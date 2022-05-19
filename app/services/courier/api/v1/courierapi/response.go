package courierapi

type CourierResponse struct {
	ID        uint64 `json:"id" yaml:"id"`
	Username  string `json:"username" yaml:"username"`
	Firstname string `json:"firstname" yaml:"firstname"`
	Lastname  string `json:"lastname,omitempty" yaml:"lastname,omitempty"`
	Email     string `json:"email" yaml:"email"`
	Phone     string `json:"phone,omitempty" yaml:"phone,omitempty"`
	Available bool   `json:"available" yaml:"available"`
}

type ReturnCourierList struct {
	CourierList []CourierResponse
}
