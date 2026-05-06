package category

import (
	"net/http"
	"test/rest/middleware"
)

func (handler *Handler) RegisterRoutes(mux *http.ServeMux, manager *middleware.Manager) {
	mux.Handle("POST /categories",
		manager.With(http.HandlerFunc(handler.CreateCategory), handler.middleware.AuthenticateJwt),
	)

	mux.Handle("GET /categories",
		manager.With(http.HandlerFunc(handler.GetCategories), handler.middleware.AuthenticateJwt),
	)

	mux.Handle("DELETE /categories/{categoryId}",
		manager.With(http.HandlerFunc(handler.DeleteCategory), handler.middleware.AuthenticateJwt),
	)
}
