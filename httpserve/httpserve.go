package httpserve

import (
	"fmt"
	"github.com/build-smile/backend-7solution/middleware"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func Run() {
	r := gin.Default()
	m := middleware.NewMiddleware(viper.GetString("jwt.secret-key"))
	r.Use(middleware.ErrorHandlerMiddleware)
	r.Use(gin.Recovery())
	r.Use(m.JWTMiddleware())

	bindPing(r)
	bindGetUsers(r)
	bindGetUser(r)
	bindCreateUser(r)
	bindRegister(r)
	bindLogin(r)
	bindUpdateUser(r)
	bindDeleteUser(r)

	port := viper.GetString("app.port")
	if port == "" {
		port = "1323"
	}
	err := r.Run(fmt.Sprintf(":%s", port))
	if err != nil {
		return
	} // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
