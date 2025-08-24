.PHONY: build clean build-all

# Build all services
build-all: build-auth build-chat build-rest-auth

# Build individual services
build-auth:
	@echo "Building auth-service..."
	@go build -o bin/auth-service ./cmd/auth-service

build-chat:
	@echo "Building chat-service..."
	@go build -o bin/chat-service ./cmd/chat-service

build-rest-auth:
	@echo "Building rest-auth-service..."
	@go build -o bin/rest-auth-service ./cmd/rest-auth-service

# Build chat client
build-client:
	@echo "Building chat-client..."
	@go build -o bin/chat-client ./cmd/chat-client

# Clean build artifacts
clean:
	@echo "Cleaning build artifacts..."
	@rm -rf bin/
	@go clean

# Install dependencies
deps:
	@echo "Installing dependencies..."
	@go mod download
	@go mod tidy

# Run tests
test:
	@echo "Running tests..."
	@go test ./...

# Run with race detection
test-race:
	@echo "Running tests with race detection..."
	@go test -race ./...

# Format code
fmt:
	@echo "Formatting code..."
	@go fmt ./...

# Lint code
lint:
	@echo "Linting code..."
	@golangci-lint run

# Help
help:
	@echo "Available targets:"
	@echo "  build-all     - Build all services"
	@echo "  build-auth    - Build auth service"
	@echo "  build-chat    - Build chat service"
	@echo "  build-rest-auth - Build REST auth service"
	@echo "  build-client  - Build chat client"
	@echo "  clean         - Clean build artifacts"
	@echo "  deps          - Install dependencies"
	@echo "  test          - Run tests"
	@echo "  test-race     - Run tests with race detection"
	@echo "  fmt           - Format code"
	@echo "  lint          - Lint code"
	@echo "  help          - Show this help"
