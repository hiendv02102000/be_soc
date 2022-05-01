package output

import "github.com/graphql-go/graphql"

func GetUserProfile() *graphql.Object {
	return graphql.NewObject(
		graphql.ObjectConfig{
			Name: "GetUserProfile",
			Fields: graphql.Fields{
				"id": &graphql.Field{
					Type: graphql.Int,
				},
				"first_name": &graphql.Field{
					Type: graphql.String,
				},
				"last_name": &graphql.Field{
					Type: graphql.String,
				},
				"username": &graphql.Field{
					Type: graphql.String,
				},
				"novel": &graphql.Field{
					Type: &graphql.List{OfType: graphql.NewObject(
						graphql.ObjectConfig{
							Name: "NovelList",
							Fields: graphql.Fields{
								"id": &graphql.Field{
									Type: graphql.Int,
								},
								"name": &graphql.Field{
									Type: graphql.String,
								},
								"img_url": &graphql.Field{
									Type: graphql.String,
								},
								"view": &graphql.Field{
									Type: graphql.Int,
								},
							},
						},
					)},
				},
			},
		},
	)
}
