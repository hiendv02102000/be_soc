package input

import "github.com/graphql-go/graphql"

func CreateNovelsInput() *graphql.InputObject {
	return graphql.NewInputObject(graphql.InputObjectConfig{
		Name: "CreateNovelsInput",
		Fields: graphql.InputObjectConfigFieldMap{
			"name": &graphql.InputObjectFieldConfig{
				Type: graphql.String,
			},
			"categories_id": &graphql.InputObjectFieldConfig{
				Type: &graphql.List{OfType: graphql.Int},
			},
		},
	})
}
