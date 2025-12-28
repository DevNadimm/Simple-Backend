package user

import (
	"encoding/json"
	"errors"
	"net/http"
	"test/models"
	"test/repo"
	"test/utils"
)

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	User        models.PublicUser `json:"user"`
	AccessToken string            `json:"access_token"`
}

func (handler *Handler) Login(w http.ResponseWriter, r *http.Request) {
	var loginReq LoginRequest

	if err := json.NewDecoder(r.Body).Decode(&loginReq); err != nil {
		utils.SendData(w, http.StatusBadRequest, false, "Invalid JSON body", nil)
		return
	}

	user, err := handler.userRepo.GetByEmail(loginReq.Email)
	if err != nil {
		if errors.Is(err, repo.ErrUserNotFound) {
			utils.SendData(w, http.StatusNotFound, false, "User not registered", nil)
			return
		}
		utils.SendData(w, http.StatusInternalServerError, false, "Internal server error", nil)
		return
	}

	// Plain-text password check
	if user.Password != loginReq.Password {
		utils.SendData(w, http.StatusUnauthorized, false, "Invalid password", nil)
		return
	}

	accessToken, err := utils.CreateJwt(
		handler.config.JwtSecretKey,
		utils.Payload{
			ID:          user.ID,
			FirstName:   user.FirstName,
			LastName:    user.LastName,
			Email:       user.Email,
			IsShopOwner: user.IsShopOwner,
		},
	)

	if err != nil {
		utils.SendData(w, http.StatusInternalServerError, false, "Failed to create access token", nil)
		return
	}

	utils.SendData(w, http.StatusOK, true, "Login successful", LoginResponse{
		User:        user.ToPublicUser(),
		AccessToken: accessToken,
	})
}
