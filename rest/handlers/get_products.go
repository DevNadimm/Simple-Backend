package handlers

import (
	"net/http"
	"test/database"
	"test/utils"
)

func GetProducts(w http.ResponseWriter, r *http.Request) {
	utils.SendData(w, http.StatusOK, true, "Products fetched successfully", database.GetProducts())
}
