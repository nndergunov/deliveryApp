package order

type CreateOrder struct {
	FromUserID       int
	RestaurantID     int
	DeliveryLocation string
	SpecialRequests  string
	Order            map[string]int
}

type OrderStatusUpdate struct {
	OrderID     int
	OrderStatus string
}
