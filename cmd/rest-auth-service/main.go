package main

import (
	"log"
	"strings"

	"golang-chat/internal/rest-auth/database"
	"golang-chat/internal/rest-auth/handler"
	authmiddleware "golang-chat/internal/rest-auth/middleware"
	"golang-chat/internal/rest-auth/model"
	"golang-chat/internal/rest-auth/repository"
	"golang-chat/internal/rest-auth/service"
	"golang-chat/internal/rest-auth/validation"
	"golang-chat/pkg/config"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
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
		log.Printf("⚠️ Warning: Database migration failed: %v", err)
		log.Println("🔄 Continuing without migration...")
	} else {
		log.Println("✅ Database migration completed successfully")
	}

	// Создаем валидатор
	validator := validation.NewCustomValidator()

	// Создаем репозиторий пользователей с GORM
	userRepository := repository.NewGormUserRepository(db)

	// Создаем Auth Service
	authService := service.NewAuthService(cfg, userRepository)

	// Создаем Auth Handler
	authHandler := handler.NewAuthHandler(authService, validator)

	// Создаем Fiber приложение
	app := fiber.New(fiber.Config{
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			code := fiber.StatusInternalServerError
			if e, ok := err.(*fiber.Error); ok {
				code = e.Code
			}
			return c.Status(code).JSON(fiber.Map{
				"error": err.Error(),
			})
		},
	})

	// CORS middleware - должен быть первым!
	app.Use(cors.New(cors.Config{
		AllowOrigins:     strings.Join(cfg.CORSOrigins, ","), // Объединяем массив в строку
		AllowMethods:     strings.Join(cfg.CORSMethods, ","),
		AllowHeaders:     strings.Join(cfg.CORSHeaders, ","),
		ExposeHeaders:    "Link",
		AllowCredentials: true,
		MaxAge:           300, // 5 минут
	}))

	// Middleware
	app.Use(logger.New())    // Логирование запросов
	app.Use(recover.New())   // Восстановление после panic
	app.Use(requestid.New()) // Уникальные ID для запросов

	// Маршрут для проверки здоровья сервиса
	app.Get("/health", healthCheckHandler)

	// Маршрут для корневого маршрута
	app.Get("/", rootHandler)

	// Группируем маршруты для авторизации
	auth := app.Group("/api/auth")
	auth.Post("/register", authHandler.Register)
	auth.Post("/login", authHandler.Login)
	auth.Post("/refresh", authHandler.RefreshToken)

	// Защищенные маршруты - требуют аутентификации
	protected := auth.Group("/", authmiddleware.AuthMiddleware(authService))
	protected.Get("/profile", authHandler.GetProfile)
	protected.Post("/logout", authHandler.Logout)

	// Запускаем HTTP сервер
	log.Println("REST Auth Service starting on :8080")
	log.Fatal(app.Listen(":8080"))
}

// Обработчик для проверки здоровья сервиса
func healthCheckHandler(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"status":  "ok",
		"service": "auth-service",
		"router":  "fiber",
	})
}

// Обработчик для корневого маршрута
func rootHandler(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "Golang Chat Auth Service",
		"version": "1.0.0",
		"router":  "fiber",
	})
}
