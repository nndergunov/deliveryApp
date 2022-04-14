package kitchen

type RestaurantList struct {
	List []struct {
		Name   string
		City   string
		Street string
	}
}

type Menu struct {
	ItemsByCourse map[string][]struct {
		Name string
		// Photo []byte
	}
}
