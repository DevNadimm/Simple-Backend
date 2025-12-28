package product

import (
	"errors"
	"net/http"
	"strconv"
	"test/repo"
	"test/utils"
)

func (handler *Handler) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("productId")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		utils.SendData(w, http.StatusBadRequest, false, "Invalid product ID", nil)
		return
	}

	err = handler.productRepo.Delete(id)
	if err != nil {
		switch {
		case errors.Is(err, repo.ErrProductNotFound):
			utils.SendData(w, http.StatusNotFound, false, "Product not found", nil)
		default:
			utils.SendData(w, http.StatusInternalServerError, false, "Internal server error", nil)
		}
		return
	}

	utils.SendData(w, http.StatusOK, true, "Product deleted successfully", nil)
}
