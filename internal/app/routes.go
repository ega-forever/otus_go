package app

import (
	"github.com/gorilla/mux"
	"net/http"
)

func SetProductRouter(r *mux.Router) {
	s := r.PathPrefix("/products").Subrouter()
	s.HandleFunc("/", GetProductsHandler).Methods(http.MethodGet)
	s.HandleFunc("/", AddProductsHandler).Methods(http.MethodPost)
	s.HandleFunc("/", RemoveProductsHandler).Methods(http.MethodDelete)
}
