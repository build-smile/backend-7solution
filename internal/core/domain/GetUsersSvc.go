package domain

import (
	"github.com/gin-gonic/gin"
	"time"
)

type GetUsersSvc interface {
	Execute(c *gin.Context, req GetUsersReq) ([]GetUsersRes, error)
}

type GetUsersRes struct {
	Id        string    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"createdAt"`
}

type GetUsersReq struct {
}
