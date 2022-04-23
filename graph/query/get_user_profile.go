package query

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

func GetUserProfile(containerRepo map[string]interface{}) *graphql.Field {
	return &graphql.Field{
		Type:        output.GetUserProfile(),
		Description: "User's Profile",

		Args: graphql.FieldConfigArgument{
			"user_profile": &graphql.ArgumentConfig{
				Type: input.GetUserProfile(),
			},
		},
		Resolve: func(p graphql.ResolveParams) (result interface{}, err error) {
			ctx := p.Context.(*gin.Context)
			user := middleware.GetUserFromContext(ctx)
			req := p.Args["user_profile"].(map[string]interface{})
			get_user_profile := dto.GetUserProfile{
				ID: req["id"].(int),
			}
			if user.ID != get_user_profile.ID {
				err = errors.New("Method not allowed")
			}
			if err != nil {
				return
			}
			novelRepo := containerRepo["novel_repository"].(service.NovelRepositoryInterface)
			userRepo := containerRepo["user_repository"].(service.UserRepositoryInterface)

			user_profile, err := userRepo.FirstUser(entity.Users{

				ID: get_user_profile.ID,
			})
			if err != nil {
				return
			}
			novel, err := novelRepo.FindNovelList(entity.Novels{
				UsersID: user.ID,
			})

			if err != nil {
				return
			}
			result = map[string]interface{}{

				"ID":         user_profile.ID,
				"first_name": user_profile.FirstName,
				"last_name":  user_profile.LastName,
				"username":   user_profile.Username,
				"novel":      novel,
			}

			return
		},
	}
}
