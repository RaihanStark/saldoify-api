#!/bin/bash

# Development script for Saldoify API

set -e

echo "Starting Saldoify API in development mode..."

# Generate SQL code
echo "Generating SQL code..."
sqlc generate

# Run the application
echo "Starting server..."
go run cmd/api/main.go 