package user

import (
	"encoding/json"
	"net/http"
	"test/config"
	"test/database"
	"test/models"
	"test/utils"
)

type ReqData struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	User        models.PublicUser `json:"user"`
	AccessToken string            `json:"access_token"`
}

func (handler *Handler) Login(w http.ResponseWriter, r *http.Request) {
	var reqData ReqData
	err := json.NewDecoder(r.Body).Decode(&reqData)
	if err != nil {
		utils.SendData(w, http.StatusBadRequest, false, "Invalid JSON body", nil)
		return
	}

	user := database.GetUserByEmail(reqData.Email)
	if user == nil {
		utils.SendData(w, http.StatusNotFound, false, "User not found", nil)
		return
	}

	if !database.CheckPasswordCorrect(*user, reqData.Password) {
		utils.SendData(w, http.StatusUnauthorized, false, "Invalid password", nil)
		return
	}

	jwtSecretKey := config.GetConfig().JwtSecretKey

	accessToken, err := utils.CreateJwt(jwtSecretKey, utils.Payload{
		ID:          user.ID,
		FirstName:   user.FirstName,
		LastName:    user.LastName,
		Email:       user.Email,
		IsShopOwner: user.IsShopOwner,
	})

	if err != nil {
		utils.SendData(w, http.StatusInternalServerError, false, "Failed to create access token", nil)
		return
	}

	utils.SendData(w, http.StatusOK, true, "Login successful", LoginResponse{
		User:        user.ToPublicUser(),
		AccessToken: accessToken,
	})
}
