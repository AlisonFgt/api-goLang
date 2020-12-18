package main

import (
	"net/http"
	"web/models"
	"web/routes"
)

func main() {
	models.InitializedDataBase()
	routes.LoadRoute()
	http.ListenAndServe(":8000", nil)
}
