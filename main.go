package main

import (
	"net/http"
	"rest_api/controllers/productcontroller"
	"rest_api/models"

	"github.com/gorilla/mux"
)

func main() {
	models.ConnectDatabase()
	r := mux.NewRouter()

	r.HandleFunc("/products", productcontroller.Index).Methods("GET")
	r.HandleFunc("/product/{id}", productcontroller.Show).Methods("GET")
	r.HandleFunc("/product", productcontroller.Create).Methods("POST")
	r.HandleFunc("/product/{id}", productcontroller.Update).Methods("PUT")
	r.HandleFunc("/product", productcontroller.Delete).Methods("DELETE")

	http.ListenAndServe(":5000", r)
}
