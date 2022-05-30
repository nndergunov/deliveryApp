package consumerapi

type ConsumerAccountResponse struct {
	ConsumerID uint64 `json:"consumer_id,omitempty" yaml:"consumer_id,omitempty"`
	Balance    int64  `json:"balance" yaml:"balance"`
}

type ReturnAccountingResponseList struct {
	AccountingResponseList []ConsumerAccountResponse
}
