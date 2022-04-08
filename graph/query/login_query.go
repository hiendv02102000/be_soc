package query

import (
	"backend-food/graph/input"
	"backend-food/graph/output"
	"backend-food/internal/pkg/domain/domain_model/dto"
	"backend-food/internal/pkg/domain/domain_model/entity"
	"backend-food/internal/pkg/repository"
	"backend-food/pkg/share/middleware"
	"time"

	"github.com/graphql-go/graphql"
	uuid "github.com/satori/go.uuid"
)

func LoginQuery(containerRepo map[string]interface{}) *graphql.Field {
	return &graphql.Field{
		Type:        output.LoginOutput(),
		Description: "User Login",
		// Args: graphql.FieldConfigArgument{
		// 	"username": &graphql.ArgumentConfig{
		// 		Type: graphql.String,
		// 	},
		// 	"password": &graphql.ArgumentConfig{
		// 		Type: graphql.NewInputObject(),
		// 	},
		//
		// },
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
			userRepo := containerRepo["user_repository"].(repository.UserRepository)
			user, err := userRepo.FirstUser(entity.Users{
				Username: loginReq.Username,
				Password: loginReq.Password,
			})
			if err != nil {
				return
			}
			// if user.Password != req.Password || user.ID == 0 {
			// 	return dto.LoginResponse{}, errors.New("Login fail")
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
			}
			err = userRepo.UpdateUser(newUser, user)
			result = newUser
			return
		},
	}
}
