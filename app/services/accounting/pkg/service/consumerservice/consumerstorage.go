package consumerservice

import "accounting/pkg/domain"

// ConsumerStorage is the interface for the accounting storage.
type ConsumerStorage interface {
	InsertNewConsumerAccount(account domain.ConsumerAccount) (*domain.ConsumerAccount, error)
	GetConsumerAccountByID(id uint64) (*domain.ConsumerAccount, error)
	DeleteConsumerAccount(consumerID uint64) error

	AddToBalanceConsumerAccount(account domain.ConsumerAccount) (*domain.ConsumerAccount, error)
	SubFromBalanceConsumerAccount(account domain.ConsumerAccount) (*domain.ConsumerAccount, error)
}
