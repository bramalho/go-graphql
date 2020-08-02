package mutations

import (
	"go-graphql/resolvers"
	"go-graphql/types"

	"github.com/graphql-go/graphql"
)

// CreateCategoryMutation mutation
var CreateCategoryMutation = graphql.Field{
	Type: types.Category,
	Args: graphql.FieldConfigArgument{
		"name": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"description": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"image": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
	},
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		category, err := resolvers.CreateCategory(params)
		return category, err
	},
}

var UpdateCategoryMutation = graphql.Field{
	Type: types.Category,
	Args: graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"name": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"description": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"image": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
	},
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		category, err := resolvers.UpdateCategory(params)
		return category, err
	},
}
