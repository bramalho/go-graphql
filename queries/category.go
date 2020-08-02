package queries

import (
	"go-graphql/resolvers"
	"go-graphql/types"

	"github.com/graphql-go/graphql"
)

// GetCategoryQuery query
var GetCategoryQuery = graphql.Field{
	Type: types.Category,
	Args: graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
	},
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		return resolvers.GetCategory(params)
	},
}

// GetCategoriesQuery query
var GetCategoriesQuery = graphql.Field{
	Type: graphql.NewList(types.Category),
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		return resolvers.GetCategories(params)
	},
}
