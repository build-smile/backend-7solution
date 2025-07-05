package domain

import (
	"github.com/gin-gonic/gin"
	"time"
)

type CreateUserSvc interface {
	Execute(c *gin.Context, req CreateUserSvcReq) error
}
type CreateUserSvcReq struct {
	Id        string    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"createdAt"`
}

type CreateUserSvcRes struct {
}
