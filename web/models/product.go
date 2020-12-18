package models

import (
	"context"
	"fmt"
	"log"

	mongodb "web/db"

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
	_collection = mongodb.GetCollectionToDataBase("product")
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

// InsertProduct is
func InsertProduct(product Product) {

	insertResult, err := _collection.InsertOne(context.TODO(), product)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted a single document: ", insertResult.InsertedID)
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

// InitializedDataBase is
func InitializedDataBase() {
	products := []*Product{
		{Nome: "Camiseta", Descricao: "Azul, bem bonita", Preco: 39, Quantidade: 5},
		{Nome: "Tenis", Descricao: "Confortável", Preco: 89, Quantidade: 3},
		{Nome: "Fone", Descricao: "Muito bom", Preco: 59, Quantidade: 2},
		{Nome: "Produto novo", Descricao: "Muito legal", Preco: 1.99, Quantidade: 1},
	}

	existsProduct := SelectAllProducts()

	if existsProduct == nil && len(existsProduct) == 0 {
		InsertProducts(products)
	}
}

// CreateNewProduct is
func CreateNewProduct(nome, descricao string, preco float64, quantidade int) {
	value := Product{Nome: nome, Descricao: descricao, Preco: preco, Quantidade: quantidade}
	InsertProduct(value)
}
