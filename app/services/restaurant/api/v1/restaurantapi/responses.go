package restaurantapi

// ReturnRestaurant contains information about the requested restaurant.
// swagger:model
type ReturnRestaurant struct {
	ID              int
	Name            string
	AcceptingOrders bool
	City            string
	Address         string
	Longitude       float64
	Latitude        float64
}

// ReturnRestaurantList contains data about all the requested restaurants.
// swagger:model
type ReturnRestaurantList struct {
	List []ReturnRestaurant
}

// ReturnMenu contains information about the menu from requested restaurant.
// swagger:model
type ReturnMenu struct {
	RestaurantID int
	MenuItems    []ReturnMenuItem
}

// ReturnMenuItem contains information about requested menu item.
// swagger:model
type ReturnMenuItem struct {
	ID     int
	Name   string
	Price  float64
	Course string
}
