# Golang Chat Application

Микросервисное приложение чата, построенное на Go с использованием gRPC.

## Архитектура

Проект состоит из трех основных компонентов:

- **Auth Service** - сервис аутентификации и авторизации
- **Chat Service** - сервис чата
- **Chat Client** - CLI клиент для взаимодействия с сервисами

## Структура проекта

```
golang-chat/
├── cmd/                    # Точки входа приложений
├── internal/               # Внутренняя логика
├── pkg/                    # Переиспользуемые пакеты
├── proto/                  # Protobuf определения
├── configs/                # Конфигурационные файлы
├── scripts/                # Скрипты для сборки и развертывания
└── docs/                   # Документация
```

## Запуск

1. Установите зависимости: `go mod tidy`
2. Запустите Auth Service: `go run cmd/auth-service/main.go`
3. Запустите Chat Service: `go run cmd/chat-service/main.go`
4. Запустите Chat Client: `go run cmd/chat-client/main.go`

## API

- Auth Service: gRPC на порту 50051
- Chat Service: gRPC на порту 50052
- Chat Client: CLI интерфейс
