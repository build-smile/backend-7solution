package middleware

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"testing"
	"time"
)

func TestGenerateJWT(t *testing.T) {
	// Create a new token
	token := jwt.New(jwt.SigningMethodHS256)

	// Set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["sub"] = "user@example.com"
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix() // Token expires in 24 hours

	// Sign the token with the secret key
	tokenString, err := token.SignedString([]byte("12345"))
	if err != nil {
		fmt.Print(err)
	}

	fmt.Print(tokenString)
}
