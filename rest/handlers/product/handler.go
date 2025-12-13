package product

import "test/rest/middleware"

type Handler struct {
	middleware *middleware.Middleware
}

func NewHandler(middleware *middleware.Middleware) *Handler {
	return &Handler{
		middleware: middleware,
	}
}
