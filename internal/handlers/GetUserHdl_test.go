package handlers

import (
	"errors"
	"github.com/build-smile/backend-7solution/internal/core/domain"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

// Mock service
type mockGetUserSvc struct {
	executeFunc func(c *gin.Context, username string) (*domain.GetUserRes, error)
}

func (m *mockGetUserSvc) Execute(c *gin.Context, username string) (*domain.GetUserRes, error) {
	return m.executeFunc(c, username)
}

func TestGetUserHdl_Handle(t *testing.T) {
	gin.SetMode(gin.TestMode)

	tests := []struct {
		name         string
		userID       string
		mockRes      *domain.GetUserRes
		mockErr      error
		wantStatus   int
		wantContains string
	}{
		{
			name:   "user found",
			userID: "testuser",
			mockRes: &domain.GetUserRes{
				Id:        "1",
				Name:      "testuser",
				Email:     "test@example.com",
				CreatedAt: time.Now(),
			},
			mockErr:      nil,
			wantStatus:   http.StatusOK,
			wantContains: `"name":"testuser"`,
		},
		{
			name:         "user not found",
			userID:       "nouser",
			mockRes:      nil,
			mockErr:      errors.New("not found"),
			wantStatus:   http.StatusOK, // Gin will still write 200 unless you handle error with JSON
			wantContains: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockSvc := &mockGetUserSvc{
				executeFunc: func(c *gin.Context, username string) (*domain.GetUserRes, error) {
					return tt.mockRes, tt.mockErr
				},
			}
			handler := NewGetUserHdl(mockSvc)

			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)
			c.Params = gin.Params{{Key: "id", Value: tt.userID}}

			handler.Handle(c)

			if w.Code != tt.wantStatus {
				t.Errorf("got status %d, want %d", w.Code, tt.wantStatus)
			}
			if tt.wantContains != "" && !contains(w.Body.String(), tt.wantContains) {
				t.Errorf("response body does not contain %q: %s", tt.wantContains, w.Body.String())
			}
		})
	}
}

func contains(s, substr string) bool {
	return len(substr) == 0 || (len(s) > 0 && (string(s) == substr || (len(s) > len(substr) && (s[0:len(substr)] == substr || contains(s[1:], substr)))))
}
