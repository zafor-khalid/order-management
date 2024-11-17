package controllers

import (
	"net/http"
	"order-management/internal/models"
	"order-management/internal/services"
	"strconv"

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




func GetOrders(c *gin.Context) {
	// Parse query parameters
	transferStatus := c.Query("transfer_status")
	archive := c.Query("archive")
	limit := c.DefaultQuery("limit", "10")
	page := c.DefaultQuery("page", "1")

	// Convert limit and page to integers
	limitInt, err := strconv.Atoi(limit)
	if err != nil || limitInt <= 0 {
		limitInt = 10 // Default to 10 if invalid
	}

	pageInt, err := strconv.Atoi(page)
	if err != nil || pageInt <= 0 {
		pageInt = 1 // Default to 1 if invalid
	}

	// Fetch orders from the service layer
	orders, total, err := services.FetchOrders(transferStatus, archive, limitInt, pageInt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to fetch orders",
			"type":    "error",
			"code":    http.StatusInternalServerError,
		})
		return
	}

	// Calculate pagination details
	lastPage := (total + limitInt - 1) / limitInt 
	totalInPage := len(orders)

	// Prepare the response
	response := gin.H{
		"message": "Orders successfully fetched.",
		"type":    "success",
		"code":    http.StatusOK,
		"data": gin.H{
			"data":          orders,
			"total":         total,
			"current_page":  pageInt,
			"per_page":      limitInt,
			"total_in_page": totalInPage,
			"last_page":     lastPage,
		},
	}
	c.JSON(http.StatusOK, response)
}


// CancelOrder handles the cancellation of an order
func CancelOrder(c *gin.Context) {
	// Retrieve the order ID from the URL parameter
	consignmentID := c.Param("consignment_id")

	// Call the service to cancel the order
	err := services.CancelOrder(consignmentID)
	if err != nil {
		// If cancellation fails, return a 400 error
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Please contact cx to cancel order",
			"type":    "error",
			"code":    400,
		})
		return
	}

	// If cancellation is successful, return a 200 success response
	c.JSON(http.StatusOK, gin.H{
		"message": "Order Cancelled Successfully",
		"type":    "success",
		"code":    200,
	})
}
