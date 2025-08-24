package validation

import (
	"regexp"

	"golang-chat/internal/rest-auth/model"

	"github.com/go-playground/validator/v10"
)

type Validation struct {
	validator *validator.Validate
}

func NewCustomValidator() *Validation {
	v := validator.New()

	// Регистрируем кастомные валидаторы
	v.RegisterValidation("username_format", validateUsernameFormat)
	v.RegisterValidation("password_strength", validatePasswordStrength)
	v.RegisterValidation("email_unique", validateEmailUnique)

	return &Validation{
		validator: v,
	}
}

// Валидация формата username
func validateUsernameFormat(fl validator.FieldLevel) bool {
	username := fl.Field().String()

	// Проверяем, что username содержит только буквы, цифры и подчеркивания
	matched, _ := regexp.MatchString(`^[a-zA-Z0-9_]+$`, username)
	return matched
}

// Валидация силы пароля
func validatePasswordStrength(fl validator.FieldLevel) bool {
	password := fl.Field().String()

	// Проверяем наличие заглавных букв
	hasUpper := regexp.MustCompile(`[A-Z]`).MatchString(password)
	// Проверяем наличие строчных букв
	hasLower := regexp.MustCompile(`[a-z]`).MatchString(password)
	// Проверяем наличие цифр
	hasNumber := regexp.MustCompile(`[0-9]`).MatchString(password)
	// Проверяем наличие спецсимволов
	hasSpecial := regexp.MustCompile(`[!@#$%^&*(),.?":{}|<>]`).MatchString(password)

	return hasUpper && hasLower && hasNumber && hasSpecial
}

// Валидация уникальности email (заглушка - нужно будет реализовать проверку в БД)
func validateEmailUnique(fl validator.FieldLevel) bool {
	// TODO: Здесь должна быть проверка в базе данных
	// Пока возвращаем true (email считается уникальным)
	return true
}

// Метод для валидации CreateUserRequest
func (v *Validation) ValidateCreateUserRequest(req *model.CreateUserRequest) error {
	return v.validator.Struct(req)
}
