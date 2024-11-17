package services

import (
	"fmt"
	"order-management/internal/models"
	"order-management/internal/repositories"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

// CreateUser saves the provided user to the database and returns the created user
func CreateUser(user models.User) error {
	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("failed to process password: %w", err)
	}
	user.Password = string(hashedPassword)
	
	// Save the user to the database
	err = repositories.CreateUser(user)
	if err != nil {
		return fmt.Errorf("failed to create user: %w", err)
	}
	return nil
}

// AuthenticateUser checks if the provided email and password match a user in the databases
func AuthenticateUser(email, password string) (string, error) {
	user, err := repositories.GetUserByEmail(email)
	if err != nil {
		return "", fmt.Errorf("failed to get user: %w", err)
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "", fmt.Errorf("invalid password: %w", err)
	}
	
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24).Unix(), // Default to 24 hour expiry
	}).SignedString([]byte(os.Getenv("JWT_SECRET")))
	
	if err != nil {
		return "", fmt.Errorf("failed to generate token: %w", err)
	}
	
	return token, nil
}
