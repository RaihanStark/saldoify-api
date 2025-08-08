package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
	sqlc "github.com/raihansoftware/saldoify-api/internal/pkg/database/sqlc"
)

// Handler represents the API handlers
type Handler struct {
	queries *sqlc.Queries
}

// NewHandler creates a new handler instance
func NewHandler(queries *sqlc.Queries) *Handler {
	return &Handler{
		queries: queries,
	}
}

// Welcome handles the welcome endpoint
func (h *Handler) Welcome(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{
		"message": "Welcome to Saldoify API",
	})
}

// Health handles the health check endpoint
func (h *Handler) Health(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{
		"status": "healthy",
	})
}

// ListUsers handles the get users endpoint
func (h *Handler) ListUsers(c echo.Context) error {
	users, err := h.queries.ListUsers(c.Request().Context())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to fetch users",
		})
	}
	return c.JSON(http.StatusOK, users)
}
