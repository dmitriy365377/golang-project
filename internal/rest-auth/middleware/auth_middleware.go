package middleware

import (
	"context"
	"net/http"

	"golang-chat/internal/rest-auth/service"
)

// Константы для контекста
const (
	UserIDKey = "user_id"
	RoleKey   = "role"
)

// AuthMiddleware проверяет JWT токен из cookies и добавляет информацию о пользователе в контекст
func AuthMiddleware(authService *service.AuthService) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Получаем access token из cookie
			accessTokenCookie, err := r.Cookie("access_token")
			if err != nil {
				// Если cookie не найден, возвращаем 401
				http.Error(w, "Unauthorized: No access token", http.StatusUnauthorized)
				return
			}

			// Валидируем токен
			userID, err := authService.ValidateToken(accessTokenCookie.Value)
			if err != nil {
				// Если токен невалиден, возвращаем 401
				http.Error(w, "Unauthorized: Invalid token", http.StatusUnauthorized)
				return
			}

			// Получаем пользователя для получения роли
			user, err := authService.GetUserByID(userID)
			if err != nil {
				http.Error(w, "Unauthorized: User not found", http.StatusUnauthorized)
				return
			}

			// Добавляем информацию о пользователе в контекст
			ctx := context.WithValue(r.Context(), UserIDKey, userID)
			ctx = context.WithValue(ctx, RoleKey, user.Role)

			// Создаем новый запрос с обновленным контекстом
			r = r.WithContext(ctx)

			// Продолжаем обработку запроса
			next.ServeHTTP(w, r)
		})
	}
}

// RoleMiddleware проверяет роль пользователя для доступа к защищенным маршрутам
func RoleMiddleware(requiredRole string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Получаем роль пользователя из контекста
			role, ok := r.Context().Value(RoleKey).(string)
			if !ok {
				http.Error(w, "Forbidden: Role not found", http.StatusForbidden)
				return
			}

			// Проверяем, есть ли у пользователя требуемая роль
			if role != requiredRole && role != "admin" { // admin имеет доступ ко всему
				http.Error(w, "Forbidden: Insufficient permissions", http.StatusForbidden)
				return
			}

			// Продолжаем обработку запроса
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

// GetUserID получает user_id из контекста
func GetUserID(r *http.Request) (string, bool) {
	userID, ok := r.Context().Value(UserIDKey).(string)
	return userID, ok
}

// GetUserRole получает роль пользователя из контекста
func GetUserRole(r *http.Request) (string, bool) {
	role, ok := r.Context().Value(RoleKey).(string)
	return role, ok
}

// RequireAuth проверяет, что пользователь аутентифицирован
func RequireAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, ok := GetUserID(r)
		if !ok {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	})
}
