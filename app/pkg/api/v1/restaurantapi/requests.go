package restaurantapi

type RestaurantData struct {
	Name   string
	City   string
	Street string
}

type MenuData struct {
	MenuItems []MenuItemData
}

type MenuItemData struct {
	ID   int
	Name string
	// Photo []byte
	Course string // first/main/salad etc.
}
