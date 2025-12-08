package handlers

import (
	"net/http"
	"strconv"
	"test/database"
	"test/utils"
)

func GetProductById(w http.ResponseWriter, r *http.Request) {
	productId := r.PathValue("productId")
	id, err := strconv.Atoi(productId)

	if err != nil {
		utils.SendData(w, http.StatusBadRequest, false, "Invalid product ID", nil)
		return
	}

	product := database.GetProduct(id)
	if product != nil {
		utils.SendData(w, http.StatusOK, true, "Product found", *product)
	} else {
		utils.SendData(w, http.StatusNotFound, false, "Product not found", nil)
	}
}
