package input

import "github.com/graphql-go/graphql"

func ListNovelsInput() *graphql.InputObject {
	return graphql.NewInputObject(graphql.InputObjectConfig{
		Name: "ListNovelsInput",
		Fields: graphql.InputObjectConfigFieldMap{
			"id": &graphql.InputObjectFieldConfig{
				Type: graphql.Int,
			},
			"user_id": &graphql.InputObjectFieldConfig{
				Type: graphql.Int,
			},
			"name": &graphql.InputObjectFieldConfig{
				Type: graphql.String,
			},
			"categories": &graphql.InputObjectFieldConfig{
				Type: graphql.String,
			},
			"is_get_chapters": &graphql.InputObjectFieldConfig{
				Type: graphql.Boolean,
			},
		},
	})
}
