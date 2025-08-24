package repository

import (
	"database/sql"
	"errors"
	"time"

	"golang-chat/internal/rest-auth/model"
)

// UserRepository интерфейс для работы с пользователями
type UserRepository interface {
	CreateUser(user *model.User) error
	GetUserByID(id string) (*model.User, error)
	GetUserByUsername(username string) (*model.User, error)
	GetUserByEmail(email string) (*model.User, error)
	UpdateUser(user *model.User) error
	DeleteUser(id string) error
}

// PostgresUserRepository реализация репозитория для PostgreSQL
type PostgresUserRepository struct {
	db *sql.DB
}

// NewPostgresUserRepository создает новый репозиторий
func NewPostgresUserRepository(db *sql.DB) *PostgresUserRepository {
	return &PostgresUserRepository{db: db}
}

// CreateUser создает нового пользователя в базе данных
func (r *PostgresUserRepository) CreateUser(user *model.User) error {
	query := `
		INSERT INTO users (id, username, email, password_hash, role, first_name, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
	`

	_, err := r.db.Exec(query,
		user.ID,
		user.Username,
		user.Email,
		user.Password,
		user.Role,
		user.FirstName,
		user.CreatedAt,
		user.UpdatedAt,
	)

	return err
}

// GetUserByID получает пользователя по ID
func (r *PostgresUserRepository) GetUserByID(id string) (*model.User, error) {
	query := `
		SELECT id, username, email, password_hash, role, first_name, created_at, updated_at
		FROM users WHERE id = $1
	`

	user := &model.User{}
	err := r.db.QueryRow(query, id).Scan(
		&user.ID,
		&user.Username,
		&user.Email,
		&user.Password,
		&user.Role,
		&user.FirstName,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("user not found")
		}
		return nil, err
	}

	return user, nil
}

// GetUserByUsername получает пользователя по username
func (r *PostgresUserRepository) GetUserByUsername(username string) (*model.User, error) {
	query := `
		SELECT id, username, email, password_hash, role, first_name, created_at, updated_at
		FROM users WHERE username = $1
	`

	user := &model.User{}
	err := r.db.QueryRow(query, username).Scan(
		&user.ID,
		&user.Username,
		&user.Email,
		&user.Password,
		&user.Role,
		&user.FirstName,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("user not found")
		}
		return nil, err
	}

	return user, nil
}

// GetUserByEmail получает пользователя по email
func (r *PostgresUserRepository) GetUserByEmail(email string) (*model.User, error) {
	query := `
		SELECT id, username, email, password_hash, role, first_name, created_at, updated_at
		FROM users WHERE email = $1
	`

	user := &model.User{}
	err := r.db.QueryRow(query, email).Scan(
		&user.ID,
		&user.Username,
		&user.Email,
		&user.Password,
		&user.Role,
		&user.FirstName,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("user not found")
		}
		return nil, err
	}

	return user, nil
}

// UpdateUser обновляет пользователя
func (r *PostgresUserRepository) UpdateUser(user *model.User) error {
	query := `
		UPDATE users 
		SET username = $2, email = $3, password_hash = $4, role = $5, first_name = $6, updated_at = $7
		WHERE id = $1
	`

	user.UpdatedAt = time.Now()
	_, err := r.db.Exec(query,
		user.ID,
		user.Username,
		user.Email,
		user.Password,
		user.Role,
		user.FirstName,
		user.UpdatedAt,
	)

	return err
}

// DeleteUser удаляет пользователя
func (r *PostgresUserRepository) DeleteUser(id string) error {
	query := `DELETE FROM users WHERE id = $1`

	result, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errors.New("user not found")
	}

	return nil
}
