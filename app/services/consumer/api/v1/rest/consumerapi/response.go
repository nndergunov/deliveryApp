package consumerapi

// ConsumerResponse  contains information about the consumer.
// swagger:model
type ConsumerResponse struct {
	ID        int    `json:"id,omitempty" yaml:"id,omitempty"`
	Firstname string `json:"firstname,omitempty" yaml:"firstname,omitempty"`
	Lastname  string `json:"lastname,omitempty" yaml:"lastname,omitempty"`
	Email     string `json:"email,omitempty" yaml:"email,omitempty"`
	Phone     string `json:"phone,omitempty" yaml:"phone,omitempty"`
}

// ReturnConsumerResponseList  contains information about the consumer list.
// swagger:model
type ReturnConsumerResponseList struct {
	ConsumerResponseList []ConsumerResponse
}

// LocationResponse  contains information about the location response.
// swagger:model
type LocationResponse struct {
	UserID     int    `json:"user_id,omitempty" yaml:"user_id,omitempty"`
	Latitude   string `json:"latitude,omitempty" yaml:"latitude,omitempty"`
	Longitude  string `json:"longitude,omitempty" yaml:"longitude,omitempty"`
	Country    string `json:"country,omitempty" yaml:"country,omitempty"`
	City       string `json:"city,omitempty" yaml:"city,omitempty"`
	Region     string `json:"region,omitempty" yaml:"region,omitempty"`
	Street     string `json:"street,omitempty" yaml:"street,omitempty"`
	HomeNumber string `json:"home_number,omitempty" yaml:"home_number,omitempty"`
	Floor      string `json:"floor,omitempty" yaml:"floor,omitempty"`
	Door       string `json:"door,omitempty" yaml:"door,omitempty"`
}
