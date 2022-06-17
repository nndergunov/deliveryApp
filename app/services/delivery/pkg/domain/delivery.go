package domain

type EstimateDeliveryRequest struct {
	ConsumerID   int
	RestaurantID int
}

type EstimateDeliveryResponse struct {
	Time string
	Cost float64
}

type DeliveryCost struct {
	Cost float64
}

type AssignOrder struct {
	OrderID   int
	CourierID int
}

type Coordinate struct {
	Lat float64
	Lon float64
}
