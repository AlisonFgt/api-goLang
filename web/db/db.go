package mongodb

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// GetCollectionToDataBase is
func GetCollectionToDataBase(collect string) *mongo.Collection {

	// Set client options
	clientOptions := options.Client().ApplyURI("mongodb://localhost:17017/?safe=true")

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")

	// Get a handle for your collection
	return client.Database("go_test").Collection(collect)
}
