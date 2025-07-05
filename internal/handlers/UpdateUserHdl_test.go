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
type mockUpdateUserSvc struct {
	executeFunc func(c *gin.Context, req domain.UpdateUserSvcReq) error
}

func (m *mockUpdateUserSvc) Execute(c *gin.Context, req domain.UpdateUserSvcReq) error {
	return m.executeFunc(c, req)
}

func TestUpdateUserHdl_Handle(t *testing.T) {
	gin.SetMode(gin.TestMode)

	tests := []struct {
		name       string
		id         string
		body       string
		mockErr    error
		wantStatus int
	}{
		{
			name:       "valid update",
			id:         "507f1f77bcf86cd799439011",
			body:       `{"name":"newname","email":"new@example.com"}`,
			mockErr:    nil,
			wantStatus: http.StatusOK,
		},
		{
			name:       "invalid json",
			id:         "507f1f77bcf86cd799439011",
			body:       `{"name":"newname","email":"new@example.com"`, // malformed
			mockErr:    nil,
			wantStatus: http.StatusBadRequest,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockSvc := &mockUpdateUserSvc{
				executeFunc: func(c *gin.Context, req domain.UpdateUserSvcReq) error {
					return tt.mockErr
				},
			}
			handler := NewUpdateUserHdl(mockSvc)

			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)
			c.Params = gin.Params{{Key: "id", Value: tt.id}}
			c.Request, _ = http.NewRequest("PATCH", "/user/"+tt.id, bytes.NewBufferString(tt.body))
			c.Request.Header.Set("Content-Type", "application/json")

			handler.Handle(c)

			if w.Code != tt.wantStatus {
				t.Errorf("got status %d, want %d", w.Code, tt.wantStatus)
			}
		})
	}
}
