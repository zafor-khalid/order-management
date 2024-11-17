package utils

import (
	"errors"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// ValidateToken validates the JWT token and returns an error if the token is invalid
func ValidateToken(token string) (string, error) {
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	if err != nil {
		return "", err
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok || !parsedToken.Valid {
		return "", jwt.ErrInvalidKey
	}

	if exp, ok := claims["exp"].(float64); !ok || float64(time.Now().Unix()) > exp {
		return "", errors.New("token expired")
	}

	sub, ok := claims["sub"].(string)
	if !ok {
		return "", errors.New("invalid subject claim")
	}
	
	return sub, nil
}