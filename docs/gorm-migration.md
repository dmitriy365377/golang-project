# Миграция на GORM

## Обзор изменений

Проект был успешно мигрирован с ручных SQL запросов на GORM (Go Object Relational Mapper).

## Что изменилось

### 1. Модель User (`internal/rest-auth/model/user.go`)

- Добавлены GORM теги для автоматической работы с базой данных
- Автоматическое управление временными метками (`CreatedAt`, `UpdatedAt`)
- Поддержка soft delete через `DeletedAt`
- Автоматическая генерация UUID для ID

```go
type User struct {
    ID        string         `json:"id" gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
    Username  string         `json:"username" gorm:"uniqueIndex;not null;size:50"`
    Email     string         `json:"email" gorm:"uniqueIndex;not null;size:100"`
    Password  string         `json:"-" gorm:"column:password_hash;not null;size:255"`
    Role      string         `json:"role" gorm:"default:'user';size:20"`
    FirstName string         `json:"first_name" gorm:"size:50"`
    CreatedAt time.Time      `json:"created_at" gorm:"autoCreateTime"`
    UpdatedAt time.Time      `json:"updated_at" gorm:"autoUpdateTime"`
    DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}
```

### 2. Репозиторий (`internal/rest-auth/repository/user_repository.go`)

- Заменен `PostgresUserRepository` на `GormUserRepository`
- Убраны все ручные SQL запросы
- Добавлены новые методы для демонстрации возможностей GORM

#### Основные методы:
```go
// Создание пользователя
func (r *GormUserRepository) CreateUser(user *model.User) error {
    return r.db.Create(user).Error
}

// Поиск по ID
func (r *GormUserRepository) GetUserByID(id string) (*model.User, error) {
    var user model.User
    err := r.db.Where("id = ?", id).First(&user).Error
    // ... обработка ошибок
    return &user, nil
}

// Обновление
func (r *GormUserRepository) UpdateUser(user *model.User) error {
    user.UpdatedAt = time.Now()
    return r.db.Save(user).Error
}
```

#### Новые методы:
- `GetUsersWithPagination` - пагинация
- `GetUsersByRole` - фильтрация по роли
- `SearchUsers` - поиск по нескольким полям
- `UpdateUserRole` - обновление роли

### 3. База данных (`internal/rest-auth/database/database.go`)

- Изменен тип возвращаемого значения с `*sql.DB` на `*gorm.DB`
- Добавлено логирование SQL запросов
- Сохранена настройка пула соединений

### 4. Главный файл (`cmd/rest-auth-service/main.go`)

- Добавлена автоматическая миграция таблиц
- Обновлен создатель репозитория

## Преимущества GORM

### 1. **Автоматическая миграция**
```go
// GORM автоматически создаст таблицы на основе структуры
db.AutoMigrate(&model.User{})
```

### 2. **Безопасность**
- Защита от SQL инъекций
- Автоматическое экранирование параметров

### 3. **Удобство**
- Нет необходимости писать SQL вручную
- Автоматическое управление временными метками
- Встроенная поддержка soft delete

### 4. **Производительность**
- Поддержка prepared statements
- Кэширование запросов
- Оптимизация N+1 проблем

### 5. **Расширяемость**
- Легко добавлять новые поля и методы
- Поддержка различных типов баз данных
- Встроенные хуки и callbacks

## Примеры использования

### Создание пользователя
```go
user := &model.User{
    Username:  "john_doe",
    Email:     "john@example.com",
    Password:  hashedPassword,
    FirstName: "John",
}
err := userRepository.CreateUser(user)
```

### Поиск с фильтрацией
```go
// Поиск по роли
admins, err := userRepository.GetUsersByRole("admin")

// Поиск по тексту
users, err := userRepository.SearchUsers("john")

// Пагинация
users, total, err := userRepository.GetUsersWithPagination(1, 10)
```

### Обновление
```go
user.Role = "admin"
err := userRepository.UpdateUser(user)

// Или обновление конкретного поля
err := userRepository.UpdateUserRole(userID, "moderator")
```

## Зависимости

Добавлены новые зависимости в `go.mod`:
```
gorm.io/gorm v1.30.1
gorm.io/driver/postgres v1.6.0
```

## Запуск

После миграции сервис запускается как обычно:

```bash
go run ./cmd/rest-auth-service
```

GORM автоматически:
1. Подключится к базе данных
2. Создаст таблицы (если их нет)
3. Выполнит миграции
4. Запустит HTTP сервер

## Миграции

При изменении структуры модели просто перезапустите сервис - GORM автоматически обновит схему базы данных.

## Откат изменений

Если нужно вернуться к ручным SQL запросам:
1. Откатите изменения в файлах
2. Удалите GORM зависимости: `go mod tidy`
3. Восстановите старую реализацию репозитория
