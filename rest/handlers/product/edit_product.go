package product

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"strings"
	"test/models"
	"test/repo"
	"test/utils"
)

func (handler *Handler) EditProduct(w http.ResponseWriter, r *http.Request) {
	productRepo := handler.productRepo

	// Parse product ID
	productId := r.PathValue("productId")
	id, err := strconv.Atoi(productId)
	if err != nil {
		utils.SendData(w, http.StatusBadRequest, false, "Invalid product ID", nil)
		return
	}

	// Parse request body
	var body models.Product
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		utils.SendData(w, http.StatusBadRequest, false, "Invalid JSON body", nil)
		return
	}

	// Validate fields
	body.Title = strings.TrimSpace(body.Title)
	if body.Title == "" {
		utils.SendData(w, http.StatusBadRequest, false, "Title is required", nil)
		return
	}
	if body.Price <= 0 {
		utils.SendData(w, http.StatusBadRequest, false, "Price must be greater than 0", nil)
		return
	}

	// Set the ID
	body.ID = id

	// Update the product
	updatedProduct, err := productRepo.Update(body)
	if err != nil {
		switch {
		case errors.Is(err, repo.ErrProductNotFound):
			utils.SendData(w, http.StatusNotFound, false, "Product not found", nil)
		default:
			utils.SendData(w, http.StatusInternalServerError, false, "Internal server error", nil)
		}
		return
	}

	// Return the updated product directly
	utils.SendData(w, http.StatusOK, true, "Product edited successfully", *updatedProduct)
}
