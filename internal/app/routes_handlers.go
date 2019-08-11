package app

import (
	"encoding/json"
	"net/http"
)

type genericMessage struct {
	Result interface{} `json:"result"`
	Status int         `json:"status"`
}

const (
	SuccessStatus = 1
	ErrorStatus   = 0
)

func GetProductsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	products := GetProducts()

	replyMessage := genericMessage{Result: &products, Status: SuccessStatus}
	_ = json.NewEncoder(w).Encode(&replyMessage)
}

func AddProductsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var product Product

	_ = json.NewDecoder(r.Body).Decode(&product)
	savedProduct := AddProduct(product)

	replyMessage := genericMessage{Result: &savedProduct, Status: SuccessStatus}
	_ = json.NewEncoder(w).Encode(&replyMessage)
}

func RemoveProductsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var product Product
	_ = json.NewDecoder(r.Body).Decode(&product)
	RemoveProduct(product.Id)

	replyMessage := genericMessage{Result: product, Status: SuccessStatus}
	_ = json.NewEncoder(w).Encode(&replyMessage)
}
