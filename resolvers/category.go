package resolvers

import (
	"go-graphql/models"

	"github.com/graphql-go/graphql"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// GetCategory resolver
func GetCategory(p graphql.ResolveParams) (*models.Category, error) {
	id, ok := p.Args["id"].(string)
	if ok == false {
		return nil, nil
	}

	category, err := models.GetCategoryByID(id)
	if err != nil {
		return nil, err
	}

	return category, nil
}

// GetCategories resolver
func GetCategories(p graphql.ResolveParams) ([]models.Category, error) {
	categories, err := models.GetCategories()
	if err != nil {
		return nil, err
	}

	return categories, nil
}

// CreateCategory resolver
func CreateCategory(p graphql.ResolveParams) (*models.Category, error) {
	category, err := models.CreateCategory(&models.Category{
		Name:        p.Args["name"].(string),
		Description: p.Args["description"].(string),
		Image:       p.Args["image"].(string),
	})

	if err != nil {
		return nil, err
	}

	return category, nil
}

// UpdateCategory resolver
func UpdateCategory(p graphql.ResolveParams) (*models.Category, error) {
	id, _ := primitive.ObjectIDFromHex(p.Args["id"].(string))
	category, err := models.UpdateCategory(&models.Category{
		ID:          id,
		Name:        p.Args["name"].(string),
		Description: p.Args["description"].(string),
		Image:       p.Args["image"].(string),
	})

	if err != nil {
		return nil, err
	}

	return category, nil
}
