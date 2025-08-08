package api

import (
	"github.com/labstack/echo/v4"
	sqlc "github.com/raihansoftware/saldoify-api/internal/pkg/database/sqlc"
)

// SetupRoutes configures all API routes
func SetupRoutes(e *echo.Echo, queries *sqlc.Queries) {
	handler := NewHandler(queries)

	// Welcome route
	e.GET("/", handler.Welcome)

	// Health check route
	e.GET("/health", handler.Health)

	// Users routes
	setupUserRoutes(e, handler)
}

// setupUserRoutes configures user-related routes
func setupUserRoutes(e *echo.Echo, handler *Handler) {
	// Get all users
	e.GET("/users", handler.ListUsers)
}
