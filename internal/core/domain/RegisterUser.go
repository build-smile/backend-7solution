package domain

import "github.com/gin-gonic/gin"

type RegisterUserSvc interface {
	Execute(c *gin.Context, req RegisterUserSvcReq) error
}
type RegisterUserSvcReq struct {
	Id        string `json:"id" `
	Name      string `json:"name" validate:"required"`
	Email     string `json:"email" validate:"required"`
	Password  string `json:"password" validate:"required"`
	CreatedAt string `json:"createdAt"`
}
