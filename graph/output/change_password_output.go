package output

import "github.com/graphql-go/graphql"

func ChangePasswordOutput() *graphql.Object {
	return graphql.NewObject(
		graphql.ObjectConfig{
			Name: "ChangePasswordOutput",
			Fields: graphql.Fields{
				"password": &graphql.Field{
					Type: graphql.String,
				},
			},
		},
	)
}
