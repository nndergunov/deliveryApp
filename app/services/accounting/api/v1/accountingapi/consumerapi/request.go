package consumerapi

type NewConsumerAccountRequest struct {
	ConsumerID int `json:"consumer_id" yaml:"consumerID"`
}

type AddBalanceConsumerAccountRequest struct {
	Amount int `json:"amount" yaml:"amount"`
}

type SubBalanceConsumerAccountRequest struct {
	Amount int `json:"amount" yaml:"amount"`
}
