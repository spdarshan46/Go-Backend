# User Management API

A RESTful API built using GoFiber, PostgreSQL, SQLC, Uber Zap Logger, and Validator.

## Features

* Create User
* Get User By ID
* Update User
* Delete User
* List All Users
* Dynamic Age Calculation
* Input Validation
* PostgreSQL Database
* SQLC Generated Queries
* Request Logging using Zap

---

## Tech Stack

* Go
* GoFiber
* PostgreSQL
* SQLC
* Uber Zap
* go-playground/validator

---

## Project Structure

```text
cmd/
└── server/
    └── main.go

internal/
├── config
├── database
├── handlers
├── middleware
├── models
├── repository
├── service
└── utils

pkg/
└── logger
```

---

## Database Schema

```sql
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    dob DATE NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
```

---

## Setup

### 1. Clone Repository

```bash
git clone https://github.com/spdarshan46/Go-Backend.git
cd Go-Backend
```

### 2. Install Dependencies

```bash
go mod tidy
```

### 3. Create PostgreSQL Database

```sql
CREATE DATABASE user_management;
```

### 4. Create Users Table

```sql
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    dob DATE NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
```

### 5. Configure Environment Variables

Create a `.env` file:

```env
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=your_password
DB_NAME=user_management
SERVER_PORT=8080
```

### 6. Run Application

```bash
go run cmd/server/main.go
```

Server starts at:

```text
http://localhost:8080
```

---

## API Endpoints

### Create User

POST `/api/v1/users`

Request

```json
{
  "name": "Alice",
  "dob": "1990-05-10"
}
```

---

### Get User

GET `/api/v1/users/{id}`

---

### Update User

PUT `/api/v1/users/{id}`

```json
{
  "name": "Alice Updated",
  "dob": "1991-03-15"
}
```

---

### Delete User

DELETE `/api/v1/users/{id}`

Returns:

```text
204 No Content
```

---

### List Users

GET `/api/v1/users`

---

## Health Check

GET `/health`

---

## Run Tests

```bash
go test ./...
```

---

## Author

Darshan S P

GitHub: https://github.com/spdarshan46
