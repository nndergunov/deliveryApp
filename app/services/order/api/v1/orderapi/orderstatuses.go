package orderapi

const (
	// Sending indicates that the order is being sent to the restaurant.
	Sending = "Sending"
	// Accepted indicates that the order is accepted by the restaurant.
	Accepted = "Accepted by the restaurant"
	// Cooking indicates that the order is being prepared by the restaurant.
	Cooking = "Cooking"
	// SearchingForCourier indicates that the restaurant is searching for the courier.
	SearchingForCourier = "Searching for courier"
	// HandingToTheCourier indicates that the order is being handed to the courier.
	HandingToTheCourier = "Handing the order to the courier"
	// InDelivery indicates that the order is being delivered.
	InDelivery = "Delivering"
	// Complete indicates that the order had been delivered.
	Complete = "Delivered"
)
