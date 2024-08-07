// database/database.go
package database

import (
    "context"
    "log"
    "os"
    "github.com/joho/godotenv"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
    "go.mongodb.org/mongo-driver/mongo/readpref"
)

var Client *mongo.Client


func Connect() error {
    // Load environment variables from config.env file
    err := godotenv.Load("config.env")
    if err != nil {
        log.Fatal("Error loading .env file")
    }

    // Get the MongoDB URI from the environment variable
    uri := os.Getenv("MONGO_URI")
    if uri == "" {
        log.Fatal("MONGO_URI is not set in the environment")
    }

    clientOptions := options.Client().ApplyURI(uri)

    Client, err = mongo.Connect(context.TODO(), clientOptions)
    if err != nil {
        return err
    }

    err = Client.Ping(context.TODO(), readpref.Primary())
    if err != nil {
        return err
    }

    log.Println("Connected to MongoDB")
    return nil
}

func Disconnect() error {
    if Client == nil {
        return nil
    }
    err := Client.Disconnect(context.TODO())
    if err != nil {
        return err
    }

    log.Println("Disconnected from MongoDB")
    return nil
}
