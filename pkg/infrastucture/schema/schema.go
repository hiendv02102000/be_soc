package schema

import (
	"backend-food/pkg/infrastucture/db"

	"github.com/graphql-go/graphql"
)

func NewAnonymousSchema(database db.Database) *graphql.Schema {
	repoContainer := GetContainerRepo(database)
	myschema, _ := graphql.NewSchema(
		graphql.SchemaConfig{
			Query: GetAnonymousQuery(repoContainer),
		},
	)
	return &myschema
}
