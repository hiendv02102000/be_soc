package mutation

import (
	"be_soc/graph/input"
	"be_soc/graph/output"
	"be_soc/internal/pkg/domain/domain_model/dto"
	"be_soc/internal/pkg/domain/domain_model/entity"
	"be_soc/internal/pkg/domain/service"
	"be_soc/pkg/share/middleware"
	"be_soc/pkg/share/utils"
	"errors"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/graphql-go/graphql"
	uuid "github.com/satori/go.uuid"
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
				OldPassword: req["old_password"].(string),
				Password:    req["password"].(string),
			}
			if utils.EncryptPassword(change_password.OldPassword) != user.Password {
				err = errors.New("wrong password")
				return
			}

			userRepo := containerRepo["user_repository"].(service.UserRepositoryInterface)

			change_password.Password = utils.EncryptPassword(change_password.Password)
			// userRepo.UpdateUser(entity.Users{

			// }, user)
			// if err != nil {
			// 	return
			// }

			timeNow := time.Now()
			timeExpriedAt := timeNow.Add(time.Hour * 2)
			// generate uuid
			uuid := uuid.Must(uuid.NewV4(), nil)
			tokenString, err := middleware.GenerateJWTToken(middleware.JWTParam{
				UUID:       uuid,
				Authorized: true,
				ExpriedAt:  timeExpriedAt,
			})

			if err != nil {
				return
			}

			newUser := entity.Users{
				Token:          &tokenString,
				TokenExpriedAt: &timeExpriedAt,
				Password:       change_password.Password,
			}
			err = userRepo.UpdateUser(newUser, user)
			result = map[string]interface{}{
				"token":            newUser.Token,
				"token_expried_at": newUser.TokenExpriedAt,
			}

			return
		},
	}
}
