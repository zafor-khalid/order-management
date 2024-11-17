package middlewares

import (
	"fmt"
	"log"
	"net/http"
	"order-management/config"
	"order-management/internal/models"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// AuthMiddleware is a middleware that checks if the user is authenticated
func AuthMiddleware(c *gin.Context) {
	tokenString := c.GetHeader("Authorization")
	if tokenString == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized. Token is missing"})
			c.Abort()
			return
		}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	if err != nil {
		log.Fatal(err)
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		if exp, ok := claims["exp"].(float64); !ok || float64(time.Now().Unix()) > exp {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized. Token has expired"})
			c.Abort()
			return
		}
		
		var user models.User
		config.DB.First(&user, claims["sub"])
		if user.ID == 0 {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized. Token is invalid"})
			c.Abort()
			return
		}
		c.Next()

	} else {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized. Token is invalid"})
		c.Abort()
	}
}