package main

import (
	"log"
	"net/http"

	"golang-chat/internal/rest-auth/database"
	"golang-chat/internal/rest-auth/handler"
	"golang-chat/internal/rest-auth/repository"
	"golang-chat/internal/rest-auth/service"
	"golang-chat/internal/rest-auth/validation"
	"golang-chat/pkg/config"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	// Загружаем конфигурацию
	cfg := config.Load()

	// Подключаемся к базе данных
	db, err := database.ConnectToPostgres(cfg.DatabaseURL)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	defer database.CloseDatabase(db)

	// Создаем валидатор
	validator := validation.NewCustomValidator()

	// Создаем репозиторий пользователей
	userRepository := repository.NewPostgresUserRepository(db)

	// Создаем Auth Service
	authService := service.NewAuthService(cfg, userRepository)

	// Создаем Auth Handler
	authHandler := handler.NewAuthHandler(authService, validator)

	// Создаем роутер с Chi
	router := chi.NewRouter()

	// Middleware
	router.Use(middleware.Logger)    // Логирование запросов
	router.Use(middleware.Recoverer) // Восстановление после panic
	router.Use(middleware.RequestID) // Уникальные ID для запросов
	router.Use(middleware.RealIP)    // Реальные IP адреса
	router.Use(middleware.CleanPath) // Очистка путей
	router.Use(middleware.GetHead)   // Поддержка HEAD запросов

	// Маршрут для проверки здоровья сервиса
	router.Get("/health", healthCheckHandler)

	// Маршрут для корня
	router.Get("/", rootHandler)

	// Группируем маршруты для авторизации
	router.Route("/api/auth", func(r chi.Router) {
		r.Post("/register", authHandler.Register)
		r.Post("/login", authHandler.Login)
		r.Post("/refresh", authHandler.RefreshToken)
		r.Get("/profile", authHandler.GetProfile)
	})

	// Запускаем HTTP сервер
	log.Println("REST Auth Service starting on :8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}

// Обработчик для проверки здоровья сервиса
func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"status": "ok", "service": "auth-service", "router": "chi"}`))
}

// Обработчик для корневого маршрута
func rootHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "Golang Chat Auth Service", "version": "1.0.0", "router": "chi"}`))
}
