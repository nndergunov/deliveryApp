package restaurantapi

type CreateRestaurant struct {
	Name   string
	City   string
	Street string
}

type UpdateRestaurant struct {
	Name   string
	City   string
	Street string
}

type CreateMenu struct {
	MenuItems []struct {
		Name string
		// Photo []byte
		Course string // first/main/salad etc.
	}
}

type UpdateMenu struct {
	MenuItems []struct {
		Name string
		// Photo []byte
		Course string // first/main/salad etc.
	}
}
