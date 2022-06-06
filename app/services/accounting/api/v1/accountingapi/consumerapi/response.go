package consumerapi

type ConsumerAccountResponse struct {
	ConsumerID int `json:"consumer_id,omitempty" yaml:"consumer_id,omitempty"`
	Balance    int `json:"balance" yaml:"balance"`
}

type ReturnAccountingResponseList struct {
	AccountingResponseList []ConsumerAccountResponse
}
