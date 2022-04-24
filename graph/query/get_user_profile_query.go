package query

import (
	"be_soc/graph/output"
	"be_soc/internal/pkg/domain/domain_model/entity"
	"be_soc/internal/pkg/domain/service"
	"be_soc/pkg/share/middleware"

	"github.com/gin-gonic/gin"
	"github.com/graphql-go/graphql"
)

func GetUserProfileQuery(containerRepo map[string]interface{}) *graphql.Field {
	return &graphql.Field{
		Type:        output.GetUserProfile(),
		Description: "User's Profile",

		Resolve: func(p graphql.ResolveParams) (result interface{}, err error) {
			ctx := p.Context.(*gin.Context)
			user := middleware.GetUserFromContext(ctx)

			if err != nil {
				return
			}
			novelRepo := containerRepo["novel_repository"].(service.NovelRepositoryInterface)
			userRepo := containerRepo["user_repository"].(service.UserRepositoryInterface)

			userProfile, err := userRepo.FirstUser(entity.Users{

				ID: user.ID,
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

				"ID":         userProfile.ID,
				"first_name": userProfile.FirstName,
				"last_name":  userProfile.LastName,
				"username":   userProfile.Username,
				"novel":      novel,
			}

			return
		},
	}
}
