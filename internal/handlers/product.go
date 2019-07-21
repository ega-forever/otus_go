package handlers

import (
	"encoding/json"
	"github.com/ega-forever/otus_go/internal/constants"
	"github.com/ega-forever/otus_go/internal/models"
	"github.com/ega-forever/otus_go/internal/repository"
	"net/http"
)

func GetProductsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	products := repository.GetProducts()

	replyMessage := constants.GenericMessage{Result: &products, Status: constants.SuccessStatus}
	_ = json.NewEncoder(w).Encode(&replyMessage)
}

func AddProductsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var product models.Product

	_ = json.NewDecoder(r.Body).Decode(&product)
	savedProduct := repository.AddProduct(product)

	replyMessage := constants.GenericMessage{Result: &savedProduct, Status: constants.SuccessStatus}
	_ = json.NewEncoder(w).Encode(&replyMessage)
}

func RemoveProductsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var product models.Product
	_ = json.NewDecoder(r.Body).Decode(&product)
	repository.RemoveProduct(product.Id)

	replyMessage := constants.GenericMessage{Result: product, Status: constants.SuccessStatus}
	_ = json.NewEncoder(w).Encode(&replyMessage)
}
