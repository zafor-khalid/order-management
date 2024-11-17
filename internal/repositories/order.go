package repositories

import (
	"fmt"
	"order-management/config"
	"order-management/internal/models"
)

// CreateOrder saves the provided order to the database and returns the created order
func CreateOrder(order models.Order) (models.Order, error) {
	fmt.Printf("Creating order with details: %+v\n", order)

	// Save the order to the database
	if err := config.DB.Create(&order).Error; err != nil {
		return models.Order{}, fmt.Errorf("failed to save order: %v", err)
	}

	// Return the created order
	return order, nil
}

// GetOrders retrieves orders from the database with pagination
func GetOrders(page, limit int) ([]models.Order, error) {
	var orders []models.Order

	// Fetch the orders with pagination
	if err := config.DB.Offset((page - 1) * limit).Limit(limit).Find(&orders).Error; err != nil {
		return nil, fmt.Errorf("failed to fetch orders: %v", err)
	}

	// Return the fetched orders
	return orders, nil
}

// CancelOrder updates the order status to 'Cancelled' using the consignment ID
func CancelOrder(consignmentID string) error {
	var order models.Order

	// Find the order by consignment ID
	if err := config.DB.Where("consignment_id = ?", consignmentID).First(&order).Error; err != nil {
		return fmt.Errorf("order not found: %v", err)
	}

	// Update the order status to 'Cancelled'
		order.Status = "Cancelled"
	if err := config.DB.Save(&order).Error; err != nil {
		return fmt.Errorf("failed to cancel order: %v", err)
	}

	// Return success
	return nil
}
