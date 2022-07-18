// Package orderapi is used to store communication information for order service.
package orderapi

// OrderFilters struct lets user determine which orders they want to see.
// swagger:model
type OrderFilters struct {
	// required: false
	FromRestaurantID *int
	// required: false
	Statuses []string
	// required: false
	ExcludeStatuses []string
}

// OrderData contains information about created order.
// swagger:model
type OrderData struct {
	// required: true
	FromUserID int
	// required: true
	RestaurantID int
	// required: true
	OrderItems []int
}

// OrderStatusData is used by admins to update current order status.
// swagger:model
type OrderStatusData struct {
	// required: true
	Status string
}
