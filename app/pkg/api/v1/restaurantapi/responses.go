package restaurantapi

type RestaurantList struct {
	List []struct {
		Name   string
		City   string
		Street string
	}
}

type ReturnMenu struct {
	ItemsByCourse map[string][]struct {
		ID   int
		Name string
		// Photo []byte
	}
}
