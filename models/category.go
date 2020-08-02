package models

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Category struct
type Category struct {
	ID          primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name        string             `json:"name,omitempty" bson:"name,omitempty"`
	Description string             `json:"description,omitempty" bson:"description,omitempty"`
	Image       string             `json:"image,omitempty" bson:"image,omitempty"`
}

// CreateCategory into collection
func CreateCategory(category *Category) (*Category, error) {
	collection := GetDB().Collection("categories")
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)

	result, err := collection.InsertOne(ctx, category)
	if err != nil {
		return nil, err
	}

	category.ID = result.InsertedID.(primitive.ObjectID)

	return category, nil
}

// CreateCategory into collection
func UpdateCategory(category *Category) (*Category, error) {
	collection := GetDB().Collection("categories")
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)

	filter := bson.M{"_id": category.ID}
	update := bson.D{{"$set",
		bson.D{
			{"name", category.Name},
			{"description", category.Description},
			{"image", category.Image},
		},
	}}

	_, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return nil, err
	}

	return category, nil
}

// GetCategoryByID into collection
func GetCategoryByID(id string) (*Category, error) {
	category := Category{}
	collection := GetDB().Collection("categories")
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)

	oid, _ := primitive.ObjectIDFromHex(id)
	err := collection.FindOne(ctx, Category{ID: oid}).Decode(&category)
	if err != nil {
		return nil, err
	}

	return &category, nil
}

// GetCategories into collection
func GetCategories() ([]Category, error) {
	var categories []Category
	collection := GetDB().Collection("categories")
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)

	cursor, err := collection.Find(ctx, Category{})
	if err != nil {
		return nil, err
	}

	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var category Category
		cursor.Decode(&category)
		categories = append(categories, category)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return categories, nil
}
