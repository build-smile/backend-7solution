package services

import (
	"github.com/build-smile/backend-7solution/internal/core/domain"
	"github.com/build-smile/backend-7solution/internal/core/port"
	"github.com/gin-gonic/gin"
)

type updateUserSvc struct {
	repo port.UserRepo
}

func NewUpdateUserSvc(repo port.UserRepo) domain.UpadteUserSvc {
	return &updateUserSvc{
		repo: repo,
	}
}

func (s *updateUserSvc) Execute(c *gin.Context, req domain.UpdateUserSvcReq) error {
	err := s.repo.UpdateUser(c, req.Id, req.Name, req.Email)
	return err
}
