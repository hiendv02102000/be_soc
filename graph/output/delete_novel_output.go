package output

import "github.com/graphql-go/graphql"

func DeleteNovelOutput() *graphql.Object {
	return graphql.NewObject(
		graphql.ObjectConfig{
			Name: "DeleteNovelOutput",
			Fields: graphql.Fields{

				// "delete_at": &graphql.Field{
				// 	Type: graphql.DateTime,
				// },
				"id": &graphql.Field{
					Type: graphql.Int,
				},
				"name": &graphql.Field{
					Type: graphql.String,
				},
				"image_url": &graphql.Field{
					Type: graphql.String,
				},
			},
		},
	)
}
