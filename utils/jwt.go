package utils

import (
	"github.com/build-smile/backend-7solution/infrastructure"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

// Claims define the structure of JWT claims
type Claims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

// GenerateJWT generates a JWT token for a user
func GenerateJWT(username string) (string, string, error) {
	expirationTime := time.Now().Add(48 * time.Hour)
	claims := &Claims{
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	// Create a new token with the claims and sign it with the secret key
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(infrastructure.CFG.Jwt.SecretKey)
	if err != nil {
		return "", "", err
	}

	// Refresh Token (longer-lived)
	refreshExpirationTime := time.Now().Add(7 * 24 * time.Hour)
	refreshClaims := &Claims{
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(refreshExpirationTime),
		},
	}
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims)
	refreshTokenString, err := refreshToken.SignedString(infrastructure.CFG.Jwt.SecretKey)
	if err != nil {
		return "", "", err
	}

	return tokenString, refreshTokenString, nil
}
