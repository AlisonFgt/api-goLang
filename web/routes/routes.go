package routes

import (
	"net/http"
	"web/controllers"
)

// LoadRoute is
func LoadRoute() {
	http.HandleFunc("/", controllers.Index)
	http.HandleFunc("/new", controllers.New)
	http.HandleFunc("/insert", controllers.Insert)
}
