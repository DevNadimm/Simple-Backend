package product

import (
	"net/http"
	"test/utils"
)

func (handler *Handler) GetProducts(w http.ResponseWriter, r *http.Request) {
	productRepo := handler.productRepo
	products, err := productRepo.List()

	if err != nil {
		utils.SendData(w, http.StatusInternalServerError, false, "Internal server error", nil)
		return
	}

	utils.SendData(w, http.StatusOK, true, "Products fetched successfully", products)
}
