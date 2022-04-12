package output

import "github.com/graphql-go/graphql"

func ChangeProfileOutput() *graphql.Object {
	return graphql.NewObject(
		graphql.ObjectConfig{
			Name: "ChangeProfileOutput",
			Fields: graphql.Fields{
				"username": &graphql.Field{
					Type: graphql.String,
				},
				"first_name": &graphql.Field{
					Type: graphql.String,
				},
				"last_name": &graphql.Field{
					Type: graphql.String,
				},
				"role": &graphql.Field{
					Type: graphql.String,
				},
				"changed_at": &graphql.Field{
					Type: graphql.DateTime,
				},
			},
		},
	)
}
