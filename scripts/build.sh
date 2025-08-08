#!/bin/bash

# Build script for Saldoify API

set -e

echo "Building Saldoify API..."

# Generate SQL code
echo "Generating SQL code..."
sqlc generate

# Build the application
echo "Building application..."
go build -o bin/saldoify-api cmd/api/main.go

echo "Build completed successfully!"
echo "Binary location: bin/saldoify-api" 