package output

import "github.com/graphql-go/graphql"

func DeleteChapterOutput() *graphql.Object {
	return graphql.NewObject(
		graphql.ObjectConfig{
			Name: "DeleteChapterOutput",
			Fields: graphql.Fields{
				"id": &graphql.Field{
					Type: graphql.Int,
				},
				"title": &graphql.Field{
					Type: graphql.String,
				},
				"content_url": &graphql.Field{
					Type: graphql.String,
				},
			},
		},
	)
}
