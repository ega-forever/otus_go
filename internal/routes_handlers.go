package internal

import (
	"encoding/json"
	"net/http"
)

func GetProductsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	products := GetProducts()

	replyMessage := GenericMessage{Result: &products, Status: SuccessStatus}
	_ = json.NewEncoder(w).Encode(&replyMessage)
}

func AddProductsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var product Product

	_ = json.NewDecoder(r.Body).Decode(&product)
	savedProduct := AddProduct(product)

	replyMessage := GenericMessage{Result: &savedProduct, Status: SuccessStatus}
	_ = json.NewEncoder(w).Encode(&replyMessage)
}

func RemoveProductsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var product Product
	_ = json.NewDecoder(r.Body).Decode(&product)
	RemoveProduct(product.Id)

	replyMessage := GenericMessage{Result: product, Status: SuccessStatus}
	_ = json.NewEncoder(w).Encode(&replyMessage)
}
