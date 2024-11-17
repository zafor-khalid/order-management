package repositories

import (
	"fmt"
	"order-management/config"
	"order-management/internal/models"
)

// CreateOrder saves the provided order to the database and returns the created order
func CreateOrder(order models.Order) (models.Order, error) {
	// Save the order to the database
	if err := config.DB.Create(&order).Error; err != nil {
		return models.Order{}, fmt.Errorf("failed to save order: %v", err)
	}

	// Return the created order
	return order, nil
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

// GetOrders retrieves a paginated list of orders with optional filtering by transfer status and archive status
func GetOrders(transferStatus string, archive string, limit int, offset int) ([]models.Order, int, error) {
	var orders []models.Order
	var total int64

	// Query builder
	query := config.DB.Model(&models.Order{})

	// Apply filters
	if transferStatus != "" {
		query = query.Where("status = ?", transferStatus)
	}
	if archive != "" {
		query = query.Where("archived = ?", archive)
	}

	// Count total records
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Fetch paginated results
	if err := query.Offset(offset).Limit(limit).Find(&orders).Error; err != nil {
		return nil, 0, err
	}

	return orders, int(total), nil
}

// GetOrderByID fetches an order from the database by its ID
func GetOrderByID(ConsignmentID string) (*models.Order, error) {
	var order models.Order
	err := config.DB.Where("consignment_id = ?", ConsignmentID).First(&order).Error
	if err != nil {
		return nil, err
	}
	return &order, nil
}

// UpdateOrderStatus updates the status of the given order in the database
func UpdateOrderStatus(order *models.Order) error {
	err := config.DB.Save(order).Error
	if err != nil {
		return err
	}
	return nil
}