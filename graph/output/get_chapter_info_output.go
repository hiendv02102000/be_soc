package output

import "github.com/graphql-go/graphql"

func GetChapterInfoOutput() *graphql.Object {
	return graphql.NewObject(
		graphql.ObjectConfig{
			Name: "GetChapterInfo",
			Fields: graphql.Fields{
				"id": &graphql.Field{
					Type: graphql.Int,
				},
				"novel_id": &graphql.Field{
					Type: graphql.Int,
				},
				"title": &graphql.Field{
					Type: graphql.String,
				},
				"content_url": &graphql.Field{
					Type: graphql.String,
				},
				"comments": &graphql.Field{
					Type: &graphql.List{OfType: graphql.NewObject(
						graphql.ObjectConfig{
							Name: "CommentsChapter",
							Fields: graphql.Fields{
								"id": &graphql.Field{
									Type: graphql.Int,
								},
								"username": &graphql.Field{
									Type: graphql.String,
								},
								"comment_content": &graphql.Field{
									Type: graphql.String,
								},
							},
						},
					)},
				},
			},
		},
	)
}
