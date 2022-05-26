package consumerapi

type NewConsumerRequest struct {
	Firstname string `json:"firstname" yaml:"firstname"`
	Lastname  string `json:"lastname" yaml:"lastname"`
	Email     string `json:"email" yaml:"email"`
	Phone     string `json:"phone" yaml:"phone"`
}

type UpdateConsumerRequest struct {
	Firstname string `json:"firstname" yaml:"firstname"`
	Lastname  string `json:"lastname" yaml:"lastname"`
	Email     string `json:"email" yaml:"email"`
	Phone     string `json:"phone" yaml:"phone"`
}
type NewConsumerLocationRequest struct {
	Altitude   string `json:"altitude" yaml:"altitude"`
	Longitude  string `json:"longitude" yaml:"longitude"`
	Country    string `json:"country" yaml:"country"`
	City       string `json:"city" yaml:"city"`
	Region     string `json:"region" yaml:"region"`
	Street     string `json:"street" yaml:"street"`
	HomeNumber string `json:"home_number" yaml:"home_number"`
	Floor      string `json:"floor" yaml:"floor"`
	Door       string `json:"door" yaml:"door"`
}
type UpdateConsumerLocationRequest struct {
	Altitude   string `json:"altitude" yaml:"altitude"`
	Longitude  string `json:"longitude" yaml:"longitude"`
	Country    string `json:"country" yaml:"country"`
	City       string `json:"city" yaml:"city"`
	Region     string `json:"region" yaml:"region"`
	Street     string `json:"street" yaml:"street"`
	HomeNumber string `json:"home_number" yaml:"home_number"`
	Floor      string `json:"floor" yaml:"floor"`
	Door       string `json:"door" yaml:"door"`
}
