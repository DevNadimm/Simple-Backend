package cmd

import (
	"fmt"
	"net/http"
	"test/middleware"
)

func Serve() {
	manager := middleware.NewManger()

	// GLOBAL Middlewares
	manager.Use(middleware.Logger, middleware.NothingGlobal)

	mux := http.NewServeMux()

	InitRoutes(mux, manager)

	fmt.Println("Server running at port: 3000")

	err := http.ListenAndServe(":3000", middleware.GlobalRouter(mux))
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
