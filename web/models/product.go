package product

import (
	"context"
	"fmt"
	"log"

	dab "web/db"

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

func init() {
	_collection = dab.GetCollectionToDataBase("product")
}

// InsertProducts is
func InsertProducts(products []*Product) {

	collects := make([]interface{}, len(products))
	for i, s := range products {
		collects[i] = s
	}

	insertManyResult, err := _collection.InsertMany(context.TODO(), collects)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted multiple documents: ", insertManyResult.InsertedIDs)
}

// SelectAllProducts is
func SelectAllProducts() []*Product {

	var results []*Product
	findOptions := options.Find()

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
