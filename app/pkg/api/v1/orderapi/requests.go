package orderapi

type OrderData struct {
	FromUserID   int
	RestaurantID int
	OrderItems   []int
}

type OrderStatusData struct {
	OrderID int
	Status  string
}
