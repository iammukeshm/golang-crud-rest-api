package controllers

import (
	"encoding/json"
	"golang-crud-rest-api/entities"
	"golang-crud-rest-api/infrastructure"
	"net/http"
)

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var product entities.Product
	json.NewDecoder(r.Body).Decode(&product)
	infrastructure.DB.Create(&product)
	json.NewEncoder(w).Encode(product)
}

func GetProductById(w http.ResponseWriter, r *http.Request) {

}

func GetProducts(w http.ResponseWriter, r *http.Request) {
	
}