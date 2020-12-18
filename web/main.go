package main

import (
	"context"
	"fmt"
	"html/template"
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var _collection *mongo.Collection

var temp = template.Must(template.ParseGlob("templates/*.html"))

func main() {
	getCollectionToDataBase("product")
	initializedDataBase()

	http.HandleFunc("/", index)
	http.ListenAndServe(":8000", nil)
}

func index(w http.ResponseWriter, r *http.Request) {

	newProducts := selectAllProducts()

	temp.ExecuteTemplate(w, "Index", newProducts)
}

func initializedDataBase() {
	products := []*Product{
		{"Camiseta", "Azul, bem bonita", 39, 5},
		{"Tenis", "Confortável", 89, 3},
		{"Fone", "Muito bom", 59, 2},
		{"Produto novo", "Muito legal", 1.99, 1},
	}

	insertProducts(products)
}

func insertProducts(products []*Product) {

	existsProduct := selectAllProducts()

	if existsProduct == nil && len(existsProduct) == 0 {
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
}

func getCollectionToDataBase(collect string) {

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
	_collection = client.Database("go_test").Collection(collect)
}
