package handlers

import (
	"bytes"
	"github.com/build-smile/backend-7solution/internal/core/domain"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httptest"
	"testing"
)

// Mock service
type mockLoginUserSvc struct {
	executeFunc func(c *gin.Context, req domain.LoginUserReq) (*domain.LoginUserRes, error)
}

func (m *mockLoginUserSvc) Execute(c *gin.Context, req domain.LoginUserReq) (*domain.LoginUserRes, error) {
	return m.executeFunc(c, req)
}

func TestLoginUserHdl_Handle(t *testing.T) {
	gin.SetMode(gin.TestMode)

	tests := []struct {
		name         string
		body         string
		mockRes      *domain.LoginUserRes
		mockErr      error
		wantStatus   int
		wantContains string
	}{
		{
			name:         "invalid json",
			body:         `{"username":"testuser","password":"pass1234"`, // malformed
			mockRes:      nil,
			mockErr:      nil,
			wantStatus:   http.StatusBadRequest,
			wantContains: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockSvc := &mockLoginUserSvc{
				executeFunc: func(c *gin.Context, req domain.LoginUserReq) (*domain.LoginUserRes, error) {
					return tt.mockRes, tt.mockErr
				},
			}
			handler := NewLoginUserHdl(mockSvc)

			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)
			c.Request, _ = http.NewRequest("POST", "/login", bytes.NewBufferString(tt.body))
			c.Request.Header.Set("Content-Type", "application/json")

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
