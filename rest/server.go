package rest

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"test/config"
	"test/rest/middleware"
)

func Start(config config.Config) {
	manager := middleware.NewManger()

	// GLOBAL Middlewares
	manager.Use(middleware.Cors, middleware.Preflight, middleware.Logger, middleware.NothingGlobal)

	mux := http.NewServeMux()
	wrappedMux := manager.WrapMux(mux)
	initRoutes(mux, manager)

	adress := ":" + strconv.Itoa(config.HttpPort)

	fmt.Println("Server running at port", adress)
	err := http.ListenAndServe(adress, wrappedMux)
	if err != nil {
		fmt.Println("Error starting server:", err)
		os.Exit(1)
	}
}
