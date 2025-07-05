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
type mockRegisterUserSvc struct {
	executeFunc func(c *gin.Context, req domain.RegisterUserSvcReq) error
}

func (m *mockRegisterUserSvc) Execute(c *gin.Context, req domain.RegisterUserSvcReq) error {
	return m.executeFunc(c, req)
}

func TestRegisterUserHdl_Handle(t *testing.T) {
	gin.SetMode(gin.TestMode)

	tests := []struct {
		name       string
		body       string
		mockErr    error
		wantStatus int
	}{
		{
			name:       "Success",
			body:       `{"name":"testuser","email":"test@example.com","password":"pass1234"}`,
			mockErr:    nil,
			wantStatus: http.StatusOK,
		},
		{
			name:       "invalid json",
			body:       `{"name":"testuser","email":"test@example.com"`, // malformed
			mockErr:    nil,
			wantStatus: http.StatusBadRequest,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockSvc := &mockRegisterUserSvc{
				executeFunc: func(c *gin.Context, req domain.RegisterUserSvcReq) error {
					return tt.mockErr
				},
			}
			handler := NewRegisterUserHdl(mockSvc)

			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)
			c.Request, _ = http.NewRequest("POST", "/register", bytes.NewBufferString(tt.body))
			c.Request.Header.Set("Content-Type", "application/json")

			handler.Handle(c)

			if w.Code != tt.wantStatus {
				t.Errorf("got status %d, want %d", w.Code, tt.wantStatus)
			}
		})
	}
}
