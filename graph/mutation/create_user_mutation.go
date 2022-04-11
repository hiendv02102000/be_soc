package mutation

import (
	"be_soc/graph/input"
	"be_soc/graph/output"
	"be_soc/internal/pkg/domain/domain_model/dto"
	"be_soc/internal/pkg/domain/domain_model/entity"
	"be_soc/internal/pkg/domain/service"
	"be_soc/pkg/share/utils"
	"errors"

	"github.com/graphql-go/graphql"
)

func CreateUserMutation(containerRepo map[string]interface{}) *graphql.Field {
	return &graphql.Field{
		Type:        output.CreateUserOutput(),
		Description: "User Register",

		Args: graphql.FieldConfigArgument{
			"user": &graphql.ArgumentConfig{
				Type: input.CreateUserInput(),
			},
		},
		Resolve: func(p graphql.ResolveParams) (result interface{}, err error) {
			req := p.Args["user"].(map[string]interface{})
			createUserReq := dto.CreateUserRequest{
				Username:  req["username"].(string),
				Password:  req["password"].(string),
				FirstName: req["first_name"].(string),
				LastName:  req["last_name"].(string),
			}
			userRepo := containerRepo["user_repository"].(service.UserRepositoryInterface)

			user, err := userRepo.FirstUser(entity.Users{
				Username: createUserReq.Username,
			})
			if err != nil {
				return
			}
			if user.ID != 0 {
				err = errors.New("User is exist")
				return
			}
			createUserReq.Password = utils.EncryptPassword(createUserReq.Password)
			user, err = userRepo.CreateUser(entity.Users{

				FirstName: createUserReq.FirstName,
				LastName:  createUserReq.LastName,
				Username:  createUserReq.Username,
				Password:  createUserReq.Password,
				Role:      entity.ClientRole,
			})
			if err != nil {
				return
			}
			result = map[string]interface{}{
				"username":   user.Username,
				"role":       user.Role,
				"created_at": user.CreatedAt,
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
