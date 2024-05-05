package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/afonso-borges/mailflow/api"
	"github.com/afonso-borges/mailflow/internal/config"
	"github.com/afonso-borges/mailflow/internal/email"
	"github.com/afonso-borges/mailflow/internal/queue"
	"github.com/afonso-borges/mailflow/internal/templates"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Println("Error loading .env file")
	}

	// Initialize configuration
	cfg := config.New()

	// Initialize templates
	tmpl, err := templates.New()
	if err != nil {
		log.Fatalf("Error initializing templates: %v", err)
	}

	// Initialize Redis client
	redisClient, err := queue.NewRedisClient(cfg)
	if err != nil {
		log.Fatalf("Error connecting to Redis: %v", err)
	}
	defer redisClient.Close()

	// Initialize email sender
	emailService := email.NewSender(cfg, tmpl)

	// Start worker to process emails from queue
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go queue.StartWorker(ctx, redisClient, emailService)

	// Initialize HTTP server
	router := gin.Default()
	api.RegisterHandlers(router, redisClient)

	// Start HTTP server
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", cfg.Port),
		Handler: router,
	}

	// Run server
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Error starting server: %v", err)
		}
	}()

	log.Printf("Server started on port %s", cfg.Port)

	// Graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Error shutting down server: %v", err)
	}

	log.Println("Server shut down successfully")
}
