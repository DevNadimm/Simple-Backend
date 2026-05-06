package category

import (
	"errors"
	"net/http"
	"strconv"
	"test/repo"
	"test/utils"
)

func (handler *Handler) DeleteCategory(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("categoryId")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		utils.SendData(w, http.StatusBadRequest, false, "Invalid category ID", nil)
		return
	}

	err = nil

	err = handler.categoryRepo.Delete(int64(id))
	if err != nil {
		switch {
		case errors.Is(err, repo.ErrCategoryNotFound):
			utils.SendData(w, http.StatusNotFound, false, "Category not found", nil)
		default:
			utils.SendData(w, http.StatusInternalServerError, false, "Internal server error", nil)
		}
		return
	}

	utils.SendData(w, http.StatusOK, true, "Category deleted successfully", nil)
}
