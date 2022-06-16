package domain

type Restaurant struct {
	ID              int
	Name            string
	AcceptingOrders bool
	City            string
	Address         string
	Longitude       float64
	Latitude        float64
}
