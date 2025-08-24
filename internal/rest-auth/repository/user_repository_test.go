package repository

import (
	"testing"
	"time"

	"golang-chat/internal/rest-auth/model"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// setupTestDB создает тестовую базу данных в памяти
func setupTestDB(t *testing.T) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		t.Fatalf("Failed to connect to test database: %v", err)
	}

	// Автоматическая миграция для тестов
	if err := db.AutoMigrate(&model.User{}); err != nil {
		t.Fatalf("Failed to migrate test database: %v", err)
	}

	return db
}

// TestGormUserRepository_CreateUser тестирует создание пользователя
func TestGormUserRepository_CreateUser(t *testing.T) {
	db := setupTestDB(t)
	repo := NewGormUserRepository(db)

	user := &model.User{
		Username:  "testuser",
		Email:     "test@example.com",
		Password:  "hashedpassword",
		Role:      "user",
		FirstName: "Test",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err := repo.CreateUser(user)
	if err != nil {
		t.Errorf("CreateUser failed: %v", err)
	}

	if user.ID == "" {
		t.Error("User ID should be generated")
	}
}

// TestGormUserRepository_GetUserByUsername тестирует поиск пользователя по username
func TestGormUserRepository_GetUserByUsername(t *testing.T) {
	db := setupTestDB(t)
	repo := NewGormUserRepository(db)

	// Создаем пользователя
	user := &model.User{
		Username:  "testuser",
		Email:     "test@example.com",
		Password:  "hashedpassword",
		Role:      "user",
		FirstName: "Test",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if err := repo.CreateUser(user); err != nil {
		t.Fatalf("Failed to create user: %v", err)
	}

	// Ищем пользователя
	foundUser, err := repo.GetUserByUsername("testuser")
	if err != nil {
		t.Errorf("GetUserByUsername failed: %v", err)
	}

	if foundUser.Username != "testuser" {
		t.Errorf("Expected username 'testuser', got '%s'", foundUser.Username)
	}
}

// TestGormUserRepository_UpdateUser тестирует обновление пользователя
func TestGormUserRepository_UpdateUser(t *testing.T) {
	db := setupTestDB(t)
	repo := NewGormUserRepository(db)

	// Создаем пользователя
	user := &model.User{
		Username:  "testuser",
		Email:     "test@example.com",
		Password:  "hashedpassword",
		Role:      "user",
		FirstName: "Test",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if err := repo.CreateUser(user); err != nil {
		t.Fatalf("Failed to create user: %v", err)
	}

	// Обновляем пользователя
	user.FirstName = "Updated"
	user.Role = "admin"

	if err := repo.UpdateUser(user); err != nil {
		t.Errorf("UpdateUser failed: %v", err)
	}

	// Проверяем обновление
	foundUser, err := repo.GetUserByID(user.ID)
	if err != nil {
		t.Fatalf("Failed to get updated user: %v", err)
	}

	if foundUser.FirstName != "Updated" {
		t.Errorf("Expected FirstName 'Updated', got '%s'", foundUser.FirstName)
	}

	if foundUser.Role != "admin" {
		t.Errorf("Expected Role 'admin', got '%s'", foundUser.Role)
	}
}

// TestGormUserRepository_GetUsersByRole тестирует поиск пользователей по роли
func TestGormUserRepository_GetUsersByRole(t *testing.T) {
	db := setupTestDB(t)
	repo := NewGormUserRepository(db)

	// Создаем несколько пользователей с разными ролями
	users := []*model.User{
		{Username: "user1", Email: "user1@example.com", Password: "hash1", Role: "user", FirstName: "User1"},
		{Username: "user2", Email: "user2@example.com", Password: "hash2", Role: "admin", FirstName: "User2"},
		{Username: "user3", Email: "user3@example.com", Password: "hash3", Role: "user", FirstName: "User3"},
	}

	for _, user := range users {
		user.CreatedAt = time.Now()
		user.UpdatedAt = time.Now()
		if err := repo.CreateUser(user); err != nil {
			t.Fatalf("Failed to create user: %v", err)
		}
	}

	// Получаем пользователей с ролью "user"
	userUsers, err := repo.GetUsersByRole("user")
	if err != nil {
		t.Errorf("GetUsersByRole failed: %v", err)
	}

	if len(userUsers) != 2 {
		t.Errorf("Expected 2 users with role 'user', got %d", len(userUsers))
	}

	// Получаем пользователей с ролью "admin"
	adminUsers, err := repo.GetUsersByRole("admin")
	if err != nil {
		t.Errorf("GetUsersByRole failed: %v", err)
	}

	if len(adminUsers) != 1 {
		t.Errorf("Expected 1 user with role 'admin', got %d", len(adminUsers))
	}
}
