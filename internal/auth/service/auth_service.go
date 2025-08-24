package service

import (
	"errors"
	"time"

	"golang-chat/internal/auth/model"
	"golang-chat/pkg/config"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type AuthService struct {
	config *config.Config
	users  map[string]*model.User
}

func NewAuthService() *AuthService {
	return &AuthService{
		users: make(map[string]*model.User),
	}
}

func (s *AuthService) CreateUser(username, email, password string) (*model.User, error) {
	// Проверка существования пользователя
	for _, user := range s.users {
		if user.Username == username || user.Email == email {
			return nil, errors.New("user already exists")
		}
	}

	user := &model.User{
		ID:        uuid.New().String(),
		Username:  username,
		Email:     email,
		Password:  password, // В реальном приложении здесь должно быть хеширование
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	s.users[user.ID] = user
	return user, nil
}

func (s *AuthService) GetUser(id string) (*model.User, error) {
	user, exists := s.users[id]
	if !exists {
		return nil, errors.New("user not found")
	}
	return user, nil
}

func (s *AuthService) Login(username, password string) (string, string, error) {
	var user *model.User
	for _, u := range s.users {
		if u.Username == username && u.Password == password {
			user = u
			break
		}
	}

	if user == nil {
		return "", "", errors.New("invalid credentials")
	}

	accessToken, err := s.generateAccessToken(user.ID)
	if err != nil {
		return "", "", err
	}

	refreshToken, err := s.generateRefreshToken(user.ID)
	if err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil
}

func (s *AuthService) ValidateToken(tokenString string) (string, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(s.config.JWTSecret), nil
	})

	if err != nil {
		return "", err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userID, ok := claims["user_id"].(string)
		if !ok {
			return "", errors.New("invalid token claims")
		}
		return userID, nil
	}

	return "", errors.New("invalid token")
}

func (s *AuthService) generateAccessToken(userID string) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(time.Hour * 1).Unix(), // 1 час
		"type":    "access",
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(s.config.JWTSecret))
}

func (s *AuthService) generateRefreshToken(userID string) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(time.Hour * 24 * 7).Unix(), // 7 дней
		"type":    "refresh",
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(s.config.JWTSecret))
}
