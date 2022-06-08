package v1

// Status is a struct that contains Service Name and whether it is up.
type Status struct {
	ServiceName string `json:"service"`
	IsUp        string `json:"isUp"`
}

type ServiceError struct {
	HTTPStatus int
	ErrorText  string
}
