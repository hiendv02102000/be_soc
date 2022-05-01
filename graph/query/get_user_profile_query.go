package query

import (
	"be_soc/graph/output"
	"be_soc/internal/pkg/domain/domain_model/entity"
	"be_soc/internal/pkg/domain/service"
	"be_soc/pkg/share/middleware"
	"fmt"

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

			UserProfile, err := userRepo.FirstUser(entity.Users{

				ID: user.ID,
			})
			if err != nil {
				return
			}
			// novel, err := novelRepo.FindNovelList(entity.Novels{
			// 	UsersID: user.ID,
			// })
			userNovel, err := novelRepo.FindNovelList(entity.Novels{
				UsersID: user.ID,
			})
			fmt.Println(userNovel)
			novels := make([]map[string]interface{}, 0)
			for i := 0; i < len(novels); i++ {
				n, err1 := novelRepo.FirstNovel(entity.Novels{
					UsersID: userNovel[i].UsersID,
				})
				if err1 != nil {
					return
				}
				fmt.Println(n)
				novel := map[string]interface{}{
					"id":      n.ID,
					"name":    n.Name,
					"img_url": n.ImageUrl,
					"view":    n.View,
				}
				novels = append(novels, novel)
			}

			if err != nil {
				return
			}
			result = map[string]interface{}{

				"id":         UserProfile.ID,
				"first_name": UserProfile.FirstName,
				"last_name":  UserProfile.LastName,
				"username":   UserProfile.Username,
				"novel":      novels,
			}

			return
		},
	}
}
