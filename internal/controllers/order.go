package controllers

import (
	"net/http"
	"order-management/internal/models"
	"order-management/internal/services"

	"github.com/gin-gonic/gin"
)


func CreateOrder(c *gin.Context) {
	var order models.Order

	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"message": "Please fix the given errors",
			"type":    "error",
			"code":    422,
			"errors":  err.Error(),
		})
		return
	}

	// Validate the order request
	if err := services.ValidateOrderRequest(&order); err != nil {
		if validationErr, ok := err.(*services.ValidationError); ok {
			c.JSON(http.StatusUnprocessableEntity, gin.H{
				"message": validationErr.Message,
				"type":    validationErr.Type,
				"code":    validationErr.Code,
				"errors":  validationErr.Errors,
			})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "Validation failed",
				"type":    "error",
				"code":    http.StatusInternalServerError,
			})
		}
		return
	}

	// Call the service to create the order
	createdOrder, err := services.CreateOrder(order)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to create order",
			"type":    "error",
			"code":    http.StatusInternalServerError,
		})
		return
	}

	// Send success response
	c.JSON(http.StatusOK, gin.H{
		"message": "Order Created Successfully",
		"type":    "success",
		"code":    http.StatusOK,
		"data": gin.H{
			"consignment_id":    createdOrder.ConsignmentID,
			"merchant_order_id": createdOrder.MerchantOrderID,
			"order_status":      "Pending",
			"delivery_fee":      createdOrder.DeliveryFee,
		},
	})
}




// GetOrders fetches all orders
func GetOrders(c *gin.Context) {
	// page := c.DefaultQuery("page", "1")
	// limit := c.DefaultQuery("limit", "10")

	// // Call the service to fetch orders
	// orders, err := services.GetOrders(page, limit)
	// if err != nil {
	// 	c.JSON(http.StatusInternalServerError, gin.H{
	// 		"message": "Failed to fetch orders",
	// 		"type":    "error",
	// 		"code":    http.StatusInternalServerError,
	// 	})
	// 	return
	// }

	// // Send success response
	// c.JSON(http.StatusOK, gin.H{
	// 	"message": "Orders fetched successfully",
	// 	"data":    orders,
	// })
}

// CancelOrder cancels an order by its consignment ID
func CancelOrder(c *gin.Context) {
	// consignmentID := c.Param("consignment_id")

	// // Call the service to cancel the order
	// err := services.CancelOrder(consignmentID)
	// if err != nil {
	// 	c.JSON(http.StatusInternalServerError, gin.H{
	// 		"message": "Failed to cancel order",
	// 		"type":    "error",
	// 		"code":    http.StatusInternalServerError,
	// 	})
	// 	return
	// }

	// // Send success response
	// c.JSON(http.StatusOK, gin.H{
	// 	"message": "Order cancelled successfully",
	// })
}