package user

import (
	"test/config"
	"test/repo"

	"github.com/jmoiron/sqlx"
)

type Handler struct {
	config   *config.Config
	userRepo repo.UserRepo
}

func NewHandler(config *config.Config, db *sqlx.DB) *Handler {
	return &Handler{
		config:   config,
		userRepo: repo.NewUserRepo(db),
	}
}
