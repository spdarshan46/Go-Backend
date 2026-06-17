# User Management API

A production-ready RESTful API built with Go Fiber for managing users with dynamic age calculation.

## Features

- ✅ CRUD operations with dynamic age calculation
- ✅ SQLC for type-safe database access
- ✅ Input validation with go-playground/validator
- ✅ Structured logging with Uber Zap
- ✅ Docker support with Docker Compose
- ✅ Pagination support
- ✅ Request ID middleware
- ✅ Health check with database verification
- ✅ Unit tests for age calculation
- ✅ Swagger/OpenAPI documentation
- ✅ Postman collection
- ✅ CI/CD with GitHub Actions
- ✅ Graceful shutdown

## Quick Start

### Prerequisites
- Go 1.21+
- PostgreSQL 15+
- Docker & Docker Compose (optional)
- SQLC (for code generation)

### Installation

1. Clone the repository:
```bash
git clone <repository-url>
cd user-management-api