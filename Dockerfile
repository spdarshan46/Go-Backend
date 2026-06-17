FROM golang:1.21-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Install SQLC and generate code
RUN go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest && sqlc generate

# Install swag and generate docs
RUN go install github.com/swaggo/swag/cmd/swag@latest && swag init -g cmd/server/main.go -o docs

RUN go build -o main ./cmd/server

FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=builder /app/main .
COPY --from=builder /app/.env .env
COPY --from=builder /app/docs ./docs

EXPOSE 8080

CMD ["./main"]