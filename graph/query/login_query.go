package query

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

	"github.com/graphql-go/graphql"
	uuid "github.com/satori/go.uuid"
)

func LoginQuery(containerRepo map[string]interface{}) *graphql.Field {
	return &graphql.Field{
		Type:        output.LoginOutput(),
		Description: "User Login",
		Args: graphql.FieldConfigArgument{
			"user": &graphql.ArgumentConfig{
				Type: input.LoginInput(),
			},
		},
		Resolve: func(p graphql.ResolveParams) (result interface{}, err error) {

			req := p.Args["user"].(map[string]interface{})
			loginReq := dto.LoginRequest{
				Username: req["username"].(string),
				Password: req["password"].(string),
			}

			err = utils.CheckValidate(loginReq)
			if err != nil {
				return
			}
			loginReq.Password = utils.EncryptPassword(loginReq.Password)
			userRepo := containerRepo["user_repository"].(service.UserRepositoryInterface)

			user, err := userRepo.FirstUser(entity.Users{
				Username: loginReq.Username,
				Password: loginReq.Password,
			})
			if err != nil {
				return
			}
			if user.ID == 0 {
				err = errors.New("Login fail")
				return
			}
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
			}
			err = userRepo.UpdateUser(newUser, user)
			result = map[string]interface{}{
				"token":            newUser.Token,
				"token_expried_at": newUser.TokenExpriedAt,
				"role":             user.Role,
			}

			return
		},
	}
}
