package communications

type ReturnOrderList struct {
	Orders []ReturnOrder
}

type ReturnOrder struct {
	OrderID      int
	FromUserID   int
	RestaurantID int
	OrderItems   []int
	Status       string
}
