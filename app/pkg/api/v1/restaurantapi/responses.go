package restaurantapi

type ReturnRestaurant struct {
	ID      int
	Name    string
	City    string
	Address string
}

type ReturnRestaurantList struct {
	List []ReturnRestaurant
}

type ReturnMenu struct {
	RestaurantID int
	Items        []ReturnMenuItem
}

type ReturnMenuItem struct {
	ID     int
	Name   string
	Course string
	// Photo []byte
}
