package product

import (
	"log"
	"net/http"
	"test/utils"
)

func (handler *Handler) GetProducts(w http.ResponseWriter, r *http.Request) {
	productRepo := handler.productRepo
	products, err := productRepo.List()

	if err != nil {
		log.Println("PRODUCT LIST ERROR:", err)
		utils.SendData(w, http.StatusInternalServerError, false, "Internal server error", nil)
		return
	}

	utils.SendData(w, http.StatusOK, true, "Products fetched successfully", products)
}
