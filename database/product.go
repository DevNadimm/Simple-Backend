package database

import (
	"fmt"
	"test/models"
)

var productList []models.Product

func StoreProduct(product models.Product) {
	productList = append(productList, product)
}

func GetProducts() []models.Product {
	return productList
}

func GetProduct(productId int) *models.Product {
	for _, product := range productList {
		if product.ID == productId {
			return &product
		}
	}
	return nil
}

func UpdateProduct(productId int, product models.Product) {
	for idx, p := range productList {
		if p.ID == productId {
			productList[idx] = product
			return
		}
	}
}

func DeleteProduct(productId int) bool {
	newList := make([]models.Product, 0, len(productList))
	found := false

	for _, p := range productList {
		if p.ID == productId {
			found = true
			continue
		}
		newList = append(newList, p)
	}

	productList = newList
	return found
}

func GetProductCount() int {
	return len(productList)
}

func init() {
	fmt.Println("Server starting...")

	productList = append(productList,
		models.Product{ID: 1, Title: "Mango", Description: "This is good", Price: 200},
		models.Product{ID: 2, Title: "Apple", Description: "Crisp red apple", Price: 150},
		models.Product{ID: 3, Title: "Banana", Description: "Organic ripe bananas", Price: 60},
		models.Product{ID: 4, Title: "Orange", Description: "Fresh juicy oranges", Price: 120},
		models.Product{ID: 5, Title: "Watermelon", Description: "Large and sweet watermelon", Price: 300},
	)
}
