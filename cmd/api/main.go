package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	dbconn "github.com/raihansoftware/saldoify-api/internal/pkg/database"
	sqlc "github.com/raihansoftware/saldoify-api/internal/pkg/database/sqlc"
)

func main() {
	// Load environment variables
	if err := godotenv.Load("configs/config.env"); err != nil {
		log.Printf("Warning: Could not load configs/config.env file: %v", err)
	}

	// Connect to database
	database, err := dbconn.NewConnection()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer database.Close()

	// Create Queries instance
	queries := sqlc.New(database)

	// Create Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	// Routes
	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{
			"message": "Welcome to Saldoify API",
		})
	})

	e.GET("/health", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{
			"status": "healthy",
		})
	})

	// Test database connection
	e.GET("/users", func(c echo.Context) error {
		users, err := queries.ListUsers(c.Request().Context())
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{
				"error": "Failed to fetch users",
			})
		}
		return c.JSON(http.StatusOK, users)
	})

	// Start server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	e.Logger.Fatal(e.Start(":" + port))
}
