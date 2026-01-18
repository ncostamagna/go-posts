# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Build and Development Commands

This project uses [Taskfile](https://taskfile.dev/) for task automation. All tasks require a `.env` file with `DB_DNS` for database connection.

```bash
task dev          # Start development server with hot reload (uses Air)
task test         # Run tests and golangci-lint
task clean        # Run go mod tidy, go fmt, and goimports
task generate     # Regenerate sqlc database code from SQL queries
task migrate name=<name>  # Create a new migration file
task migrate-up   # Apply all pending migrations
task migrate-down # Rollback the last migration
```

**Required tools:** taskfile, air, migrate, sqlc, golangci-lint, goimports

## Architecture

This is a Go microservice for managing posts, following a layered architecture with go-kit patterns.

### Layer Structure

- **cmd/main.go** - Application entry point, wires dependencies and starts HTTP server
- **transport/http/posts/** - HTTP layer using Gin + go-kit transport
  - `router.go` - Route definitions and request/response decoders
  - `controller.go` - Endpoint factories connecting HTTP to service layer
- **internal/posts/** - Business logic layer
  - `service.go` - Core business logic (Service interface)
  - `repository.go` - Data access abstraction (Repository interface)
  - `instrumenting.go` - Prometheus metrics wrapper (decorator pattern)
- **adapters/database/** - sqlc-generated database code
  - `queries/*.sql` - SQL queries with sqlc annotations
  - Generated Go files from `task generate`
- **pkg/instance/** - Service instantiation and dependency wiring
- **pkg/log/** - Structured logging wrapper (slog)

### Key Patterns

**Dependency Flow:** `main.go` → `pkg/instance` → creates Repository → creates Service → wraps with Instrumenting → creates Endpoints → creates HTTP router

**sqlc Code Generation:** SQL queries in `adapters/database/queries/*.sql` generate Go code. After modifying queries, run `task generate`.

**Metrics:** All service methods are wrapped with Prometheus instrumentation via the decorator in `instrumenting.go`. Metrics available at `/metrics`.

### Environment Variables

- `DB_DNS` - PostgreSQL connection string
- `APP_URL` - Server address (default: `0.0.0.0:80`)
- `LOG_LEVEL` - Logging level
