package service

import "accounting/pkg/domain"

// AccountingStorage is the interface for the courier storage.
type AccountingStorage interface {
	InsertNewConsumerAccount(account domain.ConsumerAccount) (*domain.ConsumerAccount, error)
	GetConsumerAccountByID(id uint64) (*domain.ConsumerAccount, error)
	DeleteConsumerAccount(consumerID uint64) error

	AddToConsumerAccount(account domain.ConsumerAccount) (*domain.ConsumerAccount, error)
	SubToConsumerAccount(account domain.ConsumerAccount) (*domain.ConsumerAccount, error)
}
