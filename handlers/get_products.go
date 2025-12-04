package handlers

import (
	"net/http"
	"test/utils"
	"test/database"
)


func GetProducts(w http.ResponseWriter, r *http.Request) {
	utils.SendData(w, http.StatusOK, true, "Products fetched successfully", database.ProductList)
}
