package schema

import (
	"backend-food/graph/mutation"

	"github.com/graphql-go/graphql"
)

func GetAnonymousMutation(containerRepo map[string]interface{}) *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		Name: "query",
		Fields: graphql.Fields{
			"register_user": mutation.CreateUserMutation(containerRepo),
		},
	})
}
