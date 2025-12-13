package user

import (
	"net/http"
	"test/rest/middleware"
)

func (handler *Handler) RegisterRoutes(mux *http.ServeMux, manager *middleware.Manager) {
	mux.Handle("POST /users",
		manager.With(http.HandlerFunc(handler.CreateUser)),
	)

	mux.Handle("POST /users/login",
		manager.With(http.HandlerFunc(handler.Login)),
	)
}
