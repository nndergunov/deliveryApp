package communications

// OrderFilters struct lets user determine which orders they want to see.
type OrderFilters struct {
	FromRestaurantID *int
	Statuses         []string
	ExcludeStatuses  []string
}

type PostOrder struct {
	OrderData   OrderData
	UserAccount int
}

type OrderData struct {
	FromUserID     int
	RestaurantID   int
	OrderItems     []int
	PaymentHashKey string
}

type OrderStatusData struct {
	Status string
}
