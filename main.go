package main

import (
	"golang-crud-rest-api/controllers"

	"github.com/gorilla/mux"
)

func main() {
	LoadAppConfig()
	router := mux.NewRouter().StrictSlash(true)
	InitProductHandlers(router);
}

func InitProductHandlers(router *mux.Router) {
	router.HandleFunc("/products/{id}", controllers.GetProductById).Methods("GET")
}
