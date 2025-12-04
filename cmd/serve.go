package cmd

import (
	"fmt"
	"net/http"
	"test/handlers"
	"test/middleware"
)

func Serve() {
	mux := http.NewServeMux()
	mux.Handle("GET /products", middleware.Logger(http.HandlerFunc(handlers.GetProducts)))
	mux.Handle("POST /products", middleware.Logger(http.HandlerFunc(handlers.CreateProduct)))
	mux.Handle("GET /products/{productId}", middleware.Logger(http.HandlerFunc(handlers.GetProductById)))
	mux.Handle("PUT /products/{productId}", middleware.Logger(http.HandlerFunc(handlers.EditProduct)))
	mux.Handle("DELETE /products/{productId}", middleware.Logger(http.HandlerFunc(handlers.DeleteProducts)))

	fmt.Println("Server running at port: 3000")

	err := http.ListenAndServe(":3000", middleware.GlobalRouter(mux))
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
