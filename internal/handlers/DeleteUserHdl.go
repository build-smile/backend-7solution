package handlers

import (
	"github.com/build-smile/backend-7solution/internal/core/domain"
	"github.com/gin-gonic/gin"
	"net/http"
)

type DeleteUserHdl struct {
	service domain.DeleteUserSvc
}

func NewDeleteUserHdl(service domain.DeleteUserSvc) *DeleteUserHdl {
	return &DeleteUserHdl{
		service: service,
	}
}
func (h *DeleteUserHdl) Handle(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(400, gin.H{"error": "User ID is required"})
		return
	}

	err := h.service.Execute(c, id)
	if err != nil {
		c.Error(err)
		return
	}

	c.Status(http.StatusOK) // No Content
}
