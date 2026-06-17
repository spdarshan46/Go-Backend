.PHONY: run build test docker-up docker-down migrate swagger sqlc clean lint

run:
	go run cmd/server/main.go

build:
	go build -o bin/server cmd/server/main.go

test:
	go test -v -race -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o coverage.html

docker-up:
	docker-compose up -d

docker-down:
	docker-compose down

migrate:
	psql -h localhost -U postgres -d userdb -f internal/database/migrations/001_create_users_table.up.sql

sqlc:
	sqlc generate

swagger:
	swag init -g cmd/server/main.go -o docs

lint:
	golangci-lint run ./...

tidy:
	go mod tidy

clean:
	rm -rf bin/
	rm -f coverage.out coverage.html
	rm -rf docs/