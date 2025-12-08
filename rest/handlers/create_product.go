package handlers

import (
	"encoding/json"
	"net/http"
	"test/database"
	"test/models"
	"test/utils"
)

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	var newProduct models.Product
	err := json.NewDecoder(r.Body).Decode(&newProduct)

	if err != nil {
		utils.SendData(w, http.StatusBadRequest, false, "Invalid JSON body", nil)
		return
	}

	// Auto-generate ID using the encapsulated function
	newProduct.ID = database.GetProductCount() + 1
	database.StoreProduct(newProduct)

	utils.SendData(w, http.StatusCreated, true, "Product created successfully", newProduct)
}
