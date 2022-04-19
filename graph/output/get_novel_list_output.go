package output

import "github.com/graphql-go/graphql"

func NovelListOutput() *graphql.Object {
	return graphql.NewObject(
		graphql.ObjectConfig{
			Name: "NovelListOutput",
			Fields: graphql.Fields{
				"list": &graphql.Field{
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
								"user_id": &graphql.Field{
									Type: graphql.Int,
								},
								"chapter": &graphql.Field{
									Type: &graphql.List{OfType: graphql.NewObject(
										graphql.ObjectConfig{
											Name: "ChaptersList",
											Fields: graphql.Fields{
												"id": &graphql.Field{
													Type: graphql.Int,
												},
												"title": &graphql.Field{
													Type: graphql.String,
												},
												"content": &graphql.Field{
													Type: graphql.String,
												},
											},
										},
									)},
								},
								"categories": &graphql.Field{
									Type: &graphql.List{OfType: graphql.NewObject(
										graphql.ObjectConfig{
											Name: "ChaptersList",
											Fields: graphql.Fields{
												"name": &graphql.Field{
													Type: graphql.String,
												},
											},
										},
									)},
								},
							},
						},
					)},
				},
			},
		},
	)
}
