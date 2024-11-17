package routers

import (
	"order-management/internal/controllers"
	"order-management/internal/middlewares"

	"github.com/gin-gonic/gin"
)

const relativePath = "/api/v1"

func LoadRoutes(router *gin.Engine) {
	// Root-level health check
	router.GET("/health", controllers.HealthCheck)

	api := router.Group(relativePath)
	{
		//	Auth routes
		auth := api.Group("/auth")
		{
			auth.POST("/signup", controllers.SignUp)
			auth.POST("/signin", controllers.SignIn)
		}
		
		//	Orders routes
		orders := api.Group("/orders").Use(middlewares.AuthMiddleware)
		{
			orders.POST("", controllers.CreateOrder) 
			orders.GET("/all", controllers.GetOrders)
			orders.PUT("/:consignment_id/cancel", controllers.CancelOrder)
		}
	}
}
