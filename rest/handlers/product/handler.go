package product

import (
	"test/repo"
	"test/rest/middleware"

	"github.com/jmoiron/sqlx"
)

type Handler struct {
	middleware  *middleware.Middleware
	productRepo repo.ProductRepo
}

func NewHandler(middleware *middleware.Middleware, db *sqlx.DB) *Handler {
	return &Handler{
		middleware:  middleware,
		productRepo: repo.NewProductRepo(db),
	}
}
