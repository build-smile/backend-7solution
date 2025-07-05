package httpserve

import (
	"github.com/build-smile/backend-7solution/internal/core/services"
	"github.com/build-smile/backend-7solution/internal/handlers"
	"github.com/build-smile/backend-7solution/internal/repositories"
	"github.com/gin-gonic/gin"
	"net/http"
)

func bindPing(r *gin.Engine) gin.IRoutes {
	return r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
}

func bindGetUsers(r *gin.Engine) gin.IRoutes {
	repo := repositories.NewUserRepo()
	svc := services.NewGetUsersSvc(repo)
	hdl := handlers.NewGetUsersHdl(svc)
	return r.GET("/users", hdl.Handle)
}

func bindGetUser(r *gin.Engine) gin.IRoutes {
	repo := repositories.NewUserRepo()
	svc := services.NewGetUserSvc(repo)
	hdl := handlers.NewGetUserHdl(svc)
	return r.GET("/user/:id", hdl.Handle)
}

func bindCreateUser(r *gin.Engine) gin.IRoutes {
	repo := repositories.NewUserRepo()
	svc := services.NewCreateUserSvc(repo)
	hdl := handlers.NewCreateUserHdl(svc)
	return r.POST("/user", hdl.Handle)
}

func bindRegister(r *gin.Engine) gin.IRoutes {
	repo := repositories.NewUserRepo()
	svc := services.NewRegisterUserSvc(repo)
	hdl := handlers.NewRegisterUserHdl(svc)
	return r.POST("/register", hdl.Handle)
}
func bindLogin(r *gin.Engine) gin.IRoutes {
	repo := repositories.NewUserRepo()
	svc := services.NewLoginUserSvc(repo)
	hdl := handlers.NewLoginUserHdl(svc)
	return r.POST("/login", hdl.Handle)
}
func bindUpdateUser(r *gin.Engine) gin.IRoutes {
	repo := repositories.NewUserRepo()
	svc := services.NewUpdateUserSvc(repo)
	hdl := handlers.NewUpdateUserHdl(svc)
	return r.PATCH("/user/:id", hdl.Handle)

}
func bindDeleteUser(r *gin.Engine) gin.IRoutes {
	repo := repositories.NewUserRepo()
	svc := services.NewDeleteUserSvc(repo)
	hdl := handlers.NewDeleteUserHdl(svc)
	return r.DELETE("/user/:id", hdl.Handle)
}
