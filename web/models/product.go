package models

import (
	"context"
	"log"

	"./web/db"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var _collection *mongo.Collection

// Product is
type Product struct {
	Nome       string
	Descricao  string
	Preco      float64
	Quantidade int
}

func selectAllProducts() []*Product {

	_collection := db.GetCollectionToCollection("product")

	var results []*Product
	findOptions := options.Find()
	findOptions.SetLimit(10)

	cur, err := _collection.Find(context.TODO(), bson.D{{}}, findOptions)
	if err != nil {
		log.Fatal(err)
	}

	for cur.Next(context.TODO()) {
		var elem Product
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}

		results = append(results, &elem)
	}

	return results
}
