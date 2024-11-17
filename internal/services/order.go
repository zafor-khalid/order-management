package services

import (
	"crypto/rand"
	"encoding/hex"
	"errors"
	"fmt"
	"order-management/internal/models"
	"order-management/internal/repositories"
	"regexp"
	"strings"
	"time"

	"github.com/go-playground/validator/v10"
)

func matchAddressToCity(address string) (int) {
	// Custom logic to map address to a city ID
	address = strings.ToLower(address)

	switch {
	case strings.Contains(address, "dhaka"):
		return 1
	case strings.Contains(address, "chittagong"):
		return 2
	default:
		return 0
	}
}


// generateConsignmentID creates a unique 12-character consignment ID
func generateConsignmentID() (string, error) {
	bytes := make([]byte, 6) // 6 bytes will give us 12 hex characters
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}

// CreateOrder calculates fees and saves the order
func CreateOrder(order models.Order) (models.Order, error) {
	// Calculate delivery fee
	order.DeliveryFee = calculateDeliveryFee(order.RecipientCity, order.ItemWeight)

	// Calculate COD fee
	order.CodFee = calculateCODFee(order.AmountToCollect)
	
	matchedCity := matchAddressToCity(order.RecipientAddress)

	
	// todo: area and city id should be matched from the address
	
	order.StoreID = 131172
	order.DeliveryType = 48
	order.ItemType = 2
	order.ItemQuantity = 1
	order.ItemWeight = 0.5
	order.RecipientZone = 1
	order.RecipientArea = 1
	order.Status = "Pending"
	order.RecipientCity = matchedCity
	order.CreatedAt = time.Now().Format("2024-05-23 14:05:34")
	order.UpdatedAt = time.Now().Format("2024-05-23 14:05:34")
	order.Archived = false
	
	consignmentID, err := generateConsignmentID()
	if err != nil {
		return models.Order{}, fmt.Errorf("failed to generate consignment ID: %w", err)
	}
	order.ConsignmentID = consignmentID
	
	// Save the order to the database
	createdOrder, err := repositories.CreateOrder(order)
	if err != nil {
		return models.Order{}, fmt.Errorf("failed to create order: %w", err)
	}

	return createdOrder, nil
}

func calculateDeliveryFee(city int, weight float64) float64 {
	if city == 1 {
		// Delivery fee logic for Dhaka
		if weight <= 0.5 {
			return 60
		} else if weight <= 1 {
			return 70
		}
		return 70 + (weight-1)*15
	}
	// Delivery fee logic for other cities
	return 100
}

func calculateCODFee(amountToCollect float64) float64 {
	return amountToCollect * 0.01
}


// GetValidationErrors parses binding or validation errors into a structured map
func GetValidationErrors(err error) map[string][]string {
	errors := make(map[string][]string)

	// Parse GIN binding errors
	if bindingErrors, ok := err.(validator.ValidationErrors); ok {
		for _, fieldErr := range bindingErrors {
			fieldName := fieldErr.Field()
			errorMessage := generateErrorMessage(fieldErr)
			errors[fieldName] = append(errors[fieldName], errorMessage)
		}
	}

	// Additional errors (e.g., custom validation logic)
	// Add more error details here if needed

	return errors
}

// generateErrorMessage generates user-friendly error messages
func generateErrorMessage(fieldErr validator.FieldError) string {
	switch fieldErr.Tag() {
	case "required":
		return "This field is required."
	case "min":
		return "Value is below the minimum allowed."
	case "max":
		return "Value exceeds the maximum allowed."
	case "e164":
		return "Invalid phone number format."
	case "len":
		return "Length must be exactly " + fieldErr.Param() + " characters."
	default:
		return "Invalid value."
	}
}

type ValidationError struct {
    Message string              `json:"message"`
    Type    string              `json:"type"`
    Code    int                 `json:"code"`
    Errors  map[string][]string `json:"errors"`
}

func ValidateOrderRequest(order *models.Order) error {
    errors := make(map[string][]string)

    if order.RecipientName == "" {
        errors["recipient_name"] = append(errors["recipient_name"], "The recipient name field is required.")
    }
    if order.RecipientPhone == "" {
        errors["recipient_phone"] = append(errors["recipient_phone"], "The recipient phone field is required.")
    } else {
        matched := regexp.MustCompile(`^(01)[3-9]{1}[0-9]{8}$`).MatchString(order.RecipientPhone)
        if !matched {
            errors["recipient_phone"] = append(errors["recipient_phone"], "Invalid phone number format.")
        }
    }
	if order.RecipientAddress == "" {
		errors["recipient_address"] = append(errors["recipient_address"], "The recipient address field is required.")
	}
	if order.AmountToCollect == 0 {
		errors["amount_to_collect"] = append(errors["amount_to_collect"], "The amount to collect field is required.")
	}
	


    if len(errors) > 0 {
        return &ValidationError{
            Message: "Please fix the given errors",
            Type:    "error",
            Code:    422,
            Errors:  errors,
        }
    }
    return nil
}

func (ve *ValidationError) Error() string {
    return ve.Message
}


func FetchOrders(transferStatus string, archive string, limit int, page int) ([]map[string]interface{}, int, error) {
	// Calculate offset for pagination
	offset := (page - 1) * limit

	// Query the database
	orders, total, err := repositories.GetOrdersFromDB(transferStatus, archive, limit, offset)
	if err != nil {
		return nil, 0, err
	}

	// Format orders for the JSON response
	formattedOrders := []map[string]interface{}{}
	for _, order := range orders {
		formattedOrders = append(formattedOrders, map[string]interface{}{
			"order_consignment_id": order.ConsignmentID,
			"order_created_at":     order.CreatedAt,
			"order_description":    order.ItemDescription,
			"merchant_order_id":    order.MerchantOrderID,
			"recipient_name":       order.RecipientName,
			"recipient_address":    order.RecipientAddress,
			"recipient_phone":      order.RecipientPhone,
			"order_amount":         order.AmountToCollect,
			"total_fee":            order.AmountToCollect+ order.DeliveryFee - order.CodFee - order.Discount,
			"instruction":          order.SpecialInstruction,
			"order_type_id":        order.DeliveryType,
			"cod_fee":              order.CodFee,
			"promo_discount":       order.Discount,
			"delivery_fee":         order.DeliveryFee,
			"order_status":         order.Status,
			"order_type":           order.DeliveryType,
			"item_type":            order.ItemType,
			"archived":             order.Archived,
		})
	}

	return formattedOrders, total, nil
}


// CancelOrder attempts to cancel an order and update its status in the database
func CancelOrder(orderID string) error {
	// Retrieve the order from the database
	order, err := repositories.GetOrderByID(orderID)
	if err != nil {
		return err
	}

	// Check if the order is in a cancellable state
	if order.Status != "Pending" {
		// If the order cannot be canceled, return an error
		return errors.New("order cannot be canceled")
	}

	// Update the order status to 'Canceled'
	order.Status = "Canceled"
	err = repositories.UpdateOrderStatus(order)
	if err != nil {
		return err
	}

	return nil
}