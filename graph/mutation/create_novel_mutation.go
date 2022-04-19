package mutation

import (
	"be_soc/graph/input"
	"be_soc/graph/output"

	"github.com/graphql-go/graphql"
)

func CreateNovelMutation(containerRepo map[string]interface{}) *graphql.Field {
	return &graphql.Field{
		Type:        output.CreateUserOutput(),
		Description: "User Register",

		Args: graphql.FieldConfigArgument{
			"user": &graphql.ArgumentConfig{
				Type: input.CreateUserInput(),
			},
		},
		Resolve: func(p graphql.ResolveParams) (result interface{}, err error) {

			return
		},
	}
}
