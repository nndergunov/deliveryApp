package v1

type Status struct {
	Service string `json:"service"`
	IsUp    string `json:"isUp"`
}
