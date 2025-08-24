# Быстрый старт

## Предварительные требования

- Go 1.21 или выше
- Protocol Buffers compiler (protoc)
- Git

## Установка зависимостей

```bash
# Клонирование репозитория
git clone <repository-url>
cd golang-chat

# Установка Go зависимостей
go mod tidy
```

## Конфигурация

1. Скопируйте файл конфигурации:
```bash
cp configs/env.example .env
```

2. Отредактируйте `.env` файл под ваши нужды

## Сборка и запуск

### Вариант 1: Использование Makefile (рекомендуется)

```bash
# Показать все доступные команды
make help

# Собрать все бинарные файлы
make build

# Запустить все сервисы
make run

# Очистить бинарные файлы
make clean
```

### Вариант 2: Использование скриптов

```bash
# Сборка
./scripts/build.sh

# Запуск
./scripts/run.sh
```

### Вариант 3: Ручной запуск

```bash
# Терминал 1: Auth Service
go run cmd/auth-service/main.go

# Терминал 2: Chat Service  
go run cmd/chat-service/main.go

# Терминал 3: Chat Client
go run cmd/chat-client/main.go
```

## Использование CLI клиента

После запуска клиента доступны следующие команды:

```bash
> login username password
> create chat "My Chat"
> connect chat_id_here
> send chat_id_here "Hello, world!"
> quit
```

## Проверка работы

1. Auth Service должен запуститься на порту 50051
2. Chat Service должен запуститься на порту 50052
3. CLI клиент должен показать приветственное сообщение

## Устранение неполадок

- **Port already in use:** Измените порты в `.env` файле
- **Import errors:** Убедитесь, что выполнили `go mod tidy`
- **Permission denied:** Убедитесь, что скрипты исполняемые (`chmod +x scripts/*.sh`)
