package courierapi

// NewCourierRequest contains information about the courier.
// swagger:model
type NewCourierRequest struct {
	// required: true
	Username string `json:"username,omitempty" yaml:"username,omitempty"`
	// required: true
	Password string `json:"password,omitempty" yaml:"password,omitempty"`
	// required: true
	Firstname string `json:"firstname,omitempty" yaml:"firstname,omitempty"`
	Lastname  string `json:"lastname,omitempty" yaml:"lastname,omitempty"`
	// required: true
	Email string `json:"email,omitempty" yaml:"email,omitempty"`
	// required: true
	Phone string `json:"phone,omitempty" yaml:"phone,omitempty"`
}

// UpdateCourierRequest contains information about the courier.
// swagger:model
type UpdateCourierRequest struct {
	Username  string `json:"username,omitempty" yaml:"username,omitempty"`
	Firstname string `json:"firstname,omitempty" yaml:"firstname,omitempty"`
	Lastname  string `json:"lastname,omitempty" yaml:"lastname,omitempty"`
	Email     string `json:"email,omitempty" yaml:"email,omitempty"`
	Phone     string `json:"phone,omitempty" yaml:"phone,omitempty"`
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
