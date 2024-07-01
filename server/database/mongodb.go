package database

import (
    "context"
    "log"
    "time"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

var MongoClient *mongo.Client

func ConnectDB() {
    clientOptions := options.Client().ApplyURI("mongodb+srv://ashutosh:ashutosh@maindb.dlpzxom.mongodb.net/?retryWrites=true&w=majority&appName=maindb")
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
    return MongoClient.Database("maindb").Collection(collectionName)
}
