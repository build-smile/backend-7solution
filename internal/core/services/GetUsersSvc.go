package services

import (
	"github.com/build-smile/backend-7solution/internal/core/domain"
	"github.com/build-smile/backend-7solution/internal/core/port"
	"github.com/gin-gonic/gin"
)

type getUsersSvc struct {
	repo port.UserRepo
}

func NewGetUsersSvc(repo port.UserRepo) domain.GetUsersSvc {

	return &getUsersSvc{
		repo: repo,
	}
}

func (s getUsersSvc) Execute(c *gin.Context, req domain.GetUsersReq) ([]domain.GetUsersRes, error) {
	orderRes, err := s.repo.GetUsers(c)
	res := s.buildRes(orderRes)

	return res, err
}

func (s getUsersSvc) buildRes(userRes []port.User) []domain.GetUsersRes {
	res := make([]domain.GetUsersRes, 0)
	for _, it := range userRes {
		var user domain.GetUsersRes
		user.Id = it.ID
		user.Name = it.Name
		user.Email = it.Email
		user.CreatedAt = it.CreatedAt
		res = append(res, user)
	}
	return res
}
