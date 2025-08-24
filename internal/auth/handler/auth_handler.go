package handler

import (
	"context"

	"golang-chat/internal/auth/service"
	"golang-chat/proto/auth"
)

type AuthHandler struct {
	auth.UnimplementedAuthServiceServer
	auth.UnimplementedUserServiceServer
	auth.UnimplementedAccessServiceServer

	authService *service.AuthService
}

func NewAuthHandler(authService *service.AuthService) *AuthHandler {
	return &AuthHandler{
		authService: authService,
	}
}

// User Service Methods
func (h *AuthHandler) Create(ctx context.Context, req *auth.CreateUserRequest) (*auth.CreateUserResponse, error) {
	user, err := h.authService.CreateUser(req.Username, req.Email, req.Password)
	if err != nil {
		return &auth.CreateUserResponse{Error: err.Error()}, nil
	}

	return &auth.CreateUserResponse{
		User: &auth.User{
			Id:        user.ID,
			Username:  user.Username,
			Email:     user.Email,
			CreatedAt: user.CreatedAt.Format("2006-01-02T15:04:05Z"),
			UpdatedAt: user.UpdatedAt.Format("2006-01-02T15:04:05Z"),
		},
	}, nil
}

func (h *AuthHandler) Get(ctx context.Context, req *auth.GetUserRequest) (*auth.GetUserResponse, error) {
	user, err := h.authService.GetUser(req.Id)
	if err != nil {
		return &auth.GetUserResponse{Error: err.Error()}, nil
	}

	return &auth.GetUserResponse{
		User: &auth.User{
			Id:        user.ID,
			Username:  user.Username,
			Email:     user.Email,
			CreatedAt: user.CreatedAt.Format("2006-01-02T15:04:05Z"),
			UpdatedAt: user.UpdatedAt.Format("2006-01-02T15:04:05Z"),
		},
	}, nil
}

func (h *AuthHandler) GetList(ctx context.Context, req *auth.GetUserListRequest) (*auth.GetUserListResponse, error) {
	// TODO: Implement user list retrieval
	return &auth.GetUserListResponse{Error: "Not implemented yet"}, nil
}

func (h *AuthHandler) Update(ctx context.Context, req *auth.UpdateUserRequest) (*auth.UpdateUserResponse, error) {
	// TODO: Implement user update
	return &auth.UpdateUserResponse{Error: "Not implemented yet"}, nil
}

func (h *AuthHandler) Delete(ctx context.Context, req *auth.DeleteUserRequest) (*auth.DeleteUserResponse, error) {
	// TODO: Implement user deletion
	return &auth.DeleteUserResponse{Error: "Not implemented yet"}, nil
}

// Auth Service Methods
func (h *AuthHandler) Login(ctx context.Context, req *auth.LoginRequest) (*auth.LoginResponse, error) {
	accessToken, refreshToken, err := h.authService.Login(req.Username, req.Password)
	if err != nil {
		return &auth.LoginResponse{Error: err.Error()}, nil
	}

	return &auth.LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func (h *AuthHandler) GetAccessToken(ctx context.Context, req *auth.GetAccessTokenRequest) (*auth.GetAccessTokenResponse, error) {
	// TODO: Implement access token refresh
	return &auth.GetAccessTokenResponse{Error: "Not implemented yet"}, nil
}

func (h *AuthHandler) GetRefreshToken(ctx context.Context, req *auth.GetRefreshTokenRequest) (*auth.GetRefreshTokenResponse, error) {
	// TODO: Implement refresh token refresh
	return &auth.GetRefreshTokenResponse{Error: "Not implemented yet"}, nil
}

// Access Service Methods
func (h *AuthHandler) Check(ctx context.Context, req *auth.CheckAccessRequest) (*auth.CheckAccessResponse, error) {
	userID, err := h.authService.ValidateToken(req.AccessToken)
	if err != nil {
		return &auth.CheckAccessResponse{
			HasAccess: false,
			Error:     err.Error(),
		}, nil
	}

	return &auth.CheckAccessResponse{
		HasAccess: true,
		UserId:    userID,
	}, nil
}
