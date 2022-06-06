package deliveryapi

type EstimateDeliveryResponse struct {
	Time string  `json:"time,omitempty"`
	Cost float64 `json:"cost,omitempty"`
}

type AssignOrderResponse struct {
	OrderID   int `json:"order_id"`
	CourierID int `json:"courier_id"`
}
