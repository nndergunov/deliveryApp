package domain

type Menu struct {
	RestaurantID int
	Items        []MenuItem
}

type MenuItem struct {
	ID     int
	MenuID int
	Name   string
	// Photo  []byte
	Course string
}
