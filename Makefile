.PHONY: build run clean proto test help

# Переменные
BIN_DIR=bin
SCRIPTS_DIR=scripts

# Цели по умолчанию
help: ## Показать справку
	@echo "Доступные команды:"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'

build: ## Собрать все бинарные файлы
	@echo "Building Golang Chat Application..."
	@mkdir -p $(BIN_DIR)
	go build -o $(BIN_DIR)/auth-service cmd/auth-service/main.go
	go build -o $(BIN_DIR)/chat-service cmd/chat-service/main.go
	go build -o $(BIN_DIR)/chat-client cmd/chat-client/main.go
	@echo "Build completed!"

run: build ## Собрать и запустить все сервисы
	@echo "Starting all services..."
	@$(SCRIPTS_DIR)/run.sh

clean: ## Очистить бинарные файлы
	@echo "Cleaning binary files..."
	@rm -rf $(BIN_DIR)
	@echo "Clean completed!"

proto: ## Сгенерировать Go код из Protobuf файлов
	@echo "Generating Go code from Protobuf..."
	@export PATH=$$PATH:$$(go env GOPATH)/bin && protoc --go_out=. --go_opt=paths=source_relative \
		--go-grpc_out=. --go-grpc_opt=paths=source_relative \
		proto/auth/auth.proto proto/chat/chat.proto
	@echo "Protobuf generation completed!"

deps: ## Установить зависимости
	@echo "Installing dependencies..."
	go mod tidy
	go mod download
	@echo "Dependencies installed!"

test: ## Запустить тесты
	@echo "Running tests..."
	go test ./...
	@echo "Tests completed!"

dev: ## Запустить в режиме разработки
	@echo "Starting in development mode..."
	@go run cmd/auth-service/main.go &
	@go run cmd/chat-service/main.go &
	@go run cmd/chat-client/main.go
	@echo "Development mode completed!"

# Docker commands
.PHONY: docker-up docker-down docker-logs docker-clean docker-restart

# Запуск всех сервисов
docker-up:
	docker-compose up -d

# Остановка всех сервисов
docker-down:
	docker-compose down

# Просмотр логов
docker-logs:
	docker-compose logs -f

# Очистка данных (удаление volumes)
docker-clean:
	docker-compose down -v
	docker system prune -f

# Перезапуск сервисов
docker-restart:
	docker-compose restart

# Запуск только PostgreSQL
docker-postgres:
	docker-compose up -d postgres

# Запуск только Redis
docker-redis:
	docker-compose up -d redis

# Подключение к PostgreSQL
docker-psql:
	docker exec -it golang-chat-postgres psql -U chatuser -d chatdb

# Быстрое подключение к БД через скрипт
db-connect:
	./scripts/connect-db.sh

# Создание резервной копии БД
docker-backup:
	docker exec golang-chat-postgres pg_dump -U chatuser chatdb > backup_$(shell date +%Y%m%d_%H%M%S).sql

# Восстановление БД из резервной копии
docker-restore:
	@echo "Usage: make docker-restore FILE=backup_file.sql"
	@if [ -z "$(FILE)" ]; then echo "Please specify FILE parameter"; exit 1; fi
	docker exec -i golang-chat-postgres psql -U chatuser -d chatdb < $(FILE)

# Проверка статуса контейнеров
docker-status:
	docker-compose ps

# Просмотр использования ресурсов
docker-stats:
	docker stats golang-chat-postgres golang-chat-redis golang-chat-pgadmin
