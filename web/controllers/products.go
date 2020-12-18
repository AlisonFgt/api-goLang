package controllers

import (
	"html/template"
	"log"
	"net/http"
	"strconv"
	"web/models"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

// Index is
func Index(w http.ResponseWriter, r *http.Request) {
	products := models.SelectAllProducts()
	temp.ExecuteTemplate(w, "Index", products)
}

// New is
func New(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "New", nil)
}

// Insert is
func Insert(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		nome := r.FormValue("nome")
		descricao := r.FormValue("descricao")
		preco := r.FormValue("preco")
		quantidade := r.FormValue("quantidade")

		precoConvert, err := strconv.ParseFloat(preco, 64)

		if err != nil {
			log.Println("erro na conversao do preco:", err)
		}

		qtdConvert, err := strconv.Atoi(quantidade)

		if err != nil {
			log.Println("erro na conversao do quantidade:", err)
		}

		models.CreateNewProduct(nome, descricao, precoConvert, qtdConvert)
	}

	http.Redirect(w, r, "/", 301)
}
