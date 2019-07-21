package routes

import (
	"github.com/ega-forever/otus_go/internal/handlers"
	"github.com/gorilla/mux"
	"net/http"
)

func SetProductRouter(r *mux.Router) {
	s := r.PathPrefix("/products").Subrouter()
	s.HandleFunc("/", handlers.GetProductsHandler).Methods(http.MethodGet)
	s.HandleFunc("/", handlers.AddProductsHandler).Methods(http.MethodPost)
	s.HandleFunc("/", handlers.RemoveProductsHandler).Methods(http.MethodDelete)
}
