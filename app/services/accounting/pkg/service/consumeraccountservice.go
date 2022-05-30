package service

import "accounting/pkg/domain"

// AccountingService is the interface for the user service.
type AccountingService interface {
	InsertNewConsumerAccount(account domain.ConsumerAccount) (*domain.ConsumerAccount, error)
	GetConsumerAccount(consumerID string) (*domain.ConsumerAccount, error)
	DeleteConsumerAccount(consumerID string) (string, error)

	AddToConsumerAccount(account domain.ConsumerAccount) (*domain.ConsumerAccount, error)
	SubFromConsumerAccount(account domain.ConsumerAccount) (*domain.ConsumerAccount, error)
}
