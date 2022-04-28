package domain

type Order struct {
	OrderID      int
	FromUserID   int
	RestaurantID int
	OrderItems   []int
	Status       string
}
