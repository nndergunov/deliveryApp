package deliveryapi

// EstimateDeliveryResponse contains information about the estimate values.
// swagger:model
type EstimateDeliveryResponse struct {
	Time string  `json:"time,omitempty"`
	Cost float64 `json:"cost,omitempty"`
}

// AssignOrderResponse contains information about assign order.
// swagger:model
type AssignOrderResponse struct {
	OrderID   int `json:"order_id"`
	CourierID int `json:"courier_id"`
}
