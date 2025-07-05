package services

import (
	"github.com/build-smile/backend-7solution/internal/core/port"
	"github.com/gin-gonic/gin"
)

type DeleteUserSvc struct {
	repo port.UserRepo
}

func NewDeleteUserSvc(repo port.UserRepo) *DeleteUserSvc {
	return &DeleteUserSvc{
		repo: repo,
	}
}

func (s *DeleteUserSvc) Execute(c *gin.Context, id string) error {
	err := s.repo.DeleteUser(c, id)
	if err != nil {
		return err
	}
	return nil
}
