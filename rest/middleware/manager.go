package middleware

import (
	"net/http"
)

type Middleware func(next http.Handler) http.Handler

type Manager struct {
	globalMiddlewares []Middleware
}

func NewManger() *Manager {
	manager := Manager{
		globalMiddlewares: make([]Middleware, 0),
	}

	return &manager
}

func (manager *Manager) Use(middlewares ...Middleware) {
	manager.globalMiddlewares = append(manager.globalMiddlewares, middlewares...)
}

func (manager *Manager) With(handler http.Handler, middlewares ...Middleware) http.Handler {
	h := handler

	// Apply route-specific middleware (in reverse)
	for i := len(middlewares) - 1; i >= 0; i-- {
		h = middlewares[i](h)
	}

	return h
}

func (manager *Manager) WrapMux(handler http.Handler) http.Handler {
	h := handler

	// h = CorsWithPreflight(Logger(NothingGlobal(mux)))

	for i := len(manager.globalMiddlewares) - 1; i >= 0; i-- {
		h = manager.globalMiddlewares[i](h)
	}

	return h
}
