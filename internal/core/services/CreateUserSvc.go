package services

import (
	"github.com/build-smile/backend-7solution/internal/core/domain"
	"github.com/build-smile/backend-7solution/internal/core/port"
	"github.com/gin-gonic/gin"
	"time"
)

type createUserSvc struct {
	repo port.UserRepo
}

func NewCreateUserSvc(repo port.UserRepo) domain.CreateUserSvc {
	return &createUserSvc{
		repo: repo,
	}
}

func (s *createUserSvc) Execute(c *gin.Context, req domain.CreateUserSvcReq) error {
	t := time.Now()
	hashPassword, err := s.repo.HashPassword(req.Password)
	if err != nil {
		return err
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
