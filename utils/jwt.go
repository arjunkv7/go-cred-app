package utils

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var secretKey = []byte("secret-key")

func GenerateJwtToken(userId string, firstName string) (string, error) {

	claims := jwt.MapClaims{
		"userId":    userId,
		"firstName": firstName,
		"exp":       time.Now().Add(time.Hour * 24).Unix(), // 24 hour expiration
		"iat":       time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(secretKey)
}


func VerifyToken(tokenString string) (*jwt.MapClaims, error) {
    token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
        return secretKey, nil
    })

	if err != nil {
		return  nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return &claims, nil
	}

	return  nil, fmt.Errorf("invalid token")
}