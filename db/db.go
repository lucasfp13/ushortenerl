package db

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client
var URLCollection *mongo.Collection

func MongoConnect() error {
	dbUri := os.Getenv("MONGO_URI")
	dbName := os.Getenv("MONGO_DB")
	dbCollection := os.Getenv("MONGO_COLLECTION")
	clientOptions := options.Client().ApplyURI(dbUri)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return err
	}

	if err := client.Ping(ctx, nil); err != nil {
		log.Fatal("mongodb is not available:", err)
	}

	Client = client
	URLCollection = Client.Database(dbName).Collection(dbCollection)

	// set unique index
	_, err = URLCollection.Indexes().CreateOne(ctx, mongo.IndexModel{
		Keys:    bson.M{"short": 1},
		Options: options.Index().SetUnique(true),
	})
	if err != nil {
		return fmt.Errorf("failed to create unique index: %v", err)
	}

	fmt.Println("Connected to MongoDB...")
	return nil
}
