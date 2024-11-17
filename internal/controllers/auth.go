package controllers

import (
	"net/http"
	"order-management/internal/models"
	"order-management/internal/services"

	"github.com/gin-gonic/gin"
)

// SignUp handles user registration
// It expects a JSON payload containing the user details in the request body.
//
// @Summary Register a new user
// @Description Creates a new user account with the provided details
// @Tags auth
// @Accept json
// @Produce json
// @Param user body models.User true "User registration details"
// @Success 201 {object} gin.H "User created successfully"
// @Failure 400 {object} gin.H "Invalid request data"
// @Failure 500 {object} gin.H "Internal server error"
// @Router /api/v1/auth/signup [post]
//
func SignUp(c *gin.Context) {
	var user models.User
	
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Email or password is invalid",	
			"code": http.StatusBadRequest,
		})
		return
	}

	err := services.CreateUser(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to create user",	
			"code": http.StatusInternalServerError,
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "User created successfully",
		"code":    http.StatusCreated,
	})
}


// SignIn handles user authentication
// It expects a JSON payload containing the user credentials in the request body.
//
// @Summary Authenticate user
// @Description Authenticates a user and returns a JWT token
// @Tags auth
// @Accept json
// @Produce json
// @Param credentials body models.User true "User credentials"
// @Success 200 {object} gin.H "Successfully signed in with JWT token"
// @Failure 400 {object} gin.H "Invalid request data"
// @Failure 401 {object} gin.H "Invalid credentials"
// @Router /api/v1/auth/signin [post]
//

func SignIn(c *gin.Context) {
	var credentials models.User

	if err := c.ShouldBindJSON(&credentials); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Email or password is invalid",
			"code": http.StatusBadRequest,
		})
		return
	}

	token, err := services.AuthenticateUser(credentials.Email, credentials.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Invalid email or password",
			"code": http.StatusUnauthorized,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token": token,
		"message": "Successfully signed in",
		"code": http.StatusOK,
	})
}