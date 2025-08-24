#!/bin/bash

echo "Building Golang Chat Application..."

# Создаем директорию для бинарных файлов
mkdir -p bin

# Собираем Auth Service
echo "Building Auth Service..."
go build -o bin/auth-service cmd/auth-service/main.go

# Собираем Chat Service
echo "Building Chat Service..."
go build -o bin/chat-service cmd/chat-service/main.go

# Собираем Chat Client
echo "Building Chat Client..."
go build -o bin/chat-client cmd/chat-client/main.go

echo "Build completed! Binaries are in the bin/ directory."
echo ""
echo "To run the services:"
echo "  ./bin/auth-service"
echo "  ./bin/chat-service"
echo "  ./bin/chat-client"
