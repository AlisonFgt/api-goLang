package main

import (
	"html/template"
	"net/http"
	p "web/models"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func main() {
	initializedDataBase()

	http.HandleFunc("/", index)
	http.ListenAndServe(":8000", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	products := p.SelectAllProducts()
	temp.ExecuteTemplate(w, "Index", products)
}

func initializedDataBase() {
	products := []*p.Product{
		{Nome: "Camiseta", Descricao: "Azul, bem bonita", Preco: 39, Quantidade: 5},
		{Nome: "Tenis", Descricao: "Confortável", Preco: 89, Quantidade: 3},
		{Nome: "Fone", Descricao: "Muito bom", Preco: 59, Quantidade: 2},
		{Nome: "Produto novo", Descricao: "Muito legal", Preco: 1.99, Quantidade: 1},
	}

	existsProduct := p.SelectAllProducts()

	if existsProduct == nil && len(existsProduct) == 0 {
		p.InsertProducts(products)
	}
}
