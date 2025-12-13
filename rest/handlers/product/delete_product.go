package product

import (
	"net/http"
	"strconv"
	"test/database"
	"test/utils"
)

func (handler *Handler) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	productId := r.PathValue("productId")
	id, err := strconv.Atoi(productId)

	if err != nil {
		utils.SendData(w, http.StatusBadRequest, false, "Invalid product ID", nil)
		return
	}

	found := database.DeleteProduct(id)

	if found {
		utils.SendData(w, http.StatusOK, true, "Product deleted successfully", nil)
	} else {
		utils.SendData(w, http.StatusNotFound, false, "Product not found", nil)
	}
}
