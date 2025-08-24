package service

import (
	"errors"
	"fmt"
	"time"

	"golang-chat/internal/rest-auth/model"
	"golang-chat/internal/rest-auth/repository"
	"golang-chat/pkg/config"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

// AuthService предоставляет бизнес-логику для авторизации
type AuthService struct {
	config         *config.Config
	userRepository repository.UserRepository
}

// NewAuthService создает новый экземпляр AuthService
func NewAuthService(config *config.Config, userRepository repository.UserRepository) *AuthService {
	return &AuthService{
		config:         config,
		userRepository: userRepository,
	}
}

// CreateUser создает нового пользователя
func (s *AuthService) CreateUser(req *model.CreateUserRequest) (*model.User, error) {
	// 1. Проверяем, что запрос не nil
	if req == nil {
		return nil, errors.New("request is nil")
	}

	// 2. Проверяем уникальность username
	if _, err := s.userRepository.GetUserByUsername(req.Username); err == nil {
		return nil, errors.New("username already exists")
	}

	// 3. Проверяем уникальность email
	if _, err := s.userRepository.GetUserByEmail(req.Email); err == nil {
		return nil, errors.New("email already exists")
	}

	// 4. Хешируем пароль
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, errors.New("failed to hash password")
	}

	// 5. Создаем пользователя
	user := &model.User{
		ID:        uuid.New().String(),
		Username:  req.Username,
		Email:     req.Email,
		Password:  string(hashedPassword),
		Role:      "user", // Роль по умолчанию
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	// 6. Сохраняем пользователя в базу данных
	if err := s.userRepository.CreateUser(user); err != nil {
		return nil, fmt.Errorf("failed to save user to database: %w", err)
	}

	return user, nil
}

// GetUserByID получает пользователя по ID
func (s *AuthService) GetUserByID(id string) (*model.User, error) {
	return s.userRepository.GetUserByID(id)
}

// GetUserByUsername получает пользователя по username
func (s *AuthService) GetUserByUsername(username string) (*model.User, error) {
	return s.userRepository.GetUserByUsername(username)
}

// GetUserByEmail получает пользователя по email
func (s *AuthService) GetUserByEmail(email string) (*model.User, error) {
	return s.userRepository.GetUserByEmail(email)
}

// DeleteUserByID удаляет пользователя по ID
func (s *AuthService) DeleteUserByID(userID string) error {
	return s.userRepository.DeleteUser(userID)
}

// Login выполняет аутентификацию пользователя
// TODO: Реализовать:
// 1. Поиск пользователя по username
// 2. Проверку пароля с помощью bcrypt.CompareHashAndPassword
// 3. Генерацию JWT access и refresh токенов
// 4. Обновление времени последнего входа
// 5. Возврат токенов и информации о пользователе
func (s *AuthService) Login(req *model.LoginRequest) (*model.LoginResponse, error) {
	// TODO: Найти пользователя по username
	// TODO: Проверить пароль
	// TODO: Сгенерировать access token (время жизни: 15 минут)
	// TODO: Сгенерировать refresh token (время жизни: 7 дней)
	// TODO: Обновить LastLoginAt
	// TODO: Вернуть токены и информацию о пользователе

	return nil, errors.New("not implemented")
}

// RefreshToken обновляет access token с помощью refresh token
// TODO: Реализовать:
// 1. Валидацию refresh token
// 2. Извлечение user_id из токена
// 3. Проверку, что пользователь существует
// 4. Генерацию нового access token
// 5. Возврат нового access token
func (s *AuthService) RefreshToken(req *model.RefreshTokenRequest) (string, error) {
	// TODO: Валидировать refresh token
	// TODO: Извлечь user_id из токена
	// TODO: Проверить существование пользователя
	// TODO: Сгенерировать новый access token
	// TODO: Вернуть новый токен

	return "", errors.New("not implemented")
}

// ValidateToken проверяет валидность access token
// TODO: Реализовать:
// 1. Парсинг JWT токена
// 2. Проверку подписи
// 3. Проверку времени жизни
// 4. Извлечение user_id
func (s *AuthService) ValidateToken(tokenString string) (string, error) {
	// TODO: Парсить JWT токен
	// TODO: Проверить подпись
	// TODO: Проверить время жизни
	// TODO: Извлечь user_id из claims

	return "", errors.New("not implemented")
}

// UpdateProfile обновляет профиль пользователя
// TODO: Реализовать:
// 1. Проверку существования пользователя
// 2. Валидацию новых данных
// 3. Проверку уникальности username и email
// 4. Обновление полей
// 5. Обновление времени изменения
func (s *AuthService) UpdateProfile(userID string, req *model.UpdateProfileRequest) (*model.User, error) {
	// TODO: Найти пользователя по ID
	// TODO: Проверить уникальность новых username/email
	// TODO: Обновить поля
	// TODO: Обновить UpdatedAt
	// TODO: Сохранить изменения

	return nil, errors.New("not implemented")
}

// DeleteUser удаляет пользователя
// TODO: Реализовать:
// 1. Проверку существования пользователя
// 2. Проверку прав доступа (только админ или сам пользователь)
// 3. Удаление пользователя из базы данных
func (s *AuthService) DeleteUser(userID string) error {
	// TODO: Найти пользователя по ID
	// TODO: Проверить права доступа
	// TODO: Удалить пользователя

	return errors.New("not implemented")
}

// GetAllUsers возвращает список всех пользователей
// TODO: Реализовать:
// 1. Проверку прав доступа (только для админов)
// 2. Пагинацию результатов
// 3. Фильтрацию по ролям
// 4. Возврат только безопасной информации (без паролей)
func (s *AuthService) GetAllUsers(limit, offset int) ([]*model.UserProfile, error) {
	// TODO: Проверить права доступа
	// TODO: Реализовать пагинацию
	// TODO: Вернуть профили пользователей (без паролей)

	return nil, errors.New("not implemented")
}

// GenerateAccessToken генерирует JWT access token
func (s *AuthService) GenerateAccessToken(userID, role string) (string, error) {
	// Создаем claims для JWT токена
	claims := jwt.MapClaims{
		"user_id": userID,
		"role":    role,
		"type":    "access",
		"exp":     time.Now().Add(15 * time.Minute).Unix(), // Время жизни: 15 минут
		"iat":     time.Now().Unix(),                       // Время создания
	}

	// Создаем JWT токен
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Подписываем токен секретным ключом
	tokenString, err := token.SignedString([]byte(s.config.JWTSecret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// GenerateRefreshToken генерирует JWT refresh token
func (s *AuthService) GenerateRefreshToken(userID string) (string, error) {
	// Создаем claims для refresh токена
	claims := jwt.MapClaims{
		"user_id": userID,
		"type":    "refresh",
		"exp":     time.Now().Add(7 * 24 * time.Hour).Unix(), // Время жизни: 7 дней
		"iat":     time.Now().Unix(),                         // Время создания
	}

	// Создаем JWT токен
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Подписываем токен секретным ключом
	tokenString, err := token.SignedString([]byte(s.config.JWTSecret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// hashPassword хеширует пароль с помощью bcrypt
// TODO: Реализовать:
// 1. Импорт "golang.org/x/crypto/bcrypt"
// 2. Хеширование пароля с cost=12
func (s *AuthService) hashPassword(password string) (string, error) {
	// TODO: Хешировать пароль с помощью bcrypt.GenerateFromPassword
	// TODO: Вернуть хешированный пароль

	return "", errors.New("not implemented")
}

// checkPassword проверяет пароль с помощью bcrypt
// TODO: Реализовать:
// 1. Сравнение хешированного пароля с введенным
// 2. Использование bcrypt.CompareHashAndPassword
func (s *AuthService) checkPassword(hashedPassword, password string) error {
	// TODO: Сравнить пароли с помощью bcrypt.CompareHashAndPassword
	// TODO: Вернуть ошибку, если пароли не совпадают

	return errors.New("not implemented")
}
