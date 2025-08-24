package repository

import (
	"errors"
	"time"

	"golang-chat/internal/rest-auth/model"

	"gorm.io/gorm"
)

// UserRepository интерфейс для работы с пользователями
type UserRepository interface {
	CreateUser(user *model.User) error
	GetUserByID(id string) (*model.User, error)
	GetUserByUsername(username string) (*model.User, error)
	GetUserByEmail(email string) (*model.User, error)
	UpdateUser(user *model.User) error
	DeleteUser(id string) error
	GetUsersWithPagination(page, pageSize int) ([]*model.User, int64, error)
	GetUsersByRole(role string) ([]*model.User, error)
	SearchUsers(query string) ([]*model.User, error)
	UpdateUserRole(id, role string) error
}

// GormUserRepository реализация репозитория с использованием GORM
type GormUserRepository struct {
	db *gorm.DB
}

// NewGormUserRepository создает новый репозиторий
func NewGormUserRepository(db *gorm.DB) *GormUserRepository {
	return &GormUserRepository{db: db}
}

// CreateUser создает нового пользователя в базе данных
func (r *GormUserRepository) CreateUser(user *model.User) error {
	// GORM автоматически установит CreatedAt и UpdatedAt
	return r.db.Create(user).Error
}

// GetUserByID получает пользователя по ID
func (r *GormUserRepository) GetUserByID(id string) (*model.User, error) {
	var user model.User
	err := r.db.Where("id = ?", id).First(&user).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}

	return &user, nil
}

// GetUserByUsername получает пользователя по username
func (r *GormUserRepository) GetUserByUsername(username string) (*model.User, error) {
	var user model.User
	err := r.db.Where("username = ?", username).First(&user).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}

	return &user, nil
}

// GetUserByEmail получает пользователя по email
func (r *GormUserRepository) GetUserByEmail(email string) (*model.User, error) {
	var user model.User
	err := r.db.Where("email = ?", email).First(&user).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}

	return &user, nil
}

// UpdateUser обновляет пользователя
func (r *GormUserRepository) UpdateUser(user *model.User) error {
	// GORM автоматически обновит UpdatedAt
	user.UpdatedAt = time.Now()
	return r.db.Save(user).Error
}

// DeleteUser удаляет пользователя (soft delete)
func (r *GormUserRepository) DeleteUser(id string) error {
	result := r.db.Delete(&model.User{}, "id = ?", id)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("user not found")
	}

	return nil
}

// GetUsersWithPagination получает пользователей с пагинацией
func (r *GormUserRepository) GetUsersWithPagination(page, pageSize int) ([]*model.User, int64, error) {
	var users []*model.User
	var total int64

	// Подсчитываем общее количество пользователей
	if err := r.db.Model(&model.User{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Получаем пользователей с пагинацией
	offset := (page - 1) * pageSize
	err := r.db.Offset(offset).Limit(pageSize).Find(&users).Error

	return users, total, err
}

// GetUsersByRole получает пользователей по роли
func (r *GormUserRepository) GetUsersByRole(role string) ([]*model.User, error) {
	var users []*model.User
	err := r.db.Where("role = ?", role).Find(&users).Error
	return users, err
}

// SearchUsers выполняет поиск пользователей по имени или email
func (r *GormUserRepository) SearchUsers(query string) ([]*model.User, error) {
	var users []*model.User
	searchQuery := "%" + query + "%"

	err := r.db.Where("username ILIKE ? OR email ILIKE ? OR first_name ILIKE ?",
		searchQuery, searchQuery, searchQuery).Find(&users).Error

	return users, err
}

// UpdateUserRole обновляет роль пользователя
func (r *GormUserRepository) UpdateUserRole(id, role string) error {
	result := r.db.Model(&model.User{}).Where("id = ?", id).Update("role", role)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("user not found")
	}

	return nil
}
