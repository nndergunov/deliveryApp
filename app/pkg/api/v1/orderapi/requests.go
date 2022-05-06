package orderapi

type OrderData struct {
	FromUserID   int
	RestaurantID int
	OrderItems   []int
}

type OrderStatusData struct {
	Status string
}
