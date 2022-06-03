package deliveryapi

type DeliveryTimeResponse struct {
	Time string `json:"time,omitempty"`
}

type DeliveryCostResponse struct {
	Cost float64 `json:"cost,omitempty"`
}

type DeliveryAssignedCourierResponse struct {
	OrderID   int `json:"order_id"`
	CourierID int `json:"courier_id"`
}
