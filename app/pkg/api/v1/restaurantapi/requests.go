package restaurantapi

type RestaurantData struct {
	Name            string
	City            string
	AcceptingOrders bool
	Address         string
}

type MenuData struct {
	MenuItems []MenuItemData
}

type MenuItemData struct {
	ID    int
	Name  string
	Price float64
	// Photo []byte
	Course string // first/main/salad etc.
}
