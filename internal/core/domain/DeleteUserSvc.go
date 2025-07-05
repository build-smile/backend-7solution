package domain

import "github.com/gin-gonic/gin"

type DeleteUserSvc interface {
	Execute(c *gin.Context, id string) error
}
