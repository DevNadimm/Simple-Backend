package middleware

import "test/config"

type Middleware struct {
	config *config.Config
}

func NewMiddleware(config *config.Config) *Middleware {
	return &Middleware{
		config: config,
	}
}
