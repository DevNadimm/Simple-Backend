package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"test/database"
	"test/models"
	"test/utils"
)

func EditProduct(w http.ResponseWriter, r *http.Request) {
	productId := r.PathValue("productId")
	id, err := strconv.Atoi(productId)
	if err != nil {
		utils.SendData(w, http.StatusBadRequest, false, "Invalid product ID", nil)
		return
	}

	var body models.Product
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		utils.SendData(w, http.StatusBadRequest, false, "Invalid JSON body", nil)
		return
	}

	for i, p := range database.ProductList {
		if p.ID == id {
			body.ID = id
			database.ProductList[i] = body
			utils.SendData(w, http.StatusOK, true, "Product edited successfully", body)
			return
		}
	}

	utils.SendData(w, http.StatusNotFound, false, "Product not found", nil)
}
