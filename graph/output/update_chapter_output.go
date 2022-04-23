package output

import "github.com/graphql-go/graphql"

func UpdateChapterOutput() *graphql.Object {
	return graphql.NewObject(
		graphql.ObjectConfig{
			Name: "UpdateChapterOutput",
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
				"user_id": &graphql.Field{
					Type: graphql.ID,
				},
				"update_at": &graphql.Field{
					Type: graphql.DateTime,
				},
			},
		},
	)
}
