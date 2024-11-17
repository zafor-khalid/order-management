package main

import (
	"log"
	"order-management/config"
	"order-management/internal/routers"

	"github.com/gin-gonic/gin"
)

func main() {
	// Load configuration first
	err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	// Initialize the database using the loaded configuration
	config.InitDB()

	// Pass services to controllers if necessary
	router := gin.Default()
	routers.LoadRoutes(router)

	// Start server
	if err := router.Run(":8081"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
