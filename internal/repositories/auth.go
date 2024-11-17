package repositories

import (
	"fmt"
	"order-management/config"
	"order-management/internal/models"
)

// CreateUser saves the provided user to the database and returns the created user
func CreateUser(user models.User) error {
	if err := config.DB.Create(&user).Error; err != nil {
		return  fmt.Errorf("failed to save order: %v", err)
	}
	return nil
}

func GetUserByEmail(email string) (models.User, error) {
	var user models.User
	if err := config.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return models.User{}, fmt.Errorf("failed to get user: %v", err)
	}
	return user, nil
}