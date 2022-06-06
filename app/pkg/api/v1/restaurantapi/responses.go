package restaurantapi

type ReturnRestaurant struct {
	ID              int
	Name            string
	AcceptingOrders bool
	City            string
	Address         string
}

type ReturnRestaurantList struct {
	List []ReturnRestaurant
}

type ReturnMenu struct {
	RestaurantID int
	MenuItems    []ReturnMenuItem
}

type ReturnMenuItem struct {
	ID     int
	Name   string
	Price  float64
	Course string
	// Photo []byte
}
