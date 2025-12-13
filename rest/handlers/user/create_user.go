package user

import (
	"encoding/json"
	"net/http"
	"test/database"
	"test/models"
	"test/utils"
)

func (handler *Handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var newUser models.User
	err := json.NewDecoder(r.Body).Decode(&newUser)

	if err != nil {
		utils.SendData(w, http.StatusBadRequest, false, "Invalid JSON body", nil)
		return
	}

	newUser.ID = database.GetUserCount() + 1
	database.CreateUser(newUser)

	utils.SendData(w, http.StatusCreated, true, "User created successfully.", newUser.ToPublicUser())
}
