package domain

type Menu struct {
	RestaurantID int
	Items        []MenuItem
}

type MenuItem struct {
	ID     int
	MenuID int
	Name   string
	Price  float64
	// Photo  []byte
	Course string
}
