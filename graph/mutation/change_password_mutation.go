package mutation

import (
	"be_soc/graph/input"
	"be_soc/graph/output"
	"be_soc/internal/pkg/domain/domain_model/dto"
	"be_soc/internal/pkg/domain/domain_model/entity"
	"be_soc/internal/pkg/domain/service"
	"be_soc/pkg/share/middleware"
	"be_soc/pkg/share/utils"

	"github.com/gin-gonic/gin"
	"github.com/graphql-go/graphql"
)

func ChangePasswordMutation(containerRepo map[string]interface{}) *graphql.Field {
	return &graphql.Field{
		Type:        output.ChangePasswordOutput(),
		Description: "User Register",

		Args: graphql.FieldConfigArgument{
			"user": &graphql.ArgumentConfig{
				Type: input.ChangePasswordInput(),
			},
		},
		Resolve: func(p graphql.ResolveParams) (result interface{}, err error) {
			ctx := p.Context.(*gin.Context)
			user := middleware.GetUserFromContext(ctx)
			req := p.Args["user"].(map[string]interface{})
			change_password := dto.ChangePasswordRequest{

				Password: req["password"].(string),
			}
			userRepo := containerRepo["user_repository"].(service.UserRepositoryInterface)

			change_password.Password = utils.EncryptPassword(change_password.Password)
			userRepo.UpdateUser(entity.Users{

				Password: change_password.Password,
			}, user)
			if err != nil {
				return
			}
			result = map[string]interface{}{

				"password": user.Password,
			}

			return
		},
	}
}
