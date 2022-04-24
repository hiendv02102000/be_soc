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
			DeleteNovel, err := novelRepo.FirstNovel(entity.Novels{
				ID: delete_novel_req.ID,
			})
			if user.ID != DeleteNovel.UsersID {
				err = errors.New("method not allowed")
			}
			if err != nil {
				return
			}
			if DeleteNovel.ID == 0 {
				err = errors.New("method not allowed")
				return
			}
			novelRepo.DeleteNovel(entity.Novels{
				ID: DeleteNovel.ID, //dieu kien
			})

			result = map[string]interface{}{
				"id":        DeleteNovel.ID,
				"name":      DeleteNovel.Name,
				"image_url": DeleteNovel.ImageUrl,
			}

			return
		},
	}
}
