package models

type ConsumerLocation struct {
	ID          uint64 `json:"id" yaml:"id"`
	ConsumerID  uint64 `json:"consumerID" yaml:"consumer_ID"`
	LocationAlt string `json:"locationAlt" yaml:"locationAlt"`
	LocationLat string `json:"locationLat" yaml:"locationLat"`
	Country     string `json:"country" yaml:"country"`
	City        string `json:"city" yaml:"city"`
	Region      string `json:"region" yaml:"region"`
	Street      string `json:"street" yaml:"street"`
	HomeNumber  string `json:"homeNumber" yaml:"homeNumber"`
	Floor       string `json:"floor" yaml:"floor"`
	Door        string `json:"door"`
}

// Fields will return all fields of this type
func (c *ConsumerLocation) Fields() []interface{} {
	return []interface{}{
		&c.ID,
		&c.ConsumerID,
		&c.LocationAlt,
		&c.LocationLat,
		&c.Country,
		&c.City,
		&c.Region,
		&c.Street,
		&c.HomeNumber,
		&c.Floor,
		&c.Door,
	}
}
