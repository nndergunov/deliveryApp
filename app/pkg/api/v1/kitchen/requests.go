package kitchen

type CreateRestaurant struct {
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
