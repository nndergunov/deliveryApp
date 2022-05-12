package domain

type SearchParameters struct {
	FromRestaurantID int
	Statuses         []string
	ExcludeStatuses  []string
}

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
