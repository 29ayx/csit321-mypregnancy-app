package database

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoClient *mongo.Client

func ConnectDB() {

	// Now using .env file for the sake of security, don't want to expose the DB URI to the world.
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}

	// Use os.Getenv() to retrieve the variables from the .env file.
	clientOptions := options.Client().ApplyURI(os.Getenv("URI"))
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Connected to MongoDB!")
	MongoClient = client
}

func GetCollection(collectionName string) *mongo.Collection {
	return MongoClient.Database("my-pregnancy-dev").Collection(collectionName)
}
