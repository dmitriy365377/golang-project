package handler

import (
	"errors"
	"strings"

	"golang-chat/internal/rest-auth/model"
	"golang-chat/internal/rest-auth/service"
	"golang-chat/internal/rest-auth/validation"

	"github.com/gofiber/fiber/v2"
)

// Кастомные ошибки для лучшей обработки
var (
	ErrInvalidCredentials = errors.New("invalid credentials")
	ErrUserNotFound       = errors.New("user not found")
)

// AuthHandler обрабатывает HTTP запросы для авторизации
type AuthHandler struct {
	authService *service.AuthService
	validator   *validation.Validation
}

// NewAuthHandler создает новый экземпляр AuthHandler
func NewAuthHandler(authService *service.AuthService, validator *validation.Validation) *AuthHandler {
	return &AuthHandler{
		authService: authService,
		validator:   validator,
	}
}

func (h *AuthHandler) Register(c *fiber.Ctx) error {
	// 1. Парсинг JSON
	var req model.CreateUserRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	// 2. Валидация данных
	if err := h.validator.ValidateCreateUserRequest(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Validation error: " + err.Error(),
		})
	}

	// 3. Проверка уникальности username
	existingUser, err := h.authService.GetUserByUsername(req.Username)
	if err == nil && existingUser != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Username already exists",
		})
	}

	// 4. Проверка уникальности email
	existingUser, err = h.authService.GetUserByEmail(req.Email)
	if err == nil && existingUser != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Email already exists",
		})
	}

	// 5. Создание пользователя
	user, err := h.authService.CreateUser(&req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create user: " + err.Error(),
		})
	}

	// 6. Генерация токенов
	accessToken, err := h.authService.GenerateAccessToken(user.ID, user.Role)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to generate access token: " + err.Error(),
		})
	}

	refreshToken, err := h.authService.GenerateRefreshToken(user.ID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to generate refresh token: " + err.Error(),
		})
	}

	// 7. Формирование и отправка ответа
	response := model.LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		User:         user,
	}

	return c.Status(fiber.StatusCreated).JSON(response)
}

// Login выполняет вход пользователя
func (h *AuthHandler) Login(c *fiber.Ctx) error {
	// Парсим JSON из request body
	var req model.LoginRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	// Валидируем данные
	if err := h.validator.ValidateLoginRequest(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Validation error: " + err.Error(),
		})
	}

	// Вызываем authService.Login
	response, err := h.authService.Login(&req)
	if err != nil {
		// Обрабатываем различные типы ошибок
		switch {
		case err.Error() == "invalid credentials":
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Invalid username or password",
			})
		case err.Error() == "user not found":
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Invalid username or password",
			})
		default:
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Login failed: " + err.Error(),
			})
		}
	}

	// В случае успеха возвращаем 200 OK с токенами и информацией о пользователе
	// Устанавливаем cookies
	c.Cookie(response.AccessTokenCookie)
	c.Cookie(response.RefreshTokenCookie)

	return c.Status(fiber.StatusOK).JSON(response)
}

// RefreshToken обновляет access token
func (h *AuthHandler) RefreshToken(c *fiber.Ctx) error {
	// 1. Парсинг JSON из request body
	var req model.RefreshTokenRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	// 2. Валидация данных
	if err := h.validator.ValidateRefreshTokenRequest(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Validation error: " + err.Error(),
		})
	}

	// 3. Вызов authService.RefreshToken
	response, err := h.authService.RefreshToken(&req)
	if err != nil {
		// Обрабатываем различные типы ошибок
		switch {
		case strings.Contains(err.Error(), "invalid refresh token"):
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Invalid refresh token",
			})
		case strings.Contains(err.Error(), "user not found"):
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "User not found",
			})
		default:
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to refresh token: " + err.Error(),
			})
		}
	}

	// 4. В случае успеха возвращаем 200 OK с новым access token
	return c.Status(fiber.StatusOK).JSON(response)
}

// GetProfile возвращает профиль текущего пользователя
// TODO: Реализовать:
// 1. Извлечение user_id из JWT токена (из middleware)
// 2. Вызов authService.GetUserByID
// 3. Обработку ошибок
// 4. Возврат профиля пользователя
func (h *AuthHandler) GetProfile(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(string)
	user, err := h.authService.GetUserByID(userID)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "User not found",
		})
	}
	return c.Status(fiber.StatusOK).JSON(user)
}

// UpdateProfile обновляет профиль пользователя
// TODO: Реализовать:
// 1. Извлечение user_id из JWT токена
// 2. Парсинг JSON из request body
// 3. Валидацию данных
// 4. Вызов authService.UpdateProfile
// 5. Обработку ошибок
// 6. Возврат обновленного профиля
func (h *AuthHandler) UpdateProfile(c *fiber.Ctx) error {
	// TODO: Получить user_id из контекста
	// TODO: Парсить JSON из request body в UpdateProfileRequest
	// TODO: Валидировать данные
	// TODO: Вызвать authService.UpdateProfile
	// TODO: Обработать ошибки
	// TODO: В случае успеха вернуть 200 OK с обновленным профилем

	return c.Status(fiber.StatusNotImplemented).JSON(fiber.Map{
		"error": "not implemented",
	})
}

// Logout выполняет выход пользователя
func (h *AuthHandler) Logout(c *fiber.Ctx) error {
	// Очищаем cookies, устанавливая их в прошлое
	accessCookie := &fiber.Cookie{
		Name:     "access_token",
		Value:    "",
		Path:     "/",
		HTTPOnly: true,
		Secure:   false,
		SameSite: "Strict",
		MaxAge:   -1, // Удаляем cookie
	}

	refreshCookie := &fiber.Cookie{
		Name:     "refresh_token",
		Value:    "",
		Path:     "/",
		HTTPOnly: true,
		Secure:   false,
		SameSite: "Strict",
		MaxAge:   -1, // Удаляем cookie
	}

	// Устанавливаем cookies для удаления
	c.Cookie(accessCookie)
	c.Cookie(refreshCookie)

	// Возвращаем успешный ответ
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Successfully logged out",
	})
}

// GetAllUsers возвращает список всех пользователей (только для админов)
// TODO: Реализовать:
// 1. Проверку роли пользователя (должен быть admin)
// 2. Извлечение параметров пагинации из query string
// 3. Вызов authService.GetAllUsers
// 4. Обработку ошибок
// 5. Возврат списка пользователей
func (h *AuthHandler) GetAllUsers(c *fiber.Ctx) error {
	// TODO: Проверить роль пользователя (должен быть admin)
	// TODO: Получить limit и offset из query string (?limit=10&offset=0)
	// TODO: Установить значения по умолчанию (limit=10, offset=0)
	// TODO: Вызвать authService.GetAllUsers
	// TODO: Обработать ошибки
	// TODO: В случае успеха вернуть 200 OK со списком пользователей

	return c.Status(fiber.StatusNotImplemented).JSON(fiber.Map{
		"error": "not implemented",
	})
}

func (h *AuthHandler) GetUserByID(c *fiber.Ctx) error {
	userID := c.Params("id")
	user, err := h.authService.GetUserByID(userID)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "User not found",
		})
	}
	return c.Status(fiber.StatusOK).JSON(user)
}

// UpdateUser обновляет информацию о пользователе (только для админов)
// TODO: Реализовать:
// 1. Извлечение user_id из URL параметров
// 2. Проверку роли (должен быть admin)
// 3. Парсинг JSON из request body
// 4. Валидацию данных
// 5. Вызов authService.UpdateProfile
// 6. Обработку ошибок
func (h *AuthHandler) UpdateUser(c *fiber.Ctx) error {
	// TODO: Получить user_id из URL параметров
	// TODO: Проверить роль пользователя (должен быть admin)
	// TODO: Парсить JSON из request body
	// TODO: Валидировать данные
	// TODO: Вызвать authService.UpdateProfile
	// TODO: Обработать ошибки
	// TODO: В случае успеха вернуть 200 OK с обновленной информацией

	return c.Status(fiber.StatusNotImplemented).JSON(fiber.Map{
		"error": "not implemented",
	})
}

// DeleteUser удаляет пользователя (только для админов)
// TODO: Реализовать:
// 1. Извлечение user_id из URL параметров
// 2. Проверку роли (должен быть admin)
// 3. Вызов authService.DeleteUser
// 4. Обработку ошибок
// 5. Возврат успешного ответа
func (h *AuthHandler) DeleteUser(c *fiber.Ctx) error {
	// TODO: Получить user_id из URL параметров
	// TODO: Проверить роль пользователя (должен быть admin)
	// TODO: Вызвать authService.DeleteUser
	// TODO: Обработать ошибки
	// TODO: В случае успеха вернуть 200 OK с сообщением об удалении

	return c.Status(fiber.StatusNotImplemented).JSON(fiber.Map{
		"error": "not implemented",
	})
}

// Helper функции

// extractUserIDFromContext извлекает user_id из контекста запроса
// TODO: Реализовать:
// 1. Получение user_id из контекста (установленного middleware)
// 2. Обработку случая, когда user_id не найден
func extractUserIDFromContext(c *fiber.Ctx) (string, error) {
	// TODO: Получить user_id из контекста
	// TODO: Вернуть ошибку, если user_id не найден

	return "", nil
}

// extractUserRoleFromContext извлекает роль пользователя из контекста
// TODO: Реализовать:
// 1. Получение роли из контекста
// 2. Обработку случая, когда роль не найдена
func extractUserRoleFromContext(c *fiber.Ctx) (string, error) {
	// TODO: Получить роль из контекста
	// TODO: Вернуть ошибку, если роль не найдена

	return "", nil
}

// parsePaginationParams парсит параметры пагинации из query string
// TODO: Реализовать:
// 1. Извлечение limit и offset из query string
// 2. Установку значений по умолчанию
// 3. Валидацию параметров
func parsePaginationParams(c *fiber.Ctx) (limit, offset int) {
	// TODO: Получить limit из query string (?limit=10)
	// TODO: Получить offset из query string (?offset=0)
	// TODO: Установить значения по умолчанию (limit=10, offset=0)
	// TODO: Валидировать параметры (limit > 0, offset >= 0)

	return 10, 0
}
