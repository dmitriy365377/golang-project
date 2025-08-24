# Docker Setup для Golang Chat Application

## Обзор

Этот документ описывает, как настроить и использовать Docker для локальной разработки Golang Chat Application.

## Требования

- Docker Desktop
- Docker Compose
- Make (опционально, для удобства)

## Быстрый старт

### 1. Запуск всех сервисов

```bash
# Используя Makefile
make docker-up

# Или напрямую
docker-compose up -d
```

### 2. Проверка статуса

```bash
make docker-status
# или
docker-compose ps
```

### 3. Остановка сервисов

```bash
make docker-down
# или
docker-compose down
```

## Доступные сервисы

### PostgreSQL
- **Порт**: 5432
- **База данных**: chatdb
- **Пользователь**: chatuser
- **Пароль**: chatpass
- **Подключение**: `postgres://chatuser:chatpass@localhost:5432/chatdb?sslmode=disable`

### Redis (опционально)
- **Порт**: 6379
- **URL**: `redis://localhost:6379`

### pgAdmin (опционально)
- **Порт**: 5050
- **Email**: admin@chat.com
- **Пароль**: admin
- **URL**: http://localhost:5050

## Управление базой данных

### Подключение к PostgreSQL

```bash
make docker-psql
# или
docker exec -it golang-chat-postgres psql -U chatuser -d chatdb
```

### Создание резервной копии

```bash
make docker-backup
# Создает файл backup_YYYYMMDD_HHMMSS.sql
```

### Восстановление из резервной копии

```bash
make docker-restore FILE=backup_20240823_154000.sql
```

### Просмотр логов

```bash
make docker-logs
# или
docker-compose logs -f postgres
```

## Структура базы данных

После запуска контейнера PostgreSQL автоматически создаются:

- **users** - пользователи системы
- **chats** - чаты
- **chat_participants** - участники чатов
- **messages** - сообщения
- **refresh_tokens** - токены обновления
- **blacklisted_tokens** - черный список токенов

## Полезные команды

### Проверка статуса контейнеров
```bash
make docker-status
```

### Перезапуск сервисов
```bash
make docker-restart
```

### Очистка данных (осторожно!)
```bash
make docker-clean
# Удаляет все volumes и данные!
```

### Просмотр использования ресурсов
```bash
make docker-stats
```

## Troubleshooting

### Порт 5432 занят
```bash
# Остановить локальный PostgreSQL
brew services stop postgresql

# Или изменить порт в docker-compose.yml
ports:
  - "5433:5432"  # Внешний порт 5433
```

### Контейнер не запускается
```bash
# Проверить логи
make docker-logs

# Перезапустить
make docker-restart
```

### Проблемы с правами доступа
```bash
# Остановить и удалить volumes
make docker-clean

# Запустить заново
make docker-up
```

## Разработка

### Подключение к БД из Go

```go
import (
    "database/sql"
    _ "github.com/lib/pq"
)

func connectDB() (*sql.DB, error) {
    return sql.Open("postgres", "postgres://chatuser:chatpass@localhost:5432/chatdb?sslmode=disable")
}
```

### Переменные окружения

Создайте `.env` файл на основе `configs/env.example`:

```env
DATABASE_URL=postgres://chatuser:chatpass@localhost:5432/chatdb?sslmode=disable
JWT_SECRET=your-secret-key
AUTH_SERVICE_PORT=:8080
CHAT_SERVICE_PORT=:8081
```

## Production

⚠️ **Внимание**: Эта конфигурация предназначена только для разработки!

Для production:
- Измените пароли
- Настройте SSL
- Используйте внешние volumes
- Настройте backup стратегии
- Добавьте мониторинг

## Полезные ссылки

- [PostgreSQL Docker Hub](https://hub.docker.com/_/postgres)
- [Redis Docker Hub](https://hub.docker.com/_/redis)
- [pgAdmin Docker Hub](https://hub.docker.com/r/dpage/pgadmin4)
- [Docker Compose документация](https://docs.docker.com/compose/)
