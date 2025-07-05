package handlers

import (
	"github.com/build-smile/backend-7solution/internal/core/domain"
	"github.com/build-smile/backend-7solution/utils"
	"github.com/gin-gonic/gin"

	"net/http"
)

type UpdateUserHdl struct {
	service domain.UpadteUserSvc
}

func NewUpdateUserHdl(service domain.UpadteUserSvc) *UpdateUserHdl {
	return &UpdateUserHdl{
		service: service,
	}
}

func (h *UpdateUserHdl) Handle(c *gin.Context) {
	req := domain.UpdateUserSvcReq{}
	if err := c.Bind(&req); err != nil {
		c.Error(err)
		return
	}
	if err := utils.ValidateRequest(req); err != nil {
		c.Error(err)
		return
	}
	id := c.Param("id")
	req.Id = id
	err := h.service.Execute(c, req)
	if err != nil {
		c.Error(err)
		return
	}
	c.Status(http.StatusOK) // Updated
}
