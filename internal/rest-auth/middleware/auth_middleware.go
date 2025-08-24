package middleware

import (
	"net/http"

	"golang-chat/internal/rest-auth/service"
)

// AuthMiddleware проверяет JWT токен и добавляет информацию о пользователе в контекст
// TODO: Реализовать:
// 1. Извлечение токена из Authorization header
// 2. Валидацию JWT токена
// 3. Добавление user_id и role в контекст запроса
// 4. Обработку ошибок аутентификации
func AuthMiddleware(authService *service.AuthService) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// TODO: Получить Authorization header
			// TODO: Проверить формат "Bearer <token>"
			// TODO: Извлечь токен
			// TODO: Валидировать токен через authService.ValidateToken
			// TODO: Получить user_id из токена
			// TODO: Получить роль пользователя (можно из токена или базы данных)
			// TODO: Добавить user_id и role в контекст
			// TODO: Вызвать next.ServeHTTP для продолжения обработки
			// TODO: В случае ошибки вернуть 401 Unauthorized

			// Временная заглушка
			next.ServeHTTP(w, r)
		})
	}
}

// RoleMiddleware проверяет роль пользователя для доступа к защищенным маршрутам
// TODO: Реализовать:
// 1. Проверку роли пользователя из контекста
// 2. Сравнение с требуемой ролью
// 3. Возврат 403 Forbidden при недостаточных правах
func RoleMiddleware(requiredRole string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// TODO: Получить роль пользователя из контекста
			// TODO: Сравнить с requiredRole
			// TODO: Если роль не совпадает, вернуть 403 Forbidden
			// TODO: Если роль совпадает, вызвать next.ServeHTTP

			// Временная заглушка
			next.ServeHTTP(w, r)
		})
	}
}

// CORS middleware для обработки Cross-Origin запросов
// TODO: Реализовать:
// 1. Установку CORS заголовков
// 2. Обработку preflight запросов (OPTIONS)
// 3. Настройку разрешенных origins, methods, headers
func CORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// TODO: Установить Access-Control-Allow-Origin
		// TODO: Установить Access-Control-Allow-Methods
		// TODO: Установить Access-Control-Allow-Headers
		// TODO: Установить Access-Control-Allow-Credentials
		// TODO: Обработать OPTIONS запросы
		// TODO: Вызвать next.ServeHTTP для остальных запросов

		// Временная заглушка
		next.ServeHTTP(w, r)
	})
}

// Logging middleware для логирования HTTP запросов
// TODO: Реализовать:
// 1. Логирование входящих запросов
// 2. Логирование ответов
// 3. Измерение времени выполнения
// 4. Логирование ошибок
func Logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// TODO: Записать время начала запроса
		// TODO: Логировать метод, URL, IP адрес
		// TODO: Создать response writer wrapper для перехвата статуса
		// TODO: Измерить время выполнения
		// TODO: Логировать результат (статус, время выполнения)

		// Временная заглушка
		next.ServeHTTP(w, r)
	})
}

// RateLimiting middleware для ограничения количества запросов
// TODO: Реализовать:
// 1. Подсчет запросов по IP адресу
// 2. Ограничение количества запросов в единицу времени
// 3. Возврат 429 Too Many Requests при превышении лимита
func RateLimiting(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// TODO: Получить IP адрес клиента
		// TODO: Проверить лимит запросов для данного IP
		// TODO: Если лимит превышен, вернуть 429
		// TODO: Если лимит не превышен, увеличить счетчик и вызвать next.ServeHTTP

		// Временная заглушка
		next.ServeHTTP(w, r)
	})
}

// Helper функции

// extractTokenFromHeader извлекает JWT токен из Authorization header
// TODO: Реализовать:
// 1. Получение Authorization header
// 2. Проверку формата "Bearer <token>"
// 3. Извлечение токена
// 4. Валидацию формата
func extractTokenFromHeader(r *http.Request) (string, error) {
	// TODO: Получить Authorization header
	// TODO: Проверить, что header начинается с "Bearer "
	// TODO: Извлечь токен после "Bearer "
	// TODO: Проверить, что токен не пустой
	// TODO: Вернуть токен или ошибку

	return "", nil
}

// addUserToContext добавляет информацию о пользователе в контекст
// TODO: Реализовать:
// 1. Создание нового контекста с user_id
// 2. Создание нового контекста с role
// 3. Обновление контекста запроса
func addUserToContext(r *http.Request, userID, role string) *http.Request {
	// TODO: Создать контекст с user_id
	// TODO: Создать контекст с role
	// TODO: Обновить контекст запроса
	// TODO: Вернуть обновленный запрос

	return r
}

// getUserFromContext получает информацию о пользователе из контекста
// TODO: Реализовать:
// 1. Извлечение user_id из контекста
// 2. Извлечение role из контекста
// 3. Обработку случая, когда информация не найдена
func getUserFromContext(r *http.Request) (userID, role string, err error) {
	// TODO: Получить user_id из контекста
	// TODO: Получить role из контекста
	// TODO: Проверить, что оба значения найдены
	// TODO: Вернуть значения или ошибку

	return "", "", nil
}

// getClientIP получает реальный IP адрес клиента
// TODO: Реализовать:
// 1. Проверку X-Forwarded-For header (для прокси)
// 2. Проверку X-Real-IP header
// 3. Fallback на RemoteAddr
func getClientIP(r *http.Request) string {
	// TODO: Проверить X-Forwarded-For header
	// TODO: Проверить X-Real-IP header
	// TODO: Использовать RemoteAddr как fallback
	// TODO: Обработать случаи с несколькими IP в X-Forwarded-For

	return r.RemoteAddr
}
