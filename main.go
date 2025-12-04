package main

import (
	"fmt"
	"test/cmd"
	"test/database"
	"test/models"
)

func main() {
	cmd.Serve()
}

func init() {
	fmt.Println("Server starting...")

	database.ProductList = append(database.ProductList,
		models.Product{ID: 1, Title: "Mango", Description: "This is good", Price: 200},
		models.Product{ID: 2, Title: "Apple", Description: "Crisp red apple", Price: 150},
		models.Product{ID: 3, Title: "Banana", Description: "Organic ripe bananas", Price: 60},
		models.Product{ID: 4, Title: "Orange", Description: "Fresh juicy oranges", Price: 120},
		models.Product{ID: 5, Title: "Watermelon", Description: "Large and sweet watermelon", Price: 300},
	)
}
