package utils

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
)

// JWT Header structure
type Header struct {
	Alg string `json:"alg"` // Algorithm = HS256
	Typ string `json:"typ"` // Type = JWT
}

// JWT Payload structure
type Payload struct {
	ID          int    `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	IsShopOwner bool   `json:"is_shop_owner"`
}

func CreateJwt(secret string, data Payload) (string, error) {

	// STEP 1: Create the header
	header := Header{
		Alg: "HS256",
		Typ: "JWT",
	}

	// Convert header struct → JSON → bytes
	byteArrHeader, err := json.Marshal(header)
	if err != nil {
		return "", err
	}

	// Base64URL-encode header (NO padding)
	base64Header := base64.URLEncoding.WithPadding(base64.NoPadding).EncodeToString(byteArrHeader)

	// STEP 2: Create the payload
	byteArrPayload, err := json.Marshal(data)
	if err != nil {
		return "", err
	}

	// Base64URL-encode payload (NO padding)
	base64Payload := base64.URLEncoding.WithPadding(base64.NoPadding).EncodeToString(byteArrPayload)

	// STEP 3: Combine header + payload → signing input
	// Format MUST be: header.payload
	message := base64Header + "." + base64Payload

	// STEP 4: Create HMAC-SHA256 signature
	byteArrSecret := []byte(secret)  // Convert secret to bytes
	byteArrMessage := []byte(message) // Convert message to bytes

	h := hmac.New(sha256.New, byteArrSecret) // HMAC-SHA256 object
	h.Write(byteArrMessage)                  // Add message
	signature := h.Sum(nil)                  // Final signature bytes

	// STEP 5: Base64URL encode the signature
	base64Signature := base64.URLEncoding.WithPadding(base64.NoPadding).EncodeToString(signature)

	// STEP 6: Final JWT = header.payload.signature
	jwt := base64Header + "." + base64Payload + "." + base64Signature

	return jwt, nil
}
