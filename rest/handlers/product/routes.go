package product

import (
	"net/http"
	"test/rest/middleware"
)

func (handler *Handler) RegisterRoutes(mux *http.ServeMux, manager *middleware.Manager) {
	mux.Handle("GET /products",
		manager.With(http.HandlerFunc(handler.GetProducts)),
	)

	mux.Handle("GET /products/{productId}",
		manager.With(http.HandlerFunc(handler.GetProductById)),
	)

	mux.Handle("POST /products",
		manager.With(http.HandlerFunc(handler.CreateProduct), handler.middleware.AuthenticateJwt),
	)

	mux.Handle("PUT /products/{productId}",
		manager.With(http.HandlerFunc(handler.EditProduct), handler.middleware.AuthenticateJwt),
	)

	mux.Handle("DELETE /products/{productId}",
		manager.With(http.HandlerFunc(handler.DeleteProduct), handler.middleware.AuthenticateJwt),
	)
}
