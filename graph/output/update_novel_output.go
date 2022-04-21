package output

import "github.com/graphql-go/graphql"

func UpdateNovelOutput() *graphql.Object {
	return graphql.NewObject(
		graphql.ObjectConfig{
			Name: "UpdateNovelOutput",
			Fields: graphql.Fields{
				"id": &graphql.Field{
					Type: graphql.Int,
				},
				"name": &graphql.Field{
					Type: graphql.String,
				},
				"imageurl": &graphql.Field{
					Type: graphql.String,
				},
				"view": &graphql.Field{
					Type: graphql.String,
				},
				"user_id": &graphql.Field{
					Type: graphql.ID,
				},
				"changed_at": &graphql.Field{
					Type: graphql.DateTime,
				},
			},
		},
	)
}
