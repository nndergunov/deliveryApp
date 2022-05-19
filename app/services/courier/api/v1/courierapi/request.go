package courierapi

type NewCourierRequest struct {
	Username  string `json:"username" yaml:"username"`
	Password  string `json:"password" yaml:"password"`
	Firstname string `json:"firstname" yaml:"firstname"`
	Lastname  string `json:"lastname" yaml:"lastname"`
	Email     string `json:"email" yaml:"email"`
	Phone     string `json:"phone,omitempty" yaml:"phone,omitempty"`
}

type UpdateCourierRequest struct {
	Username  string `json:"username" yaml:"username"`
	Firstname string `json:"firstname" yaml:"firstname"`
	Lastname  string `json:"lastname" yaml:"lastname"`
	Email     string `json:"email" yaml:"email"`
	Phone     string `json:"phone,omitempty" yaml:"phone,omitempty"`
}
