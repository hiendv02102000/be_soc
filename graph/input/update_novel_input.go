package input

import "github.com/graphql-go/graphql"

func UpdateNovelInput() *graphql.InputObject {
	return graphql.NewInputObject(graphql.InputObjectConfig{
		Name: "UpdateNovelInput",
		Fields: graphql.InputObjectConfigFieldMap{
			"name": &graphql.InputObjectFieldConfig{
				Type: graphql.String,
			},
			"id": &graphql.InputObjectFieldConfig{
				Type: graphql.Int,
			},
			"image_url": &graphql.InputObjectFieldConfig{
				Type: graphql.String,
			},
		},
	})
}
