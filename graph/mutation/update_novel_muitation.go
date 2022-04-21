package mutation

import (
	"be_soc/graph/input"
	"be_soc/graph/output"
	"be_soc/internal/pkg/domain/domain_model/dto"
	"be_soc/internal/pkg/domain/domain_model/entity"
	"be_soc/internal/pkg/domain/service"
	"be_soc/pkg/share/middleware"
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/graphql-go/graphql"
)

func UpdateNovelsMutation(containerRepo map[string]interface{}) *graphql.Field {
	return &graphql.Field{
		Type:        output.UpdateNovelOutput(),
		Description: "Update novel",

		Args: graphql.FieldConfigArgument{
			"novel": &graphql.ArgumentConfig{
				Type: input.UpdateNovelInput(),
			},
		},
		Resolve: func(p graphql.ResolveParams) (result interface{}, err error) {
			ctx := p.Context.(*gin.Context)
			user := middleware.GetUserFromContext(ctx)
			req := p.Args["novel"].(map[string]interface{})
			updateNovelReq := dto.UpdateNovelRequest{
				ID: req["id"].(int),
			}
			if req["imageurl"] != nil {
				updateNovelReq.Imageurl = req["imageurl"].(string)
			}
			if req["name"] != nil {
				updateNovelReq.Name = req["name"].(string)
			}
			novelRepo := containerRepo["novel_repository"].(service.NovelRepositoryInterface)
			oldnovel, err := novelRepo.FirstNovel(entity.Novels{
				ID: updateNovelReq.ID,
			})
			if user.ID != oldnovel.UsersID {
				err = errors.New("Method not allowed")
			}
			if err != nil {
				return
			}
			if updateNovelReq.ID == 0 {
				err = errors.New("Novel is unexist")
				return
			}
			novelRepo.UpdateNovel(entity.Novels{
				Name:     updateNovelReq.Name,
				ImageUrl: &updateNovelReq.Imageurl,
			}, oldnovel)
			if err != nil {
				return
			}
			result = map[string]interface{}{
				"id":       oldnovel.ID,
				"name":     oldnovel.Name,
				"imageurl": oldnovel.ImageUrl,
			}
			// timeNow := time.Now()

			// newUser := entity.Users{Username: loginReq.Username,
			// 	Password: loginReq.Password}
			return
		},
	}
}
