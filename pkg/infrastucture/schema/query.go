package schema

import (
	"be_soc/graph/query"

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

func GetClientQuery(containerRepo map[string]interface{}) *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		Name: "query",
		Fields: graphql.Fields{
			"list_novels": query.GetListNovelsQuery(containerRepo),
		},
	})
}
func GetAdminQuery(containerRepo map[string]interface{}) *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		Name:   "query",
		Fields: graphql.Fields{},
	})
}
