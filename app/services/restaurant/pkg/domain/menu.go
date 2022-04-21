package domain

type Menu struct {
	RestaurantID int
	Items        map[int]MenuItem
}

type MenuItem struct {
	ID   int
	Name string
	// Photo  []byte
	Course string
}
