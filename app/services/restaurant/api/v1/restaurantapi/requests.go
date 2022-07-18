package restaurantapi

// RestaurantData contains information about the restaurant.
// swagger:model
type RestaurantData struct {
	// required: true
	Name string
	// required: true
	City string
	// required: true
	AcceptingOrders bool
	// required: true
	Address string
	// required: true
	Longitude float64
	// required: true
	Latitude float64
}

// MenuData contains information about the menu.
// swagger:model
type MenuData struct {
	// required: true
	MenuItems []MenuItemData
}

// MenuItemData contains information about the menu item.
// swagger:model
type MenuItemData struct {
	// required: true
	ID int
	// required: true
	Name string
	// required: true
	Price float64
	// required: true
	Course string // first/main/salad etc.
}
