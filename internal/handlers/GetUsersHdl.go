package handlers

import (
	"github.com/build-smile/backend-7solution/internal/core/domain"
	"github.com/gin-gonic/gin"
	"net/http"
)

type GetUsersHdl struct {
	service domain.GetUsersSvc
}

func NewGetUsersHdl(service domain.GetUsersSvc) *GetUsersHdl {
	return &GetUsersHdl{
		service: service,
	}
}

func (h *GetUsersHdl) Handle(c *gin.Context) {
	req := domain.GetUsersReq{}
	res, err := h.service.Execute(c, req)
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, res)
}
