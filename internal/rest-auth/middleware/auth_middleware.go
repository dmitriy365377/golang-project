package middleware

import (
	"golang-chat/internal/rest-auth/service"

	"github.com/gofiber/fiber/v2"
)

// Константы для контекста
const (
	UserIDKey = "user_id"
	RoleKey   = "role"
)

// AuthMiddleware проверяет JWT токен из cookies и добавляет информацию о пользователе в контекст
func AuthMiddleware(authService *service.AuthService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Получаем access token из cookie
		accessToken := c.Cookies("access_token")
		if accessToken == "" {
			// Если cookie не найден, возвращаем 401
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Unauthorized: No access token",
			})
		}

		// Валидируем токен
		userID, err := authService.ValidateToken(accessToken)
		if err != nil {
			// Если токен невалиден, возвращаем 401
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Unauthorized: Invalid token",
			})
		}

		// Получаем пользователя для получения роли
		user, err := authService.GetUserByID(userID)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Unauthorized: User not found",
			})
		}

		// Добавляем информацию о пользователе в локальное хранилище Fiber
		c.Locals(UserIDKey, userID)
		c.Locals(RoleKey, user.Role)

		// Продолжаем обработку запроса
		return c.Next()
	}
}

// RoleMiddleware проверяет роль пользователя для доступа к защищенным маршрутам
func RoleMiddleware(requiredRole string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Получаем роль пользователя из локального хранилища
		role, ok := c.Locals(RoleKey).(string)
		if !ok {
			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
				"error": "Forbidden: Role not found",
			})
		}

		// Проверяем, есть ли у пользователя требуемая роль
		if role != requiredRole && role != "admin" { // admin имеет доступ ко всему
			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
				"error": "Forbidden: Insufficient permissions",
			})
		}

		// Продолжаем обработку запроса
		return c.Next()
	}
}

// CORS middleware для обработки Cross-Origin запросов
// TODO: Реализовать:
// 1. Установку CORS заголовков
// 2. Обработку preflight запросов (OPTIONS)
// 3. Настройку разрешенных origins, methods, headers
func CORS() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// TODO: Установить Access-Control-Allow-Origin
		// TODO: Установить Access-Control-Allow-Methods
		// TODO: Установить Access-Control-Allow-Headers
		// TODO: Установить Access-Control-Allow-Credentials
		// TODO: Обработать OPTIONS запросы
		// TODO: Вызвать c.Next() для остальных запросов

		// Временная заглушка
		return c.Next()
	}
}

// Logging middleware для логирования HTTP запросов
// TODO: Реализовать:
// 1. Логирование входящих запросов
// 2. Логирование ответов
// 3. Измерение времени выполнения
// 4. Логирование ошибок
func Logging() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// TODO: Записать время начала запроса
		// TODO: Логировать метод, URL, IP адрес
		// TODO: Создать response writer wrapper для перехвата статуса
		// TODO: Измерить время выполнения
		// TODO: Логировать результат (статус, время выполнения)

		// Временная заглушка
		return c.Next()
	}
}

// RateLimiting middleware для ограничения количества запросов
// TODO: Реализовать:
// 1. Подсчет запросов по IP адресу
// 2. Ограничение количества запросов в единицу времени
// 3. Возврат 429 Too Many Requests при превышении лимита
func RateLimiting() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// TODO: Получить IP адрес клиента
		// TODO: Проверить лимит запросов для данного IP
		// TODO: Если лимит превышен, вернуть 429
		// TODO: Если лимит не превышен, увеличить счетчик и вызвать c.Next()

		// Временная заглушка
		return c.Next()
	}
}

// Helper функции

// GetUserID получает user_id из локального хранилища
func GetUserID(c *fiber.Ctx) (string, bool) {
	userID, ok := c.Locals(UserIDKey).(string)
	return userID, ok
}

// GetUserRole получает роль пользователя из локального хранилища
func GetUserRole(c *fiber.Ctx) (string, bool) {
	role, ok := c.Locals(RoleKey).(string)
	return role, ok
}

// RequireAuth проверяет, что пользователь аутентифицирован
func RequireAuth() fiber.Handler {
	return func(c *fiber.Ctx) error {
		_, ok := GetUserID(c)
		if !ok {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Unauthorized",
			})
		}
		return c.Next()
	}
}
