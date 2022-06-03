package domain

type Location struct {
	Latitude   string
	Longitude  string
	Country    string
	City       string
	Region     string
	Street     string
	HomeNumber string
	Floor      string
	Door       string
}

// Coord represents a geographic coordinate
type Coord struct {
	Lat float64
	Lon float64
}
