package courierapi

type CourierAccountResponse struct {
	CourierID uint64 `json:"courier_id,omitempty" yaml:"courier_id,omitempty"`
	Balance    int64  `json:"balance" yaml:"balance"`
}

type ReturnAccountingResponseList struct {
	AccountingResponseList []CourierAccountResponse
}
