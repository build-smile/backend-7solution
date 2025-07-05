package domain

import (
	"github.com/gin-gonic/gin"
	"time"
)

type GetUserSvc interface {
	Execute(c *gin.Context, username string) (*GetUserRes, error)
}

type GetUserRes struct {
	Id        string    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"createdAt"`
}

type GetUserReq struct {
	Username string `json:"username"` // Assuming you want to get user by username
}
