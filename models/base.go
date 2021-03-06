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
		log.Println(err)
	}

	if len(os.Getenv("DB_CONNECTION")) == 0 {
		log.Fatal("DB connection string not found")
	}

	setUpDB()
}

func setUpDB() {
	log.Println("Connection to database " + os.Getenv("DB_CONNECTION"))

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	clientOptions := options.Client().ApplyURI(os.Getenv("DB_CONNECTION"))
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	db = client.Database(os.Getenv("DB_NAME"))

	log.Println("Database connected to " + os.Getenv("DB_NAME"))
}

// GetDB instance
func GetDB() *mongo.Database {
	return db
}
