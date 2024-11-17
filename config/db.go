package config

import (
	"fmt"
	"log"
	"order-management/internal/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// DB is the global variable for the database connection
var DB *gorm.DB

// InitDB initializes the database connection and runs migrations
func InitDB() {
	// Construct the DSN (Data Source Name) for PostgreSQL using the global AppConfig
	// dsn := fmt.Sprintf(
	// 	"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
	// 	AppConfig.DBHost,
	// 	AppConfig.DBUser,
	// 	AppConfig.DBPassword,
	// 	AppConfig.DBName,
	// 	AppConfig.DBPort,
	// )

	// Open the connection to the database
	var err error
	 dsn := "host=localhost user=user password=password dbname=order_db port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	fmt.Println(dsn)

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	log.Println("Successfully connected to the database.")

	// Auto Migrate models (this can be expanded as needed)
	if err := DB.AutoMigrate(&models.Order{}, &models.User{}); err != nil {
		log.Fatalf("Error migrating models: %v", err)
	}

	log.Println("Database migration completed successfully.")
}
