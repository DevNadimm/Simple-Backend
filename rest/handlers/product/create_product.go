package product

import (
	"encoding/json"
	"net/http"
	"strings"
	"test/models"
	"test/utils"
)

func (handler *Handler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	productRepo := handler.productRepo
	var newProduct models.Product

	err := json.NewDecoder(r.Body).Decode(&newProduct)

	if err != nil {
		utils.SendData(w, http.StatusBadRequest, false, "Invalid JSON body", nil)
		return
	}

	title := strings.TrimSpace(newProduct.Title)

	if title == "" {
		utils.SendData(w, http.StatusBadRequest, false, "Title is required", nil)
		return
	}

	if newProduct.Price <= 0 {
		utils.SendData(w, http.StatusBadRequest, false, "Price must be greater than 0", nil)
		return
	}

	product, err := productRepo.Create(newProduct)

	if err != nil {
		utils.SendData(w, http.StatusInternalServerError, false, "Internal server error", nil)
		return
	}

	utils.SendData(w, http.StatusCreated, true, "Product created successfully", product)
}
