package courierapi

type NewCourierAccountRequest struct {
	CourierID uint64 `json:"courier_id" yaml:"courierID"`
}

type AddBalanceCourierAccountRequest struct {
	Amount int64 `json:"amount" yaml:"amount"`
}

type SubBalanceCourierAccountRequest struct {
	Amount int64 `json:"amount" yaml:"amount"`
}
