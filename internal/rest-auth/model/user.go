package model

import "time"

// User представляет пользователя в системе
// TODO: Добавить поля для ролей и дополнительной информации
type User struct {
	ID        string    `json:"id" db:"id"`
	Username  string    `json:"username" db:"username"`
	Email     string    `json:"email" db:"email"`
	Password  string    `json:"-" db:"password"` // "-" означает, что поле не будет сериализоваться в JSON
	Role      string    `json:"role" db:"role"`  // TODO: Добавить роли (user, admin, moderator)
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
	FirstName string    `json:"first_name" db:"first_name"`
	// TODO: Добавить дополнительные поля:
	// - FirstName string
	// - LastName string
	// - Phone string
	// - IsActive bool
	// - LastLoginAt time.Time
}

// CreateUserRequest - запрос на создание пользователя
type CreateUserRequest struct {
	Username string `json:"username" validate:"required,min=3,max=50,alphanum"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8,containsany=ABCDEFGHIJKLMNOPQRSTUVWXYZ,containsany=abcdefghijklmnopqrstuvwxyz,containsany=0123456789,containsany=!@#$%^&*"`
}

// LoginRequest - запрос на вход пользователя
type LoginRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

// LoginResponse - ответ на успешный вход
type LoginResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	User         *User  `json:"user"`
	// TODO: Добавить дополнительные поля:
	// - ExpiresIn int (время жизни токена в секундах)
	// - TokenType string (обычно "Bearer")
}

// RefreshTokenRequest - запрос на обновление токена
type RefreshTokenRequest struct {
	RefreshToken string `json:"refresh_token" validate:"required"`
}

// UpdateProfileRequest - запрос на обновление профиля
type UpdateProfileRequest struct {
	Username  string `json:"username" validate:"omitempty,min=3,max=50,username_format"`
	Email     string `json:"email" validate:"omitempty,email,email_unique"`
	FirstName string `json:"first_name" validate:"omitempty,min=2,max=50"`
	LastName  string `json:"last_name" validate:"omitempty,min=2,max=50"`
	Phone     string `json:"phone" validate:"omitempty,len=10"`
	// TODO: Добавить поля для обновления:
	// - FirstName string
	// - LastName string
	// - Phone string
}

// UserProfile - профиль пользователя для отображения
type UserProfile struct {
	ID        string    `json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Role      string    `json:"role"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	// TODO: Добавить поля:
	// - FirstName string
	// - LastName string
	// - Phone string
	// - IsActive bool
	// - LastLoginAt time.Time
}
