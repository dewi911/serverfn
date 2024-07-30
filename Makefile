.PHONY: build run test docker-build docker-compose-up docker-compose-down swag-init migrate-up migrate-down


GOBASE=$(shell pwd)
GOBIN=$(GOBASE)/bin
GOFILES=$(wildcard *.go)


BINARY_NAME=task_manager

build:
	@echo "Building..."
	@go build -o $(GOBIN)/$(BINARY_NAME) $(GOFILES)

run: build
	@echo "Running..."
	@$(GOBIN)/$(BINARY_NAME)

test:
	@echo "Testing..."
	@go test ./internal/... ./pkg/...

docker-build:
	@echo "Building Docker image..."
	@docker build -t $(BINARY_NAME) .

docker-compose-up:
	@echo "Starting Docker services..."
	@docker-compose up -d

docker-compose-down:
	@echo "Stopping Docker services..."
	@docker-compose down

swag-init:
	@echo "Initializing Swagger docs..."
	@swag init -g ./internal/app/app.go

migrate-up:
	@echo "Running database migrations..."
	@docker-compose run --rm migrate -path /migrations -database "postgres://${DB_USERNAME}:${DB_PASSWORD}@postgres:5432/${DB_NAME}?sslmode=disable" up

migrate-down:
	@echo "Reverting database migrations..."
	@docker-compose run --rm migrate -path /migrations -database "postgres://${DB_USERNAME}:${DB_PASSWORD}@postgres:5432/${DB_NAME}?sslmode=disable" down

clean:
	@echo "Cleaning..."
	@go clean
	@rm -f $(GOBIN)/$(BINARY_NAME)

env:
	@echo "Creating .env file..."
	@cp .env.example .env