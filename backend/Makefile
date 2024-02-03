# Simple Makefile for a Go project

# Build the application
all: build

build:
	@echo "Building..."
	
	@go build -o main cmd/api/main.go

# Run the application
run:
	@go run cmd/api/main.go

# Create DB container
docker-run:
	@if docker compose up 2>/dev/null; then \
		: ; \
	else \
		echo "Falling back to Docker Compose V1"; \
		docker-compose up; \
	fi

# Shutdown DB container
docker-down:
	@if docker compose down 2>/dev/null; then \
		: ; \
	else \
		echo "Falling back to Docker Compose V1"; \
		docker-compose down; \
	fi

# Test the application
test:
	@echo "Testing..."
	@go test ./tests -v

# Clean the binary
clean:
	@echo "Cleaning..."
	@rm -f main

doc:
	@if command - swag > /dev/null; then \
		echo "Swag is already installed on your machine. Skipping..."; \
	else \
		go install github.com/swaggo/swag/cmd/swag@latest; \
	fi
	@echo "Generating Swagger docs..."
	@swag fmt cmd/api/main.go
	@swag init -g cmd/api/main.go

# Live Reload
watch:
	@if command -v air > /dev/null; then \
	    air; \
	    echo "Watching...";\
	else \
	    read -p "Go's 'air' is not installed on your machine. Do you want to install it? [Y/n] " choice; \
	    if [ "$$choice" != "n" ] && [ "$$choice" != "N" ]; then \
	        go install github.com/cosmtrek/air@latest; \
	        air; \
	        echo "Watching...";\
	    else \
	        echo "You chose not to install air. Exiting..."; \
	        exit 1; \
	    fi; \
	fi

.PHONY: all build run test clean doc

CONNECTION_STRING=postgres://${DB_USERNAME}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_DATABASE}?sslmode=disable
migrate-check:
	@if ! command -v migrate &> /dev/null; then \
        go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest; \
    fi

migrate-up:
	@migrate -path db/migrations -database "$(CONNECTION_STRING)" up

migrate-down:
	@migrate -path db/migrations -database "$(CONNECTION_STRING)" down

migrate-create:
	@if [ -z "$(NAME)" ]; then \
		read -p "Enter the name of the migration: " name; \
	else \
		name=$(NAME); \
	fi; \
	migrate create -ext sql -dir db/migrations -seq $$name

migrate-version:
	@migrate -path db/migrations -database "$(CONNECTION_STRING)" version

migrate-force:
	@migrate -path db/migrations -database "$(CONNECTION_STRING)" force $(VERSION)

.PHONY: migrate-check migrate-up migrate-down migrate-create migrate-version migrate-force



