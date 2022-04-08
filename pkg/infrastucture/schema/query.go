package schema

import (
	"backend-food/graph/query"

	"github.com/graphql-go/graphql"
)

// var QueryTypes = graphql.NewObject(graphql.ObjectConfig{
// 	Name: "query",
// 	Fields: graphql.Fields{
// 		"login": query.LoginQuery(OutputTypes),
// 	},
// })

func GetAnonymousQuery(containerRepo map[string]interface{}) *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		Name: "query",
		Fields: graphql.Fields{
			"login": query.LoginQuery(containerRepo),
		},
	})
}
