package handlers

import (
	"github.com/build-smile/backend-7solution/internal/core/domain"
	"github.com/build-smile/backend-7solution/utils"
	"github.com/gin-gonic/gin"

	"net/http"
)

type CreateUserHdl struct {
	service domain.CreateUserSvc
}

func NewCreateUserHdl(service domain.CreateUserSvc) *CreateUserHdl {
	return &CreateUserHdl{
		service: service,
	}
}

func (h *CreateUserHdl) Handle(c *gin.Context) {
	req := domain.CreateUserSvcReq{}
	if err := c.Bind(&req); err != nil {
		c.Error(err)
		return
	}
	if err := utils.ValidateRequest(req); err != nil {
		c.Error(err)
		return
	}
	err := h.service.Execute(c, req)
	if err != nil {
		c.Error(err)
		return
	}
	c.Status(http.StatusCreated) // Created
}
