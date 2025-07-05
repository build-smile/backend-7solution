package services

import (
	"github.com/build-smile/backend-7solution/internal/core/domain"
	"github.com/build-smile/backend-7solution/internal/core/port"
	"github.com/build-smile/backend-7solution/utils"
	"github.com/gin-gonic/gin"
)

type LoginUserSvc struct {
	repo port.UserRepo
}

func NewLoginUserSvc(repo port.UserRepo) *LoginUserSvc {
	return &LoginUserSvc{
		repo: repo,
	}
}
func (s *LoginUserSvc) Execute(c *gin.Context, req domain.LoginUserReq) (*domain.LoginUserRes, error) {
	user, err := s.repo.GetUser(port.User{Name: req.Username}, c)
	if err != nil {
		return nil, err
	}

	err = s.repo.CheckPasswordHash(req.Password, user.Password)
	if err != nil {
		return nil, err
	}

	jwt, _, err := utils.GenerateJWT(user.Name)
	if err != nil {
		return nil, err
	}

	return &domain.LoginUserRes{AccessToken: jwt}, nil
}
