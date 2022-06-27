package deliveryapi

type EstimateDeliveryRequest struct {
	ConsumerID   int `json:"consumer_id"`
	RestaurantID int `json:"restaurant_id"`
}

type AssignOrderRequest struct {
	FromUserID   int `json:"from_user_id"`
	RestaurantID int `json:"restaurant_id"`
}
