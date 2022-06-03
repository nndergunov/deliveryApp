package domain

type DeliveryDistanceLocation struct {
	FromLocation *Location
	ToLocation   *Location
}

type DeliveryTime struct {
	Time string
}

type DeliveryCost struct {
	Cost float64
}

type AssignedCourier struct {
	OrderID   int
	CourierID int
}
