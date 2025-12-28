package product

import (
	"errors"
	"net/http"
	"strconv"
	"test/repo"
	"test/utils"
)

func (handler *Handler) GetProductById(w http.ResponseWriter, r *http.Request) {
	productRepo := handler.productRepo

	// Parse product ID
	productId := r.PathValue("productId")
	id, err := strconv.Atoi(productId)
	if err != nil {
		utils.SendData(w, http.StatusBadRequest, false, "Invalid product ID", nil)
		return
	}

	// Fetch product
	product, err := productRepo.Get(id)
	if err != nil {
		switch {
		case errors.Is(err, repo.ErrProductNotFound):
			utils.SendData(w, http.StatusNotFound, false, "Product not found", nil)
		default:
			utils.SendData(w, http.StatusInternalServerError, false, "Internal server error", nil)
		}
		return
	}

	// Return the product
	utils.SendData(w, http.StatusOK, true, "Product found", *product)
}
