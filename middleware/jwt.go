package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
)

func (m Middleware) JWTMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if skipVerifyJwt(c) {
			return
		}
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte(m.secretKey), nil
		})

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		if !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token claims"})
			c.Abort()
			return
		}

		fmt.Println("Claims:", claims)

		c.Next()
	}
}
func skipVerifyJwt(c *gin.Context) bool {
	if (c.FullPath() == "/login" || c.FullPath() == "/register" || c.FullPath() == "") &&
		(c.Request.Method == "POST" || c.Request.Method == "OPTIONS") {
		return true
	}
	return false
}
