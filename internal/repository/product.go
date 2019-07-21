package repository

import (
	"github.com/ega-forever/otus_go/internal/models"
	"sync"
)

var products = make(map[int]models.Product, 0)
var productLockMutex = &sync.Mutex{}

func GetProducts() []models.Product {

	arr := make([]models.Product, len(products))

	index := 0
	for _, val := range products {
		arr[index] = val
		index++
	}

	return arr
}

func AddProduct(product models.Product) *models.Product {
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
