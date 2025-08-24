# 🔐 REST Auth Service

REST API сервис для авторизации пользователей с использованием JWT токенов.

## 📁 Структура проекта

```
cmd/rest-auth-service/
├── main.go                    # Точка входа (с TODO заданиями)
└── README.md                  # Этот файл

internal/rest-auth/
├── model/
│   └── user.go               # Модели пользователей (с TODO)
├── service/
│   └── auth_service.go       # Бизнес-логика авторизации (с TODO)
├── handler/
│   └── auth_handler.go       # HTTP обработчики (с TODO)
└── middleware/
    └── auth_middleware.go    # Middleware для безопасности (с TODO)

docs/
└── rest-auth-tasks.md        # Подробные задания по реализации
```

## 🎯 Ваша задача

Реализовать полнофункциональный REST API для авторизации, следуя TODO комментариям в каждом файле.

## 🚀 Как начать

1. **Изучите технологии:**
   - JWT токены
   - bcrypt для хеширования паролей
   - HTTP middleware
   - Валидация данных

2. **Реализуйте код пошагово:**
   - Начните с моделей данных
   - Затем сервис
   - Потом handlers
   - В конце middleware

3. **Протестируйте API:**
   - Используйте Postman или curl
   - Проверьте все endpoints
   - Убедитесь в безопасности

## 📚 Документация

- **Подробные задания:** `docs/rest-auth-tasks.md`
- **Архитектура:** `docs/architecture.md`
- **Быстрый старт:** `docs/quickstart.md`

## 🔧 Зависимости

```bash
go get github.com/gorilla/mux
go get github.com/golang-jwt/jwt/v5
go get golang.org/x/crypto/bcrypt
go get github.com/go-playground/validator/v10
```

## 🌐 API Endpoints

### Публичные
- `POST /api/auth/register` - Регистрация
- `POST /api/auth/login` - Вход
- `POST /api/auth/refresh` - Обновление токена

### Защищенные
- `GET /api/auth/profile` - Профиль
- `PUT /api/auth/profile` - Обновление профиля
- `POST /api/auth/logout` - Выход

### Админские
- `GET /api/users` - Список пользователей
- `GET /api/users/{id}` - Информация о пользователе
- `PUT /api/users/{id}` - Обновление пользователя
- `DELETE /api/users/{id}` - Удаление пользователя

## 🎯 Готово к разработке!

Все файлы созданы с подробными TODO комментариями. Начинайте реализацию! 🚀
