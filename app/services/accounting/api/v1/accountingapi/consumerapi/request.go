package consumerapi

type NewConsumerAccountRequest struct {
	ConsumerID uint64 `json:"consumer_id" yaml:"consumerID"`
}

type AddBalanceConsumerAccountRequest struct {
	Amount int64 `json:"amount" yaml:"amount"`
}

type SubBalanceConsumerAccountRequest struct {
	Amount int64 `json:"amount" yaml:"amount"`
}
