package courierapi

type NewCourierAccountRequest struct {
	CourierID uint64 `json:"courier_id" yaml:"courierID"`
}

type AddBalanceCourierAccountRequest struct {
	CourierID uint64 `json:"courier_id" yaml:"courierID"`
	Amount    int64  `json:"amount" yaml:"amount"`
}

type SubBalanceCourierAccountRequest struct {
	CourierID uint64 `json:"courier_id" yaml:"courierID"`
	Amount    int64  `json:"amount" yaml:"amount"`
}
