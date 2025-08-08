package main

import (
	"log"

	"github.com/joho/godotenv"
	api "github.com/raihansoftware/saldoify-api/internal/app/api"
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

	// Create and setup server
	server := api.NewServer(queries)
	server.Setup()

	// Start server
	if err := server.Start(); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
