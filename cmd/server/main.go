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

	// Allow CORS for all origins
	corsMiddleware := func(handler fasthttp.RequestHandler) fasthttp.RequestHandler {
		return func(ctx *fasthttp.RequestCtx) {
			ctx.Response.Header.Set("Access-Control-Allow-Origin", "*")
			handler(ctx)
		}
	}

	// Health check route
	r.GET("/health", func(ctx *fasthttp.RequestCtx) {
		ctx.SetStatusCode(fasthttp.StatusOK)
	})

	// Register routes
	r.GET("/users", userHandler.HandleGetAllUsers)
	r.GET("/users/{id}", userHandler.HandleGetUser)
	r.POST("/users", userHandler.HandleCreateUser)
	r.PUT("/users/{id}", userHandler.HandleUpdateUser)
	r.DELETE("/users/{id}", userHandler.HandleDeleteUser)

	// Start server
	addr := fmt.Sprintf(":%d", *port)
	log.Printf("Server starting on %s", addr)
	if err := fasthttp.ListenAndServe(addr, corsMiddleware(r.Handler)); err != nil {
		log.Fatalf("Error starting server: %s", err)
	}
}
