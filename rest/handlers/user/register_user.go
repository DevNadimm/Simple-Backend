package user

import (
	"encoding/json"
	"net/http"
	"test/models"
	"test/utils"
)

func (handler *Handler) RegisterUser(w http.ResponseWriter, r *http.Request) {
	var newUser models.User

	if err := json.NewDecoder(r.Body).Decode(&newUser); err != nil {
		utils.SendData(w, http.StatusBadRequest, false, "Invalid JSON body", nil)
		return
	}

	if newUser.FirstName == "" || newUser.LastName == "" || newUser.Email == "" || newUser.Password == "" {
		utils.SendData(w, http.StatusBadRequest, false, "Missing required fields", nil)
		return
	}

	// Optional: check if user already exists
	existingUser, _ := handler.userRepo.GetByEmail(newUser.Email)
	if existingUser != nil {
		utils.SendData(w, http.StatusConflict, false, "Email already exists", nil)
		return
	}

	createdUser, err := handler.userRepo.Create(newUser)
	if err != nil {
		utils.SendData(w, http.StatusInternalServerError, false, "Internal server error", nil)
		return
	}

	utils.SendData(
		w,
		http.StatusCreated,
		true,
		"Registration successful",
		createdUser.ToPublicUser(),
	)
}
