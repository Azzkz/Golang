package utils

import (
	"github.com/golang-jwt/jwt/v4"
	"time"
)

var jwtKey = []byte("your_secret_key")

func GenerateJWT(userID string) (string, error) {
	claims := &jwt.StandardClaims{
		Subject:   userID,
		ExpiresAt: time.Now().Add(24 * time.Hour).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}

func ValidateJWT(tokenStr string) bool {
	_, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	return err == nil
}
