package input

import "github.com/graphql-go/graphql"

func DeleteNovelInput() *graphql.InputObject {
	return graphql.NewInputObject(graphql.InputObjectConfig{
		Name: "DeleteNovelInput",
		Fields: graphql.InputObjectConfigFieldMap{
			"id": &graphql.InputObjectFieldConfig{
				Type: graphql.Int,
			},
		},
	})
}
