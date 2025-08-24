# 🚀 Задания по реализации REST авторизации

## 📋 **Обзор проекта**

Вы создаете REST API для авторизации пользователей с использованием JWT токенов, ролевой системой и middleware для безопасности.

## 🎯 **Что нужно реализовать**

### **1. Основная структура (main.go)**
- ✅ Создан файл с TODO комментариями
- 🔄 **Ваша задача:** Реализовать все TODO пункты

### **2. Модели данных (model/user.go)**
- ✅ Созданы структуры с TODO комментариями
- 🔄 **Ваша задача:** Дополнить модели согласно комментариям

### **3. Бизнес-логика (service/auth_service.go)**
- ✅ Создан сервис с подробными TODO
- 🔄 **Ваша задача:** Реализовать всю логику авторизации

### **4. HTTP обработчики (handler/auth_handler.go)**
- ✅ Созданы handlers с TODO комментариями
- 🔄 **Ваша задача:** Реализовать все HTTP endpoints

### **5. Middleware (middleware/auth_middleware.go)**
- ✅ Создан middleware с TODO комментариями
- 🔄 **Ваша задача:** Реализовать все middleware функции

## 🛠️ **Технологии для изучения**

### **JWT токены**
```bash
go get github.com/golang-jwt/jwt/v5
```
- [JWT.io](https://jwt.io/) - изучение структуры токенов
- [JWT в Go](https://pkg.go.dev/github.com/golang-jwt/jwt/v5)

### **Хеширование паролей**
```bash
go get golang.org/x/crypto/bcrypt
```
- [bcrypt в Go](https://pkg.go.dev/golang.org/x/crypto/bcrypt)
- [Почему bcrypt](https://auth0.com/blog/hashing-in-action-understanding-bcrypt/)

### **Валидация данных**
```bash
go get github.com/go-playground/validator/v10
```
- [Validator в Go](https://pkg.go.dev/github.com/go-playground/validator/v10)
- [Теги валидации](https://github.com/go-playground/validator/blob/master/README.md#baked-in-validations)

### **HTTP роутинг**
```bash
go get github.com/gorilla/mux
```
- [Gorilla Mux](https://github.com/gorilla/mux)
- [HTTP в Go](https://golang.org/pkg/net/http/)

## 📚 **Пошаговый план реализации**

### **Шаг 1: Настройка зависимостей**
1. Добавить недостающие пакеты в go.mod
2. Изучить документацию по JWT, bcrypt, validator
3. Понять, как работают HTTP handlers и middleware

### **Шаг 2: Реализация моделей**
1. Дополнить структуру User (роли, дополнительные поля)
2. Добавить валидационные теги
3. Создать функции для конвертации между моделями

### **Шаг 3: Реализация сервиса**
1. Начать с простых функций (CreateUser, GetUserByID)
2. Реализовать хеширование паролей
3. Реализовать JWT токены
4. Добавить валидацию и обработку ошибок

### **Шаг 4: Реализация handlers**
1. Начать с Register и Login
2. Добавить валидацию входных данных
3. Реализовать обработку ошибок
4. Добавить правильные HTTP статусы

### **Шаг 5: Реализация middleware**
1. Начать с AuthMiddleware
2. Реализовать извлечение токенов
3. Добавить проверку ролей
4. Реализовать CORS и логирование

### **Шаг 6: Интеграция и тестирование**
1. Связать все компоненты в main.go
2. Протестировать API endpoints
3. Проверить безопасность и валидацию

## 🔐 **Ключевые концепции безопасности**

### **JWT токены**
- **Access Token:** короткое время жизни (15 минут), содержит user_id и role
- **Refresh Token:** длительное время жизни (7 дней), только для обновления access token
- **Claims:** user_id, role, exp (expiration), iat (issued at)

### **Хеширование паролей**
- Используйте bcrypt с cost=12
- Никогда не храните пароли в открытом виде
- Сравнивайте хеши, а не пароли

### **Валидация данных**
- Валидируйте все входные данные
- Используйте теги валидации в структурах
- Проверяйте уникальность username и email

### **Middleware**
- **AuthMiddleware:** проверяет JWT токены
- **RoleMiddleware:** проверяет роли пользователей
- **CORS:** обрабатывает cross-origin запросы
- **Logging:** логирует все запросы
- **RateLimiting:** ограничивает количество запросов

## 📡 **API Endpoints для реализации**

### **Публичные endpoints**
```
POST /api/auth/register    - Регистрация пользователя
POST /api/auth/login       - Вход пользователя
POST /api/auth/refresh     - Обновление access token
```

### **Защищенные endpoints (требуют JWT)**
```
GET  /api/auth/profile     - Профиль текущего пользователя
PUT  /api/auth/profile     - Обновление профиля
POST /api/auth/logout      - Выход пользователя
```

### **Админские endpoints (требуют роль admin)**
```
GET    /api/users          - Список всех пользователей
GET    /api/users/{id}     - Информация о пользователе
PUT    /api/users/{id}     - Обновление пользователя
DELETE /api/users/{id}     - Удаление пользователя
```

## 🧪 **Тестирование API**

### **Инструменты для тестирования**
- **Postman** - для ручного тестирования
- **curl** - для командной строки
- **Thunder Client** - расширение VS Code
- **Insomnia** - альтернатива Postman

### **Примеры запросов**
```bash
# Регистрация
curl -X POST http://localhost:8080/api/auth/register \
  -H "Content-Type: application/json" \
  -d '{"username":"testuser","email":"test@example.com","password":"password123"}'

# Вход
curl -X POST http://localhost:8080/api/auth/login \
  -H "Content-Type: application/json" \
  -d '{"username":"testuser","password":"password123"}'

# Получение профиля (с токеном)
curl -X GET http://localhost:8080/api/auth/profile \
  -H "Authorization: Bearer YOUR_ACCESS_TOKEN"
```

## 📝 **Структура ответов**

### **Успешный ответ**
```json
{
  "success": true,
  "data": {
    "user": {
      "id": "uuid",
      "username": "testuser",
      "email": "test@example.com",
      "role": "user",
      "created_at": "2024-01-01T00:00:00Z"
    }
  }
}
```

### **Ответ с ошибкой**
```json
{
  "success": false,
  "error": "Invalid credentials",
  "code": "AUTH_001"
}
```

## 🚨 **Обработка ошибок**

### **HTTP статусы**
- **200 OK** - успешный запрос
- **201 Created** - ресурс создан
- **400 Bad Request** - неверные данные
- **401 Unauthorized** - неверные учетные данные
- **403 Forbidden** - недостаточно прав
- **404 Not Found** - ресурс не найден
- **429 Too Many Requests** - превышен лимит запросов
- **500 Internal Server Error** - внутренняя ошибка сервера

### **Коды ошибок**
- **AUTH_001** - неверные учетные данные
- **AUTH_002** - токен истек
- **AUTH_003** - неверный токен
- **AUTH_004** - недостаточно прав
- **VAL_001** - неверные данные
- **DB_001** - ошибка базы данных

## 🎯 **Критерии оценки**

### **Обязательные требования**
- ✅ Все TODO пункты реализованы
- ✅ JWT токены работают корректно
- ✅ Пароли хешируются с помощью bcrypt
- ✅ Валидация входных данных
- ✅ Правильные HTTP статусы
- ✅ Middleware для аутентификации
- ✅ Проверка ролей

### **Дополнительные бонусы**
- 🏆 Rate limiting
- 🏆 Логирование запросов
- 🏆 CORS настройки
- 🏆 Обработка ошибок с кодами
- 🏆 Тесты для API
- 🏆 Документация API (Swagger)

## 🚀 **Готово к началу!**

Теперь у вас есть:
1. **Полная структура проекта** с TODO комментариями
2. **Подробные задания** по каждому компоненту
3. **Ссылки на документацию** для изучения
4. **Примеры и шаблоны** для реализации

**Начинайте с изучения технологий, затем реализуйте код пошагово!**

Когда закончите, я проверю ваш код и дам обратную связь. 🎯
