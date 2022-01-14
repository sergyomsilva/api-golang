package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sergyomsilva/api-golang/database"
	"github.com/sergyomsilva/api-golang/models"
)

func Index(w http.ResponseWriter, r *http.Request) {
	var products []models.Product
	database.DB.Find(&products)
	json.NewEncoder(w).Encode(products)
}

func Show(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var product models.Product
	database.DB.First(&product, id)
	json.NewEncoder(w).Encode(product)
}

func Create(w http.ResponseWriter, r *http.Request) {
	var product models.Product
	json.NewDecoder(r.Body).Decode(&product)
	database.DB.Create(&product)
	json.NewEncoder(w).Encode(&product)
}

func Update(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var product models.Product
	database.DB.First(&product, id)
	json.NewDecoder(r.Body).Decode(&product)
	database.DB.Save(&product)
	json.NewEncoder(w).Encode(&product)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var product models.Product
	database.DB.Delete(&product, id)
	json.NewEncoder(w).Encode(&product)
}
