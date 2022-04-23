package input

import "github.com/graphql-go/graphql"

func CreateChapterInput() *graphql.InputObject {
	return graphql.NewInputObject(graphql.InputObjectConfig{
		Name: "CreateChapterInput",
		Fields: graphql.InputObjectConfigFieldMap{
			"novel_id": &graphql.InputObjectFieldConfig{
				Type: graphql.Int,
			},
			"title": &graphql.InputObjectFieldConfig{
				Type: graphql.String,
			},
		},
	})
}
