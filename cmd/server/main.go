package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"go.uber.org/zap"
	"user-management-api/internal/config"
	"user-management-api/internal/database"
	"user-management-api/internal/handlers"
	"user-management-api/internal/middleware"
	"user-management-api/internal/repository"
	"user-management-api/internal/service"
	"user-management-api/pkg/logger"
)

// @title User Management API
// @version 1.0
// @description RESTful API for managing users with dynamic age calculation
// @host localhost:8080
// @BasePath /api/v1

func main() {
	// Initialize logger
	if err := logger.InitLogger(); err != nil {
		log.Fatalf("Failed to initialize logger: %v", err)
	}
	defer logger.Sync()

	// Load configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		logger.Log.Fatal("Failed to load configuration", zap.Error(err))
	}

	// Initialize database
	db, err := database.NewPostgresDB(cfg)
	if err != nil {
		logger.Log.Fatal("Failed to connect to database", zap.Error(err))
	}
	defer db.Close()

	// Test database connection
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := db.PingContext(ctx); err != nil {
		logger.Log.Fatal("Database ping failed", zap.Error(err))
	}

	logger.Log.Info("Database connection established")

	// Initialize repository
	userRepo := repository.NewUserRepository(db)

	// Initialize service
	userService := service.NewUserService(userRepo)

	// Initialize handler
	userHandler := handlers.NewUserHandler(userService)

	// Initialize Fiber app
	app := fiber.New(fiber.Config{
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			logger.Log.Error("Unhandled error", zap.Error(err))
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"success": false,
				"message": "Internal server error",
			})
		},
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	})

	// Middleware
	app.Use(recover.New())
	app.Use(middleware.Headers())
	app.Use(middleware.Logger())
	app.Use(middleware.RequestID())

	// Routes
	api := app.Group("/api/v1")

	api.Post("/users", userHandler.CreateUser)
	api.Get("/users", userHandler.ListUsers)
	api.Get("/users/:id", userHandler.GetUserByID)
	api.Put("/users/:id", userHandler.UpdateUser)
	api.Delete("/users/:id", userHandler.DeleteUser)

	// Health check
	app.Get("/health", func(c *fiber.Ctx) error {
		ctx, cancel := context.WithTimeout(c.Context(), 2*time.Second)
		defer cancel()

		if err := db.PingContext(ctx); err != nil {
			return c.Status(fiber.StatusServiceUnavailable).JSON(fiber.Map{
				"status":   "unhealthy",
				"database": "disconnected",
				"error":    err.Error(),
			})
		}
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"status":    "healthy",
			"database":  "connected",
			"timestamp": time.Now().Format(time.RFC3339),
		})
	})

	// Start server
	port := cfg.ServerPort
	logger.Log.Info("Starting server", zap.Int("port", port))

	// Graceful shutdown
	go func() {
		if err := app.Listen(fmt.Sprintf(":%d", port)); err != nil {
			logger.Log.Fatal("Failed to start server", zap.Error(err))
		}
	}()

	// Wait for interrupt signal
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	logger.Log.Info("Shutting down server...")

	shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer shutdownCancel()

	if err := app.Shutdown(); err != nil {
		logger.Log.Error("Server shutdown error", zap.Error(err))
	}

	<-shutdownCtx.Done()
	logger.Log.Info("Server shutdown complete")
}