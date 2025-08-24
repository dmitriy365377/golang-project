package validation

import (
	"errors"
	"regexp"
	"strings"

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

// Метод для валидации LoginRequest
func (v *Validation) ValidateLoginRequest(req *model.LoginRequest) error {
	if err := v.validator.Struct(req); err != nil {
		// Преобразуем ошибки валидации в более понятные сообщения
		return v.translateValidationErrors(err)
	}
	return nil
}

// translateValidationErrors преобразует ошибки валидации в понятные сообщения
func (v *Validation) translateValidationErrors(err error) error {
	if validationErrors, ok := err.(validator.ValidationErrors); ok {
		var messages []string

		for _, e := range validationErrors {
			switch e.Tag() {
			case "required":
				messages = append(messages, e.Field()+" is required")
			case "min":
				messages = append(messages, e.Field()+" must be at least "+e.Param()+" characters")
			case "max":
				messages = append(messages, e.Field()+" must be no more than "+e.Param()+" characters")
			case "email":
				messages = append(messages, e.Field()+" must be a valid email address")
			case "alphanum":
				messages = append(messages, e.Field()+" must contain only letters and numbers")
			default:
				messages = append(messages, e.Field()+" validation failed: "+e.Tag())
			}
		}

		return errors.New(strings.Join(messages, "; "))
	}

	return err
}

// Метод для валидации RefreshTokenRequest
func (v *Validation) ValidateRefreshTokenRequest(req *model.RefreshTokenRequest) error {
	return v.validator.Struct(req)
}

// Метод для валидации UpdateProfileRequest
func (v *Validation) ValidateUpdateProfileRequest(req *model.UpdateProfileRequest) error {
	return v.validator.Struct(req)
}

// ValidateStruct - универсальный метод для валидации любой структуры
func (v *Validation) ValidateStruct(s interface{}) error {
	return v.validator.Struct(s)
}
