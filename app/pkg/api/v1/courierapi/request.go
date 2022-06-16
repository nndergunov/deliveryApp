package courierapi

type NewCourierRequest struct {
	Username  string `json:"username,omitempty" yaml:"username,omitempty"`
	Password  string `json:"password,omitempty" yaml:"password,omitempty"`
	Firstname string `json:"firstname,omitempty" yaml:"firstname,omitempty"`
	Lastname  string `json:"lastname,omitempty" yaml:"lastname,omitempty"`
	Email     string `json:"email,omitempty" yaml:"email,omitempty"`
	Phone     string `json:"phone,omitempty" yaml:"phone,omitempty"`
}

type UpdateCourierRequest struct {
	Username  string `json:"username,omitempty" yaml:"username,omitempty"`
	Firstname string `json:"firstname,omitempty" yaml:"firstname,omitempty"`
	Lastname  string `json:"lastname,omitempty" yaml:"lastname,omitempty"`
	Email     string `json:"email,omitempty" yaml:"email,omitempty"`
	Phone     string `json:"phone,omitempty" yaml:"phone,omitempty"`
}
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
