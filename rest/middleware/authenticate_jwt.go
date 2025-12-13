package middleware

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"net/http"
	"strings"
	"test/utils"
)

func (middleware Middleware) AuthenticateJwt(next http.Handler) http.Handler {
	handler := func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			utils.SendData(w, http.StatusUnauthorized, false, "Authorization header missing", nil)
			return
		}

		authHeaderArr := strings.Split(authHeader, " ")
		if len(authHeaderArr) != 2 || authHeaderArr[0] != "Bearer" {
			utils.SendData(w, http.StatusUnauthorized, false, "Invalid Authorization header format", nil)
			return
		}

		accessToken := authHeaderArr[1]
		if accessToken == "" {
			utils.SendData(w, http.StatusUnauthorized, false, "Access token missing", nil)
			return
		}

		tokenParts := strings.Split(accessToken, ".")
		if len(tokenParts) != 3 {
			utils.SendData(w, http.StatusUnauthorized, false, "Invalid access token", nil)
			return
		}

		jwtHeader := tokenParts[0]
		jwtPayload := tokenParts[1]
		signature := tokenParts[2]

		message := jwtHeader + "." + jwtPayload
		jwtSecret := middleware.config.JwtSecretKey

		byteArrSecret := []byte(jwtSecret)
		byteArrMessage := []byte(message)

		h := hmac.New(sha256.New, byteArrSecret)
		h.Write(byteArrMessage)
		hash := h.Sum(nil)

		expectedSignature := base64.URLEncoding.WithPadding(base64.NoPadding).EncodeToString(hash)
		if signature != expectedSignature {
			utils.SendData(w, http.StatusUnauthorized, false, "Invalid access token", nil)
			return
		}

		next.ServeHTTP(w, r)
	}

	return http.HandlerFunc(handler)
}
