package rest

import (
	"net/http"
	"test/rest/handlers"
	"test/rest/middleware"
)

func initRoutes(mux *http.ServeMux, manager *middleware.Manager) {

	// ROUTES (no local middleware)
	mux.Handle("GET /products",
		manager.With(http.HandlerFunc(handlers.GetProducts)),
	)

	mux.Handle("GET /products/{productId}",
		manager.With(http.HandlerFunc(handlers.GetProductById)),
	)

	mux.Handle("POST /products",
		manager.With(http.HandlerFunc(handlers.CreateProduct)),
	)

	mux.Handle("PUT /products/{productId}",
		manager.With(http.HandlerFunc(handlers.EditProduct)),
	)

	mux.Handle("DELETE /products/{productId}",
		manager.With(http.HandlerFunc(handlers.DeleteProducts)),
	)

	// ROUTE WITH LOCAL MIDDLEWARES
	mux.Handle("GET /test",
		manager.With(
			http.HandlerFunc(handlers.GetProducts),
			middleware.NothingLocal, // local
		),
	)
}
