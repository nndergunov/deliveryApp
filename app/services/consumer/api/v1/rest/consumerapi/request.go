package consumerapi

// NewConsumerRequest contains information about the consumer.
// swagger:model
type NewConsumerRequest struct {
	// required: true
	Firstname string `json:"firstname" yaml:"firstname"`
	Lastname  string `json:"lastname" yaml:"lastname"`
	// required: true
	Email string `json:"email" yaml:"email"`
	// required: true
	Phone string `json:"phone" yaml:"phone"`
}

// UpdateConsumerRequest contains information about the consumer.
// swagger:model
type UpdateConsumerRequest struct {
	Firstname string `json:"firstname" yaml:"firstname"`
	Lastname  string `json:"lastname" yaml:"lastname"`
	Email     string `json:"email" yaml:"email"`
	Phone     string `json:"phone" yaml:"phone"`
}

// NewLocationRequest contains information about the location.
// swagger:model
type NewLocationRequest struct {
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

// UpdateLocationRequest contains information about the location.
// swagger:model
type UpdateLocationRequest struct {
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
