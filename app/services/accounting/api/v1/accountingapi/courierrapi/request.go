package courierapi

type NewCourierAccountRequest struct {
	CourierID int `json:"courier_id" yaml:"courierID"`
}

type AddBalanceCourierAccountRequest struct {
	Amount int `json:"amount" yaml:"amount"`
}

type SubBalanceCourierAccountRequest struct {
	Amount int `json:"amount" yaml:"amount"`
}
