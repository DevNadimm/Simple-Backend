package cmd

import (
	"os"
	"test/config"
	"test/infra/db"
	"test/rest"
	"test/rest/handlers/product"
	"test/rest/handlers/user"
	"test/rest/middleware"
)

func Serve() {
	config := config.GetConfig()
	middleware := middleware.NewMiddleware(config)
	db, err := db.NewConnection()

	if err != nil {
		os.Exit(1)
	}

	server := rest.NewServer(config, user.NewHandler(config, db), product.NewHandler(middleware, db))

	server.Start()
}
