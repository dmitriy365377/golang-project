package main

import (
	"log"
	"net/http"

	"golang-chat/internal/rest-auth/database"
	"golang-chat/internal/rest-auth/handler"
	authmiddleware "golang-chat/internal/rest-auth/middleware"
	"golang-chat/internal/rest-auth/model"
	"golang-chat/internal/rest-auth/repository"
	"golang-chat/internal/rest-auth/service"
	"golang-chat/internal/rest-auth/validation"
	"golang-chat/pkg/config"

	"github.com/go-chi/chi/v5"
	chimiddleware "github.com/go-chi/chi/v5/middleware"
	"github.com/rs/cors"
)

func main() {
	// Загружаем конфигурацию
	cfg := config.Load()

	// Подключаемся к базе данных через GORM
	db, err := database.ConnectToPostgres(cfg.DatabaseURL)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	defer database.CloseDatabase(db)

	// Выполняем автоматическую миграцию таблиц
	if err := db.AutoMigrate(&model.User{}); err != nil {
		log.Fatal("Failed to migrate database:", err)
	}
	log.Println("✅ Database migration completed successfully")

	// Создаем валидатор
	validator := validation.NewCustomValidator()

	// Создаем репозиторий пользователей с GORM
	userRepository := repository.NewGormUserRepository(db)

	// Создаем Auth Service
	authService := service.NewAuthService(cfg, userRepository)

	// Создаем Auth Handler
	authHandler := handler.NewAuthHandler(authService, validator)

	// Создаем роутер с Chi
	router := chi.NewRouter()

	// CORS middleware - должен быть первым!
	corsMiddleware := cors.New(cors.Options{
		AllowedOrigins:   cfg.CORSOrigins,
		AllowedMethods:   cfg.CORSMethods,
		AllowedHeaders:   cfg.CORSHeaders,
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300, // 5 минут
	})
	router.Use(corsMiddleware.Handler)

	// Middleware
	router.Use(chimiddleware.Logger)    // Логирование запросов
	router.Use(chimiddleware.Recoverer) // Восстановление после panic
	router.Use(chimiddleware.RequestID) // Уникальные ID для запросов
	router.Use(chimiddleware.RealIP)    // Реальные IP адреса
	router.Use(chimiddleware.CleanPath) // Очистка путей
	router.Use(chimiddleware.GetHead)   // Поддержка HEAD запросов

	// Маршрут для проверки здоровья сервиса
	router.Get("/health", healthCheckHandler)

	// Маршрут для корневого маршрута
	router.Get("/", rootHandler)

	// Группируем маршруты для авторизации
	router.Route("/api/auth", func(r chi.Router) {
		r.Post("/register", authHandler.Register)
		r.Post("/login", authHandler.Login)
		r.Post("/refresh", authHandler.RefreshToken)

		// Защищенные маршруты - требуют аутентификации
		r.Group(func(r chi.Router) {
			r.Use(authmiddleware.AuthMiddleware(authService))
			r.Get("/profile", authHandler.GetProfile)
			r.Post("/logout", authHandler.Logout)
		})
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
