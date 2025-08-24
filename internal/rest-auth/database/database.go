package database

import (
	"fmt"
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// ConnectToPostgres подключается к PostgreSQL базе данных через GORM
func ConnectToPostgres(databaseURL string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(databaseURL), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info), // Логирование SQL запросов
	})
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}

	// Получаем sql.DB для настройки пула соединений
	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("failed to get sql.DB: %w", err)
	}

	// Настройка пула соединений
	sqlDB.SetMaxOpenConns(25)
	sqlDB.SetMaxIdleConns(25)
	sqlDB.SetConnMaxLifetime(5 * time.Minute)

	// Проверка подключения
	if err := sqlDB.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	log.Println("✅ Successfully connected to PostgreSQL database via GORM")
	return db, nil
}

// CloseDatabase закрывает соединение с базой данных
func CloseDatabase(db *gorm.DB) {
	if db != nil {
		sqlDB, err := db.DB()
		if err != nil {
			log.Printf("Error getting sql.DB: %v", err)
			return
		}
		
		if err := sqlDB.Close(); err != nil {
			log.Printf("Error closing database: %v", err)
		} else {
			log.Println("✅ Database connection closed")
		}
	}
}

// AutoMigrate выполняет автоматическую миграцию таблиц
func AutoMigrate(db *gorm.DB) error {
	// Импортируем модель здесь, чтобы избежать циклических зависимостей
	// Это будет вызвано из основного приложения
	log.Println("🔄 Auto-migration will be performed from main application")
	return nil
}
