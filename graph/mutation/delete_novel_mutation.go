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

func DeleteNovelMutation(containerRepo map[string]interface{}) *graphql.Field {
	return &graphql.Field{
		Type:        output.DeleteNovelOutput(),
		Description: "User Register",

		Args: graphql.FieldConfigArgument{
			"novel": &graphql.ArgumentConfig{
				Type: input.DeleteNovelInput(),
			},
		},

		Resolve: func(p graphql.ResolveParams) (result interface{}, err error) {
			ctx := p.Context.(*gin.Context)
			user := middleware.GetUserFromContext(ctx)

			req := p.Args["novel"].(map[string]interface{})
			delete_novel_req := dto.DeleteNovelRequest{
				ID: req["id"].(int),
			}

			novelRepo := containerRepo["novel_repository"].(service.NovelRepositoryInterface)
			deleteNovel, err := novelRepo.FirstNovel(entity.Novels{
				ID: delete_novel_req.ID,
			})
			if user.ID != deleteNovel.UsersID {
				err = errors.New("Method not allowed")
			}
			if err != nil {
				return
			}

			novelRepo.DeleteNovel(entity.Novels{
				ID: deleteNovel.ID, //dieu kien
			})

			result = map[string]interface{}{
				"id":        deleteNovel.ID,
				"name":      deleteNovel.Name,
				"image_url": deleteNovel.ImageUrl,
			}

			return
		},
	}
}
