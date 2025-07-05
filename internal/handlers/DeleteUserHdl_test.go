package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httptest"
	"testing"
)

// Mock service
type mockDeleteUserSvc struct {
	executeFunc func(c *gin.Context, id string) error
}

func (m *mockDeleteUserSvc) Execute(c *gin.Context, id string) error {
	return m.executeFunc(c, id)
}

func TestDeleteUserHdl_Handle(t *testing.T) {
	gin.SetMode(gin.TestMode)

	tests := []struct {
		name       string
		userID     string
		mockErr    error
		wantStatus int
	}{
		{"valid id", "507f1f77bcf86cd799439011", nil, http.StatusOK},
		{"missing id", "", nil, http.StatusBadRequest},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockSvc := &mockDeleteUserSvc{
				executeFunc: func(c *gin.Context, id string) error {
					return tt.mockErr
				},
			}
			handler := NewDeleteUserHdl(mockSvc)

			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)
			c.Params = gin.Params{{Key: "id", Value: tt.userID}}

			handler.Handle(c)

			if w.Code != tt.wantStatus {
				t.Errorf("got status %d, want %d", w.Code, tt.wantStatus)
			}
		})
	}
}
