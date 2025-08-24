package validation

import (
	"strings"
	"testing"

	"golang-chat/internal/rest-auth/model"
)

func TestValidateLoginRequest(t *testing.T) {
	validator := NewCustomValidator()

	tests := []struct {
		name    string
		request model.LoginRequest
		wantErr bool
	}{
		{
			name: "Valid login request",
			request: model.LoginRequest{
				Username: "testuser",
				Password: "password123",
			},
			wantErr: false,
		},
		{
			name: "Empty username",
			request: model.LoginRequest{
				Username: "",
				Password: "password123",
			},
			wantErr: true,
		},
		{
			name: "Empty password",
			request: model.LoginRequest{
				Username: "testuser",
				Password: "",
			},
			wantErr: true,
		},
		{
			name: "Username too short",
			request: model.LoginRequest{
				Username: "ab",
				Password: "password123",
			},
			wantErr: true,
		},
		{
			name: "Both fields empty",
			request: model.LoginRequest{
				Username: "",
				Password: "",
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validator.ValidateLoginRequest(&tt.request)

			if (err != nil) != tt.wantErr {
				t.Errorf("ValidateLoginRequest() error = %v, wantErr %v", err, tt.wantErr)
			}

			if err != nil {
				t.Logf("Validation error: %v", err)
			}
		})
	}
}

func TestTranslateValidationErrors(t *testing.T) {
	validator := NewCustomValidator()

	// Создаем невалидный запрос
	req := model.LoginRequest{
		Username: "",
		Password: "",
	}

	err := validator.ValidateLoginRequest(&req)
	if err == nil {
		t.Fatal("Expected validation error, got nil")
	}

	// Проверяем, что ошибка содержит понятные сообщения
	errorMsg := err.Error()
	t.Logf("Translated error: %s", errorMsg)

	// Проверяем, что сообщения содержат нужные слова
	if !strings.Contains(errorMsg, "Username is required") {
		t.Errorf("Error message should contain 'Username is required', got: %s", errorMsg)
	}

	if !strings.Contains(errorMsg, "Password is required") {
		t.Errorf("Error message should contain 'Password is required', got: %s", errorMsg)
	}
}
