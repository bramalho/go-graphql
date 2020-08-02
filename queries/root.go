package queries

import (
	"github.com/graphql-go/graphql"
)

// Root struct
type Root struct {
	Query *graphql.Object
}

// NewRoot query
func NewRoot() *Root {
	root := Root{
		Query: graphql.NewObject(
			graphql.ObjectConfig{
				Name: "Query",
				Fields: graphql.Fields{
					"category":   &GetCategoryQuery,
					"categories": &GetCategoriesQuery,
					"product":    &GetProductQuery,
					"products":   &GetProductsQuery,
				},
			},
		),
	}

	return &root
}
