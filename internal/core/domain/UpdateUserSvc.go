package domain

import (
	"github.com/gin-gonic/gin"
)

type UpadteUserSvc interface {
	Execute(c *gin.Context, req UpdateUserSvcReq) error
}
type UpdateUserSvcReq struct {
	Id    string `json:"id"`
	Name  string `json:"name" validate:"required"`
	Email string `json:"email" validate:"required,email"`
}

type UpadteUserSvcRes struct {
}
