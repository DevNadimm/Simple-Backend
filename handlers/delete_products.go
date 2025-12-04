package handlers

import (
	"net/http"
	"strconv"
	"test/database"
	"test/models"
	"test/utils"
)

func DeleteProducts(w http.ResponseWriter, r *http.Request) {
	productId := r.PathValue("productId")
	id, err := strconv.Atoi(productId)

	if err != nil {
		utils.SendData(w, http.StatusBadRequest, false, "Invalid product ID", nil)
		return
	}

	var newList []models.Product
	found := false

	for _, product := range database.ProductList {
		if product.ID == id {
			found = true
			continue
		}
		newList = append(newList, product)
	}

	database.ProductList = newList

	if found {
		utils.SendData(w, http.StatusOK, true, "Product deleted successfully", nil)
	} else {
		utils.SendData(w, http.StatusNotFound, false, "Product not found", nil)
	}
}
