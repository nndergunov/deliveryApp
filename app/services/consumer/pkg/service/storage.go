package service

import (
	"consumer/pkg/domain"
)

// ConsumerStorage is the interface for the consumer storage.
type ConsumerStorage interface {
	InsertConsumer(consumer domain.Consumer) (*domain.Consumer, error)
	DeleteConsumer(id uint64) error
	UpdateConsumer(consumer domain.Consumer) (*domain.Consumer, error)
	GetAllConsumer() ([]domain.Consumer, error)
	GetConsumer(id uint64, phone, email string) (*domain.Consumer, error)
	CleanConsumerTable() error

	InsertConsumerLocation(consumer domain.ConsumerLocation) (*domain.ConsumerLocation, error)
	DeleteConsumerLocation(consumerID uint64) error
	UpdateConsumerLocation(consumer domain.ConsumerLocation) (*domain.ConsumerLocation, error)
	GetConsumerLocation(id uint64) (*domain.ConsumerLocation, error)
}
