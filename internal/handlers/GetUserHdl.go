package handlers

import (
	"github.com/build-smile/backend-7solution/internal/core/domain"
	"github.com/gin-gonic/gin"
	"net/http"
)

type GetUserHdl struct {
	service domain.GetUserSvc
}

func NewGetUserHdl(service domain.GetUserSvc) *GetUserHdl {
	return &GetUserHdl{
		service: service,
	}
}

func (h *GetUserHdl) Handle(c *gin.Context) {
	username := c.Param("id")
	res, err := h.service.Execute(c, username)
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, res)
}
