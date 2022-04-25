package output

import "github.com/graphql-go/graphql"

func CreateCommentOutput() *graphql.Object {
	return graphql.NewObject(
		graphql.ObjectConfig{
			Name: "CreateCommentOutput",
			Fields: graphql.Fields{
				"id": &graphql.Field{
					Type: graphql.String,
				},
				"cmt_content": &graphql.Field{
					Type: graphql.String,
				},
				"user_id": &graphql.Field{
					Type: graphql.Int,
				},
				"chapter_id": &graphql.Field{
					Type: graphql.Int,
				},
			},
		},
	)
}
