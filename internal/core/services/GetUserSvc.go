package services

import (
	"errors"
	"fmt"
	"github.com/build-smile/backend-7solution/internal/core/domain"
	"github.com/build-smile/backend-7solution/internal/core/port"
	"github.com/build-smile/backend-7solution/utils"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
)

type getUserSvc struct {
	repo port.UserRepo
}

func (g getUserSvc) Execute(c *gin.Context, username string) (*domain.GetUserRes, error) {
	user, err := g.repo.GetUser(port.User{Name: username}, c)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, utils.NewCustomError(http.StatusNotFound, "User not found The user with the specified username does not exist")
		}
		return nil, utils.NewCustomError(http.StatusInternalServerError, fmt.Sprintf("Failed to get user: %v", err))
	}
	res := &domain.GetUserRes{
		Id:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
	}

	return res, nil
}

func NewGetUserSvc(repo port.UserRepo) domain.GetUserSvc {
	return &getUserSvc{
		repo: repo,
	}
}
