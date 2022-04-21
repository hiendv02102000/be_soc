package input

import "github.com/graphql-go/graphql"

func GetUserProfile() *graphql.InputObject {
	return graphql.NewInputObject(graphql.InputObjectConfig{
		Name: "GetUserProfile",
		Fields: graphql.InputObjectConfigFieldMap{
			"id": &graphql.InputObjectFieldConfig{
				Type: graphql.Int,
			},
		},
	})
}
