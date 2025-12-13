package rest

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"test/config"
	"test/rest/handlers/product"
	"test/rest/handlers/user"
	"test/rest/middleware"
)

type Server struct {
	config         *config.Config
	userHandler    *user.Handler
	productHandler *product.Handler
}

func NewServer(
	config *config.Config,
	userHandler *user.Handler,
	productHandler *product.Handler,
) *Server {
	return &Server{
		config:         config,
		userHandler:    userHandler,
		productHandler: productHandler,
	}
}

func (server *Server) Start() {
	manager := middleware.NewManger()
	manager.Use(middleware.Cors, middleware.Preflight, middleware.Logger)

	mux := http.NewServeMux()
	wrappedMux := manager.WrapMux(mux)

	server.userHandler.RegisterRoutes(mux, manager)
	server.productHandler.RegisterRoutes(mux, manager)

	adress := ":" + strconv.Itoa(server.config.HttpPort)
	fmt.Println("Server running at port", adress)

	err := http.ListenAndServe(adress, wrappedMux)
	if err != nil {
		fmt.Println("Error starting server:", err)
		os.Exit(1)
	}
}
