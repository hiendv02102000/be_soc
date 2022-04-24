package output

import "github.com/graphql-go/graphql"

func CreateChapterOutput() *graphql.Object {
	return graphql.NewObject(
		graphql.ObjectConfig{
			Name: "CreateChapterOutput",
			Fields: graphql.Fields{

				"id": &graphql.Field{
					Type: graphql.Int,
				},
				"title": &graphql.Field{
					Type: graphql.String,
				},
				"novel_id": &graphql.Field{
					Type: graphql.Int,
				},
				"content_url": &graphql.Field{
					Type: graphql.String,
				},
				"created_at": &graphql.Field{
					Type: graphql.DateTime,
				},
			},
		},
	)
}
