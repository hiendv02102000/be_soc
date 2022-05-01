package output

import "github.com/graphql-go/graphql"

func GetAllCategoriesOutput() *graphql.Object {
	return graphql.NewObject(
		graphql.ObjectConfig{
			Name: "CategoriesListOutput",
			Fields: graphql.Fields{
				"list": &graphql.Field{
					Type: &graphql.List{OfType: graphql.NewObject(
						graphql.ObjectConfig{
							Name: "CategoryList",
							Fields: graphql.Fields{
								"id": &graphql.Field{
									Type: graphql.Int,
								},
								"name": &graphql.Field{
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
