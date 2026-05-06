package middleware

import (
	"test/config"
	"test/repo"
)

type Middleware struct {
	config   *config.Config
	userRepo repo.UserRepo
}

func NewMiddleware(config *config.Config, userRepo repo.UserRepo) *Middleware {
	return &Middleware{
		config:   config,
		userRepo: userRepo,
	}
}
