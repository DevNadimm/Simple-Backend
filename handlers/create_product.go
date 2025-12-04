package handlers

import (
	"encoding/json"
	"net/http"
	"test/models"
	"test/database"
	"test/utils"
)

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	var newProduct models.Product
	err := json.NewDecoder(r.Body).Decode(&newProduct)

	// If invalid JSON
	if err != nil {
		utils.SendData(w, http.StatusBadRequest, false, "Invalid JSON body", nil)
		return
	}

	// Auto-generate ID
	newProduct.ID = len(database.ProductList) + 1
	database.ProductList = append(database.ProductList, newProduct)

	utils.SendData(w, http.StatusCreated, true, "Product created successfully", newProduct)
}
