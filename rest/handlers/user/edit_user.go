package user

import (
	"encoding/json"
	"net/http"
	"strconv"
	"test/repo"
	"test/utils"
)

type EditRequest struct {
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	IsShopOwner *bool  `json:"is_shop_owner"`
}

func (handler *Handler) EditUser(w http.ResponseWriter, r *http.Request) {
	userRepo := handler.userRepo

	// 1. Parse user ID
	userId := r.PathValue("userId")
	id, err := strconv.Atoi(userId)
	if err != nil {
		utils.SendData(w, http.StatusBadRequest, false, "Invalid user ID", nil)
		return
	}

	// 2. Parse body
	var body EditRequest
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		utils.SendData(w, http.StatusBadRequest, false, "Invalid JSON body", nil)
		return
	}

	// 3. Get existing user
	user, err := userRepo.GetByID(id)
	if err != nil {
		if err == repo.ErrUserNotFound {
			utils.SendData(w, http.StatusNotFound, false, "User not found", nil)
			return
		}
		utils.SendData(w, http.StatusInternalServerError, false, "Internal server error", nil)
		return
	}

	// 4. Update only provided fields
	if body.FirstName != "" {
		user.FirstName = body.FirstName
	}
	if body.LastName != "" {
		user.LastName = body.LastName
	}
	if body.IsShopOwner != nil {
		user.IsShopOwner = *body.IsShopOwner
	}

	// 5. Save
	updatedUser, err := userRepo.Update(*user)
	if err != nil {
		utils.SendData(w, http.StatusInternalServerError, false, "Internal server error", nil)
		return
	}

	utils.SendData(w, http.StatusOK, true, "User updated successfully", updatedUser.ToPublicUser())
}
