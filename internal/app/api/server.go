package api

import (
	"log"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	sqlc "github.com/raihansoftware/saldoify-api/internal/pkg/database/sqlc"
)

// Server represents the API server
type Server struct {
	echo    *echo.Echo
	queries *sqlc.Queries
}

// NewServer creates a new server instance
func NewServer(queries *sqlc.Queries) *Server {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	return &Server{
		echo:    e,
		queries: queries,
	}
}

// Setup configures the server with routes and middleware
func (s *Server) Setup() {
	SetupRoutes(s.echo, s.queries)
}

// Start starts the server
func (s *Server) Start() error {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Starting server on port %s", port)
	return s.echo.Start(":" + port)
}

// GetEcho returns the Echo instance (useful for testing)
func (s *Server) GetEcho() *echo.Echo {
	return s.echo
}
