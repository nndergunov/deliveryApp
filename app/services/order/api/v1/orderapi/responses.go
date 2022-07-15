package orderapi

// ReturnOrderList contains information about all the requested orders.
// swagger:model
type ReturnOrderList struct {
	// all the requested orders
	Orders []ReturnOrder
}

// ReturnOrder contains order information.
// swagger:model
type ReturnOrder struct {
	OrderID      int
	FromUserID   int
	RestaurantID int
	OrderItems   []int
	Status       string
}
