package orderapi

type OrderData struct {
	FromUserID   int
	RestaurantID int
	OrderItems   []int
	Status       string
}
