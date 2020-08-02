package mutations

import "github.com/graphql-go/graphql"

// Root struct
type Root struct {
	Query *graphql.Object
}

// NewRoot query
func NewRoot() *Root {
	root := Root{
		Query: graphql.NewObject(
			graphql.ObjectConfig{
				Name: "RootMutation",
				Fields: graphql.Fields{
					"createCategory": &CreateCategoryMutation,
					"updateCategory": &UpdateCategoryMutation,
					"createProduct":  &CreateProductMutation,
					"updateProduct":  &UpdateProductMutation,
				},
			},
		),
	}

	return &root
}
