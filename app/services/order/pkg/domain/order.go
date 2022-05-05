package domain

type Order struct {
	OrderID      int
	FromUserID   int
	RestaurantID int
	OrderItems   []int
	Status       OrderStatus
}

type OrderStatus struct {
	OrderID int
	Status  string
}
