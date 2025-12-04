package utils

import (
	"encoding/json"
	"net/http"
	"test/models"
)

func SendData(w http.ResponseWriter, statusCode int, status bool, message string, data interface{}) {
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(models.APIResponse{
		Status:  status,
		Message: message,
		Data:    data,
	})
}
