package output

import "github.com/graphql-go/graphql"

func CreateUserOutput() *graphql.Object {
	return graphql.NewObject(
		graphql.ObjectConfig{
			Name: "LoginOutput",
			Fields: graphql.Fields{
				"id": &graphql.Field{
					Type: graphql.Int,
				},
				"name": &graphql.Field{
					Type: graphql.String,
				},
				"user": &graphql.Field{
					Type: graphql.String,
				},
				"categories": &graphql.Field{
					Type: graphql.String,
				},
				"img_url": &graphql.Field{
					Type: graphql.String,
				},
				"created_at": &graphql.Field{
					Type: graphql.DateTime,
				},
			},
		},
	)
}
