package routers

import (
	"order-management/internal/controllers"

	"github.com/gin-gonic/gin"
)

const relativePath = "/api/v1"

// LoadRoutes now accepts services as parameters and passes them to controllers
func LoadRoutes(router *gin.Engine) {



	// Root-level health check
	router.GET("/health", controllers.HealthCheck)

	// Public routes group
	public := router.Group(relativePath)
	{
		// Auth routes
		// auth := public.Group("/auth")
		// {
		// 	auth.POST("/login", authController.Login)
		// }
		
		orders := public.Group("/orders")
		{
			orders.GET("", controllers.GetOrders)
			orders.POST("", controllers.CreateOrder) 
			orders.PUT("/:consignment_id/cancel", controllers.CancelOrder)
		}
	}
	


	// Protected routes group with JWT middleware
	// protected := router.Group(relativePath)
	// protected.Use(middlewares.JWTMiddleware())
	// {
	// 	// Order routes
	// 	orders := protected.Group("/orders")
	// 	{
	// 		orders.GET("", orderController.GetAll)
	// 		orders.POST("", orderController.Create)
	// 		orders.PUT("/:id/cancel", orderController.Cancel)
	// 	}

		
	// }
}
