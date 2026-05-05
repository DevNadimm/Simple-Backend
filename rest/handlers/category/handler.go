package category

import (
	"test/repo"
	"test/rest/middleware"

	"github.com/jmoiron/sqlx"
)

type Handler struct {
	middleware   *middleware.Middleware
	categoryRepo repo.CategoryRepo
}

func NewHandler(middleware *middleware.Middleware, db *sqlx.DB) *Handler {
	return &Handler{
		middleware:   middleware,
		categoryRepo: repo.NewCategoryRepo(db),
	}
}
