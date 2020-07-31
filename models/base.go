package models

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client
var db *mongo.Database

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	setUpDB()
}

func setUpDB() {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	clientOptions := options.Client().ApplyURI(os.Getenv("db_connection_string"))
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	db = client.Database(os.Getenv("db_name"))

	log.Println("Database connected to " + os.Getenv("db_name"))
}

// GetDB instance
func GetDB() *mongo.Database {
	return db
}
