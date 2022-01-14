package routes

import (
	"github.com/sergyomsilva/api-rest-golang/middleware"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sergyomsilva/api-rest-golang/controllers"
)

func HandleResquest() {
	route := mux.NewRouter()
	route.Use(middleware.ContentTypeMiddleware)

	route.HandleFunc("/api/products", controllers.Index).Methods("Get")
	route.HandleFunc("/api/products/{id}", controllers.Show).Methods("Get")
	route.HandleFunc("/api/products/", controllers.Create).Methods("Post")
	route.HandleFunc("/api/products/{id}", controllers.Update).Methods("Put")
	route.HandleFunc("/api/products/{id}", controllers.Delete).Methods("Delete")
	log.Fatal(http.ListenAndServe(":8000", route))
}
