package main

import (
	"fmt"
	"golang-crud-rest-api/controllers"
	"golang-crud-rest-api/infrastructure"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

var DB *gorm.DB

func main() {

	// Load Configurations from config.json using Viper
	LoadAppConfig()

	// Initialize Database
	infrastructure.InitializeDatabase(AppConfig.Database.ConnectionString)
	
	// Initialize the router
	router := mux.NewRouter().StrictSlash(true)

	// Register Routes
	RegisterProductRoutes(router)

	// Start the server
	log.Println(fmt.Sprintf("Starting Server on port %s", AppConfig.Server.Port))
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", AppConfig.Server.Port), router))
}

func RegisterProductRoutes(router *mux.Router) {
	router.HandleFunc("/api/products", controllers.GetProducts).Methods("GET")
	router.HandleFunc("/api/products/{id}", controllers.GetProductById).Methods("GET")
	router.HandleFunc("/api/product", controllers.CreateProduct).Methods("POST")
}
