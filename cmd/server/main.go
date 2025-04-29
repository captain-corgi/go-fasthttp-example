package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/captain-corgi/go-fasthttp-example/internal/domain/repository"
	"github.com/captain-corgi/go-fasthttp-example/internal/handler"
	"github.com/captain-corgi/go-fasthttp-example/internal/service"
	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
)

func main() {
	// Parse command line flags
	port := flag.Int("port", 8080, "Port to listen on")
	flag.Parse()

	// Initialize dependencies
	userRepo := repository.NewInMemoryUserRepository()
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	// Create router
	r := router.New()

	// Register routes
	r.GET("/users/{id}", userHandler.HandleGetUser)
	r.POST("/users", userHandler.HandleCreateUser)
	r.PUT("/users/{id}", userHandler.HandleUpdateUser)
	r.DELETE("/users/{id}", userHandler.HandleDeleteUser)

	// Start server
	addr := fmt.Sprintf(":%d", *port)
	log.Printf("Server starting on %s", addr)
	if err := fasthttp.ListenAndServe(addr, r.Handler); err != nil {
		log.Fatalf("Error starting server: %s", err)
	}
}
