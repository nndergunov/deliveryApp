package deliveryapi

type DeliveryTimeRequest struct {
	FromLocation *Location `json:"from_location"`
	ToLocation   *Location `json:"to_location"`
}

type DeliveryCostRequest struct {
	FromLocation *Location `json:"from_location"`
	ToLocation   *Location `json:"to_location"`
}

type Location struct {
	Latitude   string `json:"latitude" yaml:"latitude"`
	Longitude  string `json:"longitude" yaml:"longitude"`
	Country    string `json:"country" yaml:"country"`
	City       string `json:"city" yaml:"city"`
	Region     string `json:"region" yaml:"region"`
	Street     string `json:"street" yaml:"street"`
	HomeNumber string `json:"home_number" yaml:"home_number"`
	Floor      string `json:"floor" yaml:"floor"`
	Door       string `json:"door" yaml:"door"`
}

type OrderRequest struct {
	FromUserID   int `json:"from_user_id"`
	RestaurantID int `json:"restaurant_id"`
}
