package category

import (
	"net/http"
	"test/utils"
)

func (handler *Handler) GetCategories(w http.ResponseWriter, r *http.Request) {
	categories, err := handler.categoryRepo.List()
	if err != nil {
		utils.SendData(w, http.StatusInternalServerError, false, "Internal server error", nil)
		return
	}

	utils.SendData(w, http.StatusOK, true, "Categories retrieved successfully", categories)
}
