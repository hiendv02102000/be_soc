package mutation

import (
	"backend-food/graph/input"
	"backend-food/graph/output"
	"backend-food/internal/pkg/domain/domain_model/dto"
	"backend-food/internal/pkg/domain/domain_model/entity"
	"backend-food/internal/pkg/domain/service"
	"backend-food/pkg/share/middleware"

	"github.com/gin-gonic/gin"
	"github.com/graphql-go/graphql"
)

func ChangeProfileMuitation(containerRepo map[string]interface{}) *graphql.Field {
	return &graphql.Field{
		Type:        output.ChangeProfileOutput(),
		Description: "Change Profile",
		//
		Args: graphql.FieldConfigArgument{
			"user": &graphql.ArgumentConfig{
				Type: input.ChangeProfileInput(),
			},
		},
		Resolve: func(p graphql.ResolveParams) (result interface{}, err error) {
			ctx := p.Context.(*gin.Context)
			user := middleware.GetUserFromContext(ctx)
			req := p.Args["user"].(map[string]interface{})
			changeProfileReq := dto.ChangeProfileRequest{
				FirstName: req["first_name"].(string),
				LastName:  req["last_name"].(string),
			}
			userRepo := containerRepo["user_repository"].(service.UserRepositoryInterface)

			userRepo.UpdateUser(entity.Users{
				FirstName: changeProfileReq.FirstName,
				LastName:  changeProfileReq.LastName,
				Role:      entity.ClientRole,
			}, user)
			if err != nil {
				return
			}
			result = map[string]interface{}{
				"username":   user.Username,
				"role":       user.Role,
				"Changed_at": user.UpdatedAt,
				"first_name": user.FirstName,
				"last_name":  user.LastName,
			}
			// timeNow := time.Now()

			// newUser := entity.Users{Username: loginReq.Username,
			// 	Password: loginReq.Password}
			return
		},
	}
}
