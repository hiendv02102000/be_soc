package input

import "github.com/graphql-go/graphql"

func UpdateChapterInput() *graphql.InputObject {
	return graphql.NewInputObject(graphql.InputObjectConfig{
		Name: "UpdateChapterInput",
		Fields: graphql.InputObjectConfigFieldMap{
			"title": &graphql.InputObjectFieldConfig{
				Type: graphql.String,
			},
			"chapter_id": &graphql.InputObjectFieldConfig{
				Type: graphql.Int,
			},
			"novel_id": &graphql.InputObjectFieldConfig{
				Type: graphql.Int,
			},
			"content_url": &graphql.InputObjectFieldConfig{
				Type: graphql.String,
			},
		},
	})
}
