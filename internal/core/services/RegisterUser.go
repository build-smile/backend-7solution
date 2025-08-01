package services

import (
	"fmt"
	"github.com/build-smile/backend-7solution/internal/core/domain"
	"github.com/build-smile/backend-7solution/internal/core/port"
	"github.com/build-smile/backend-7solution/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type RegisterUserSvc struct {
	repo port.UserRepo
}

func NewRegisterUserSvc(repo port.UserRepo) *RegisterUserSvc {
	return &RegisterUserSvc{
		repo: repo,
	}
}
func (s *RegisterUserSvc) Execute(c *gin.Context, req domain.RegisterUserSvcReq) error {
	t := time.Now()
	hashPassword, err := s.repo.HashPassword(req.Password)
	if err != nil {
		return utils.NewCustomError(http.StatusInternalServerError, fmt.Sprintf("failed to hash password: %v", err))
	}
	err = s.repo.CreateUser(
		port.User{
			Name:      req.Name,
			Email:     req.Email,
			CreatedAt: t,
			Password:  hashPassword,
		}, c)
	return err
}
