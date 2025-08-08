.PHONY: build dev test clean generate migrate

# Build the application
build:
	@echo "Building Saldoify API..."
	@./scripts/build.sh

# Run in development mode
dev:
	@echo "Starting development server..."
	@./scripts/dev.sh

# Generate SQL code
generate:
	@echo "Generating SQL code..."
	@sqlc generate

# Run database migrations
migrate:
	@echo "Running database migrations..."
	@dbmate up

# Run tests
test:
	@echo "Running tests..."
	@go test ./...

# Clean build artifacts
clean:
	@echo "Cleaning build artifacts..."
	@rm -rf bin/
	@go clean

# Install dependencies
deps:
	@echo "Installing dependencies..."
	@go mod download

# Format code
fmt:
	@echo "Formatting code..."
	@go fmt ./...

# Vet code
vet:
	@echo "Vetting code..."
	@go vet ./...

# All-in-one setup
setup: deps generate migrate
	@echo "Setup completed!" 