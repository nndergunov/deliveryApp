package restaurantapi

type ReturnRestaurantList struct {
	List []struct {
		ID      int
		Name    string
		City    string
		Address string
	}
}

type ReturnMenu struct {
	ItemsByCourse map[string][]struct {
		ID   int
		Name string
		// Photo []byte
	}
}
