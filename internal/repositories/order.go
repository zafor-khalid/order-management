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


func GetOrdersFromDB(transferStatus string, archive string, limit int, offset int) ([]models.Order, int, error) {
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