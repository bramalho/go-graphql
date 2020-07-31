package resolvers

import (
	"go-graphql/models"

	"github.com/graphql-go/graphql"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// GetProduct resolver
func GetProduct(p graphql.ResolveParams) (*models.Product, error) {
	id, ok := p.Args["id"].(string)
	if ok == false {
		return nil, nil
	}

	product, err := models.GetProductByID(id)
	if err != nil {
		return nil, err
	}

	return product, nil
}

// GetProducts resolver
func GetProducts(p graphql.ResolveParams) ([]models.Product, error) {
	products, err := models.GetProducts()
	if err != nil {
		return nil, err
	}

	return products, nil
}

// CreateProduct resolver
func CreateProduct(p graphql.ResolveParams) (*models.Product, error) {
	product, err := models.CreateProduct(&models.Product{
		Name:        p.Args["name"].(string),
		Description: p.Args["description"].(string),
		Quantity:    p.Args["quantity"].(int),
		Status:      p.Args["status"].(bool),
	})

	if err != nil {
		return nil, err
	}

	return product, nil
}

// UpdateProduct resolver
func UpdateProduct(p graphql.ResolveParams) (*models.Product, error) {
	id, _ := primitive.ObjectIDFromHex(p.Args["id"].(string))
	product, err := models.UpdateProduct(&models.Product{
		ID:          id,
		Name:        p.Args["name"].(string),
		Description: p.Args["description"].(string),
		Quantity:    p.Args["quantity"].(int),
		Status:      p.Args["status"].(bool),
	})

	if err != nil {
		return nil, err
	}

	return product, nil
}
