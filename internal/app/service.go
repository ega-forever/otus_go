package app

import (
	"sync"
)

type Product struct {
	Name string `json:"name"`
	Id   int    `json:"id"`
}

var products = make(map[int]Product, 0)
var productLockMutex = &sync.Mutex{}

func GetProducts() []Product {

	arr := make([]Product, len(products))

	index := 0
	for _, val := range products {
		arr[index] = val
		index++
	}

	return arr
}

func AddProduct(product Product) *Product {
	productLockMutex.Lock()
	products[product.Id] = product
	productLockMutex.Unlock()
	return &product
}

func RemoveProduct(id int) {
	productLockMutex.Lock()
	delete(products, id)
	productLockMutex.Unlock()
}
