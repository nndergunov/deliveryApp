package courierapi

type CourierAccountResponse struct {
	CourierID int `json:"courier_id,omitempty" yaml:"courier_id,omitempty"`
	Balance   int `json:"balance" yaml:"balance"`
}

type ReturnAccountingResponseList struct {
	AccountingResponseList []CourierAccountResponse
}
