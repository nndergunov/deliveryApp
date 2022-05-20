package consumerapi

type ConsumerResponse struct {
	ID               uint64 `json:"id" yaml:"id"`
	Firstname        string `json:"firstname" yaml:"firstname"`
	Lastname         string `json:"lastname,omitempty" yaml:"lastname,omitempty"`
	Email            string `json:"email,omitempty" yaml:"email,omitempty"`
	Phone            string `json:"phone,omitempty" yaml:"phone,omitempty"`
	ConsumerLocation ConsumerLocationResponse
}

type ReturnConsumerList struct {
	ConsumerList []ConsumerResponse
}

type ConsumerLocationResponse struct {
	LocationAlt string `json:"location_alt,omitempty" yaml:"location_alt,omitempty"`
	LocationLat string `json:"location_lat,omitempty" yaml:"location_lat,omitempty"`
	Country     string `json:"country,omitempty" yaml:"country,omitempty"`
	City        string `json:"city,omitempty" yaml:"city,omitempty"`
	Region      string `json:"region,omitempty" yaml:"region,omitempty"`
	Street      string `json:"street,omitempty" yaml:"street,omitempty"`
	HomeNumber  string `json:"home_number,omitempty" yaml:"home_number,omitempty"`
	Floor       string `json:"floor,omitempty" yaml:"floor,omitempty"`
	Door        string `json:"door,omitempty" yaml:"door,omitempty"`
}
