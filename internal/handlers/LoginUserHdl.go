package handlers

import (
	"github.com/build-smile/backend-7solution/internal/core/domain"
	"github.com/build-smile/backend-7solution/utils"
	"github.com/gin-gonic/gin"
)

type LoginUserHdl struct {
	loginUserSvc domain.LoginUserSvc
}

func NewLoginUserHdl(loginUserSvc domain.LoginUserSvc) *LoginUserHdl {
	return &LoginUserHdl{
		loginUserSvc: loginUserSvc,
	}
}
func (h *LoginUserHdl) Handle(c *gin.Context) {
	req := domain.LoginUserReq{}
	if err := c.Bind(&req); err != nil {
		c.Error(err)
		return
	}
	if err := utils.ValidateRequest(req); err != nil {
		c.Error(err)
		return
	}
	res, err := h.loginUserSvc.Execute(c, req)
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(200, res) // OK
}
