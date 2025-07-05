package domain

import "github.com/gin-gonic/gin"

type LoginUserSvc interface {
	Execute(c *gin.Context, req LoginUserReq) (*LoginUserRes, error)
}

type LoginUserReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
type LoginUserRes struct {
	AccessToken string `json:"accessToken"`
}
