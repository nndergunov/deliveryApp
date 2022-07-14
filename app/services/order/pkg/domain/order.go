// Package domain is used to store service-wide objects.
package domain

// SearchParameters contains information about order search filters.
type SearchParameters struct {
	FromRestaurantID *int
	Statuses         []string
	ExcludeStatuses  []string
}

// Order is collection of data about the order.
type Order struct {
	OrderID      int
	FromUserID   int
	RestaurantID int
	OrderItems   []int
	Status       OrderStatus
}

// OrderStatus contains the current status of the order.
type OrderStatus struct {
	OrderID int
	Status  string
}
