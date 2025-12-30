package cmd

import (
	"fmt"
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

	dbCon, err := db.NewConnection(config.DB)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	err = db.MigrateDB(dbCon, "./migrations")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	middleware := middleware.NewMiddleware(config)

	server := rest.NewServer(config, user.NewHandler(config, dbCon), product.NewHandler(middleware, dbCon))

	server.Start()
}
