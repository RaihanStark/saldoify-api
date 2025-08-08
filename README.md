# Saldoify API

A Go-based API service for Saldoify application.

## Project Structure

This project follows the [Go project layout standards](https://github.com/golang-standards/project-layout).

### Key Directories

- `/cmd/api` - Main application entry point
- `/internal/pkg/database` - Database-related code (queries, migrations, schema)
- `/internal/app/api` - Application-specific API code
- `/configs` - Configuration files
- `/deployments` - Deployment configurations (Docker, etc.)
- `/docs` - Documentation
- `/api` - API specifications (OpenAPI/Swagger)
- `/build` - Build and CI configurations
- `/scripts` - Build and utility scripts

## Setup

1. Install dependencies:
   ```bash
   go mod download
   ```

2. Set up the database:
   ```bash
   # Run migrations
   dbmate up
   ```

3. Generate SQL code:
   ```bash
   sqlc generate
   ```

4. Run the application:
   ```bash
   go run cmd/api/main.go
   ```

## Environment Variables

Copy `configs/config.env.example` to `configs/config.env` and configure your environment variables.

## API Endpoints

- `GET /` - Welcome message
- `GET /health` - Health check
- `GET /users` - List users

## Development

- Database migrations are in `internal/pkg/database/migrations/`
- SQL queries are in `internal/pkg/database/queries/`
- Generated SQL code is in `internal/pkg/database/sqlc/` 