package cmd

import (
	"test/config"
	"test/rest"
	"test/rest/handlers/product"
	"test/rest/handlers/user"
	"test/rest/middleware"
)

func Serve() {
	config := config.GetConfig()
	middleware := middleware.NewMiddleware(config)

	server := rest.NewServer(config, user.NewHandler(), product.NewHandler(middleware))
	
	server.Start()
}
