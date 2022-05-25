package consumerapi

type ConsumerResponse struct {
	ID        uint64 `json:"id,omitempty" yaml:"id,omitempty"`
	Firstname string `json:"firstname,omitempty" yaml:"firstname,omitempty"`
	Lastname  string `json:"lastname,omitempty" yaml:"lastname,omitempty"`
	Email     string `json:"email,omitempty" yaml:"email,omitempty"`
	Phone     string `json:"phone,omitempty" yaml:"phone,omitempty"`
}

type ReturnConsumerResponseList struct {
	ConsumerResponseList []ConsumerResponse
}

type ConsumerLocationResponse struct {
	ConsumerID uint64 `json:"consumer_id,omitempty" yaml:"consumer_id,omitempty"`
	Altitude   string `json:"altitude,omitempty" yaml:"altitude,omitempty"`
	Longitude  string `json:"Longitude,omitempty" yaml:"Longitude,omitempty"`
	Country    string `json:"country,omitempty" yaml:"country,omitempty"`
	City       string `json:"city,omitempty" yaml:"city,omitempty"`
	Region     string `json:"region,omitempty" yaml:"region,omitempty"`
	Street     string `json:"street,omitempty" yaml:"street,omitempty"`
	HomeNumber string `json:"home_number,omitempty" yaml:"home_number,omitempty"`
	Floor      string `json:"floor,omitempty" yaml:"floor,omitempty"`
	Door       string `json:"door,omitempty" yaml:"door,omitempty"`
}
