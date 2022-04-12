package input

import "github.com/graphql-go/graphql"

func ChangeProfileInput() *graphql.InputObject {
	return graphql.NewInputObject(graphql.InputObjectConfig{
		Name: "ChangeProfileInput",
		Fields: graphql.InputObjectConfigFieldMap{
			"first_name": &graphql.InputObjectFieldConfig{
				Type: graphql.String,
			},
			"last_name": &graphql.InputObjectFieldConfig{
				Type: graphql.String,
			},
		},
	})
}
