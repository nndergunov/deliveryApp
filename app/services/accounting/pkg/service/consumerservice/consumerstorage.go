package consumerservice

import "accounting/pkg/domain"

// ConsumerStorage is the interface for the accounting storage.
type ConsumerStorage interface {
	InsertNewConsumerAccount(account domain.ConsumerAccount) (*domain.ConsumerAccount, error)
	GetConsumerAccountByID(id int) (*domain.ConsumerAccount, error)
	DeleteConsumerAccount(consumerID int) error

	AddToBalanceConsumerAccount(account domain.ConsumerAccount) (*domain.ConsumerAccount, error)
	SubFromBalanceConsumerAccount(account domain.ConsumerAccount) (*domain.ConsumerAccount, error)
}
