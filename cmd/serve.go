package cmd

import (
	"fmt"
	"net/http"
	"test/middleware"
)

func Serve() {
	manager := middleware.NewManger()

	// GLOBAL Middlewares
	manager.Use(middleware.Cors, middleware.Preflight, middleware.Logger, middleware.NothingGlobal)

	fmt.Println("Server running at port: 3000")

	mux := http.NewServeMux()
	wrappedMux := manager.WrapMux(mux)
	InitRoutes(mux, manager)

	err := http.ListenAndServe(":3000", wrappedMux)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
