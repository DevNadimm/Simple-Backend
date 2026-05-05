package category

import (
	"encoding/json"
	"net/http"
	"strings"
	"test/models"
	"test/repo"
	"test/utils"
)

func (handler *Handler) CreateCategory(w http.ResponseWriter, r *http.Request) {
	var category models.Category

	if err := json.NewDecoder(r.Body).Decode(&category); err != nil {
		utils.SendData(w, http.StatusBadRequest, false, "Invalid JSON body", nil)
		return
	}

	category.Name = strings.TrimSpace(category.Name)
	if category.Name == "" {
		utils.SendData(w, http.StatusBadRequest, false, "Name is required", nil)
		return
	}

	if category.ParentID != nil {
		_, err := handler.categoryRepo.GetByID(int(*category.ParentID))
		if err != nil {
			if err == repo.ErrCategoryNotFound {
				utils.SendData(w, http.StatusNotFound, false, "Parent category not found", nil)
			} else {
				utils.SendData(w, http.StatusInternalServerError, false, "Internal server error", nil)
			}
			return
		}
	}

	res, err := handler.categoryRepo.Create(category)
	if err != nil {
		utils.SendData(w, http.StatusInternalServerError, false, "Internal server error", nil)
		return
	}

	utils.SendData(w, http.StatusCreated, true, "Category created successfully", res)
}
