package consumerapi

type NewConsumerAccountRequest struct {
	ConsumerID uint64 `json:"consumer_id" yaml:"consumerID"`
}

type AddConsumerAccountRequest struct {
	ConsumerID uint64 `json:"consumer_id" yaml:"consumerID"`
	Amount     int64  `json:"amount" yaml:"amount"`
}

type SubConsumerAccountRequest struct {
	ConsumerID uint64 `json:"consumer_id" yaml:"consumerID"`
	Amount     int64  `json:"amount" yaml:"amount"`
}
