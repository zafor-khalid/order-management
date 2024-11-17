package main

import (
	"order-management/config"
	"order-management/internal/routers"
	"order-management/pkg/logger"

	"github.com/gin-gonic/gin"
)

func main() {
	logger.Initialize("INFO")

	// Load configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		logger.Error("Failed to load config: %v", err)
	}

	// Pass services to controllers if necessary
	router := gin.Default()
	routers.LoadRoutes(router)

	// Start server
	if err := router.Run(":" + cfg.Port); err != nil {
		logger.Error("Failed to start server: %v", err)
	}
}
