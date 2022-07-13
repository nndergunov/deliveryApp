package deliveryapi

// EstimateDeliveryRequest contains information about the estimate values.
// swagger:model
type EstimateDeliveryRequest struct {
	// required: true
	ConsumerID int `json:"consumer_id"`
	// required: true
	RestaurantID int `json:"restaurant_id"`
}

// AssignOrderRequest contains information about assign order.
// swagger:model
type AssignOrderRequest struct {
	// required: true
	FromUserID int `json:"from_user_id"`
	// required: true
	RestaurantID int `json:"restaurant_id"`
}
