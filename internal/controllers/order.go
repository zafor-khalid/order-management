package controllers

import (
	"net/http"
	"order-management/internal/models"
	"order-management/internal/services"
	"order-management/utils"

	"github.com/gin-gonic/gin"
)

// CreateOrder handles the creation of a new order.
// It expects a JSON payload containing the order details in the request body.
//
// @Summary Create a new order
// @Description Creates a new order with the provided details
// @Tags orders
// @Accept json
// @Produce json
// @Param order body models.Order true "Order details"
// @Success 200 {object} gin.H "Order created successfully"
// @Failure 422 {object} gin.H "Validation errors"
// @Failure 500 {object} gin.H "Internal server error"
// @Router /api/v1/orders [post]

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

// GetOrders handles the GET /api/v1/orders/all endpoint.
// It retrieves a paginated list of orders with optional filtering.
//
// Query Parameters:
//   - transfer_status: Filter orders by transfer status
//   - archive: Filter orders by archive status
//   - limit: Number of orders per page (default: 10)
//   - page: Page number (default: 1)
//
// Returns:
//   - 200 OK: Successfully retrieved orders with pagination details
//   - 500 Internal Server Error: If there was an error fetching the orders

func GetOrders(c *gin.Context) {
	transferStatus := c.Query("transfer_status")
	archive := c.Query("archive")
	limit := c.DefaultQuery("limit", "10")
	page := c.DefaultQuery("page", "1")

	limitInt, pageInt := utils.ParsePaginationParams(limit, page)
	
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
    lastPage := utils.CalculateLastPage(total, limitInt)

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
			"total_in_page": len(orders),
			"last_page":     lastPage,
		},
	}
	c.JSON(http.StatusOK, response)
}


// CancelOrder handles the cancellation of an order
//
// Parameters:
//   - consignment_id: The ID of the order to be cancelled (URL parameter)
//
// Returns:
//   - 200 OK: Successfully cancelled the order
//   - 400 Bad Request: If the order cannot be cancelled

func CancelOrder(c *gin.Context) {
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
