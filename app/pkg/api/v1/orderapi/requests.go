package orderapi

// OrderFilters struct lets user determine which orders they want to see.
type OrderFilters struct {
	FromRestaurantID *int
	Statuses         []string
	ExcludeStatuses  []string
}

type OrderData struct {
	FromUserID   int
	RestaurantID int
	OrderItems   []int
}

type OrderStatusData struct {
	Status string
}
