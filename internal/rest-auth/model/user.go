package model

import (
	"net/http"
	"time"

	"gorm.io/gorm"
)

// User представляет пользователя в системе
// TODO: Добавить поля для ролей и дополнительной информации
type User struct {
	ID        string         `json:"id" gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	Username  string         `json:"username" gorm:"uniqueIndex;not null;size:50"`
	Email     string         `json:"email" gorm:"uniqueIndex;not null;size:100"`
	Password  string         `json:"-" gorm:"column:password_hash;not null;size:255"` // "-" означает, что поле не будет сериализоваться в JSON
	Role      string         `json:"role" gorm:"default:'user';size:20"`              // TODO: Добавить роли (user, admin, moderator)
	FirstName string         `json:"first_name" gorm:"size:50"`
	CreatedAt time.Time      `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time      `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"` // Soft delete
	// TODO: Добавить дополнительные поля:
	// - LastName string
	// - Phone string
	// - IsActive bool
	// - LastLoginAt time.Time
}

// TableName указывает имя таблицы для GORM
func (User) TableName() string {
	return "users"
}

// CreateUserRequest - запрос на создание пользователя
type CreateUserRequest struct {
	Username string `json:"username" validate:"required,min=3,max=50,alphanum"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8,containsany=ABCDEFGHIJKLMNOPQRSTUVWXYZ,containsany=abcdefghijklmnopqrstuvwxyz,containsany=0123456789,containsany=!@#$%^&*"`
}

// LoginRequest - запрос на вход пользователя
type LoginRequest struct {
	Username string `json:"username" validate:"required,min=3"`
	Password string `json:"password" validate:"required,min=1"`
}

// LoginResponse - ответ на успешный вход
type LoginResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	User         *User  `json:"user"`
	// Cookie настройки
	AccessTokenCookie  *http.Cookie `json:"-"`
	RefreshTokenCookie *http.Cookie `json:"-"`
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
