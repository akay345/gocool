package main

import (
	"log"
	"net/http"
	"time"

	"gocool/internal/config"
	"gocool/internal/middleware"
	"gocool/internal/router"
)

func main() {
	// Load configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	// Create a new router
	r := router.SetupRoutes()

	// Apply middleware
	r.Use(middleware.Tracer())
	r.Use(middleware.Auth())

	// Define the server configuration
	srv := &http.Server{
		Handler:      r,
		Addr:         ":" + cfg.ServerPort,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	// Start the server
	log.Println("Starting server on port", cfg.Server.Port)
	if err := srv.ListenAndServe(); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
