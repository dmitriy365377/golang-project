#!/bin/bash

echo "Starting Golang Chat Application..."

# Проверяем, что бинарные файлы существуют
if [ ! -f "bin/auth-service" ] || [ ! -f "bin/chat-service" ] || [ ! -f "bin/chat-client" ]; then
    echo "Binary files not found. Please run ./scripts/build.sh first."
    exit 1
fi

# Запускаем Auth Service в фоне
echo "Starting Auth Service..."
./bin/auth-service &
AUTH_PID=$!

# Ждем немного для запуска Auth Service
sleep 2

# Запускаем Chat Service в фоне
echo "Starting Chat Service..."
./bin/chat-service &
CHAT_PID=$!

# Ждем немного для запуска Chat Service
sleep 2

# Запускаем Chat Client
echo "Starting Chat Client..."
./bin/chat-client

# Останавливаем сервисы при выходе
echo "Stopping services..."
kill $AUTH_PID $CHAT_PID 2>/dev/null
echo "All services stopped."
