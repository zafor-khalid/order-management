package models

// Order represents the structure of an order in the system.
// It contains all the necessary information about a delivery order,
// including merchant details, recipient information, item details,
// and various fees and statuses.
//
// Fields:
//   - ID: Unique identifier for the order in the database
//   - StoreID: ID of the store/merchant creating the order
//   - MerchantOrderID: Merchant's internal reference ID for the order
//   - RecipientName: Name of the delivery recipient (required)
//   - RecipientPhone: Contact number of the recipient (required)
//   - RecipientAddress: Delivery address (required)
//   - RecipientCity/Zone/Area: Geographic location details
//   - DeliveryType: Type of delivery service requested
//   - ItemType: Category or type of item being delivered
//   - SpecialInstruction: Additional delivery instructions
//   - ItemQuantity: Number of items in the order
//   - ItemWeight: Weight of the items in the order
//   - AmountToCollect: COD amount to collect from recipient (must be > 0)
//   - ItemDescription: Description of items being delivered
//   - DeliveryFee: Fee charged for delivery service
//   - CodFee: Fee charged for cash-on-delivery service
//   - Status: Current status of the order (e.g., Pending, Canceled)
//   - ConsignmentID: Unique tracking ID for the order
//   - Discount: Any discount applied to the order
//   - CreatedAt: Timestamp of order creation
//   - UpdatedAt: Timestamp of last order update
//   - Archived: Flag indicating if the order is archived

type Order struct {
	ID                 uint    `gorm:"primaryKey"`
	StoreID            int     `json:"store_id"`
	MerchantOrderID    string  `json:"merchant_order_id"`
	RecipientName      string  `json:"recipient_name" binding:"required"`
	RecipientPhone     string  `json:"recipient_phone" binding:"required"`
	RecipientAddress   string  `json:"recipient_address" binding:"required"`
	RecipientCity      int     `json:"recipient_city"`
	RecipientZone      int     `json:"recipient_zone"`
	RecipientArea      int     `json:"recipient_area"`
	DeliveryType       int     `json:"delivery_type"`
	ItemType           int     `json:"item_type"`
	SpecialInstruction string  `json:"special_instruction"`
	ItemQuantity       int     `json:"item_quantity"`
	ItemWeight         float64 `json:"item_weight"`
	AmountToCollect    float64 `json:"amount_to_collect" binding:"required,gt=0"`
	ItemDescription    string  `json:"item_description"`
	DeliveryFee        float64 `json:"delivery_fee"`
	CodFee             float64 `json:"cod_fee"`
	Status             string  `json:"status"`
	ConsignmentID      string  `json:"consignment_id"`
	Discount           float64 `json:"discount"`
	CreatedAt          string  `json:"created_at"`
	UpdatedAt          string  `json:"updated_at"`
	Archived           bool    `json:"archived"`
}
