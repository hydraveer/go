package database

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// DBInstance initializes and returns a MongoDB client instance
func DBInstance() *mongo.Client {
	// Load the .env file
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Retrieve MongoDB URL from environment variables
	MongoDb := os.Getenv("MONGO_URL")
	if MongoDb == "" {
		log.Fatal("MONGO_URL is not set in the environment")
	}

	// Create a MongoDB client
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(MongoDb))
	if err != nil {
		log.Fatal("Failed to create MongoDB client: ", err)
	}

	// Ping the MongoDB server to verify the connection
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal("Failed to connect to MongoDB: ", err)
	}

	fmt.Println("Successfully connected to MongoDB")
	return client
}

// Global MongoDB client instance
var Client *mongo.Client = DBInstance()
