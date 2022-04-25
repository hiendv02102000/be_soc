package input

import "github.com/graphql-go/graphql"

func CreateCommentInput() *graphql.InputObject {
	return graphql.NewInputObject(graphql.InputObjectConfig{
		Name: "CreateCommentInput",
		Fields: graphql.InputObjectConfigFieldMap{
			"chapter_id": &graphql.InputObjectFieldConfig{
				Type: graphql.Int,
			},
			"cmt_content": &graphql.InputObjectFieldConfig{
				Type: graphql.String,
			},
		},
	})
}
