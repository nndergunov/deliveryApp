package service

import (
	"consumer/pkg/domain"
)

// ConsumerStorage is the interface for the consumer storage.
type ConsumerStorage interface {
	InsertConsumer(consumer domain.Consumer) (*domain.Consumer, error)
	DeleteConsumer(id int) error
	UpdateConsumer(consumer domain.Consumer) (*domain.Consumer, error)
	GetAllConsumer() ([]domain.Consumer, error)
	GetConsumerByID(id int) (*domain.Consumer, error)
	GetConsumerDuplicateByParam(param domain.SearchParam) (*domain.Consumer, error)
	CleanConsumerTable() error

	InsertConsumerLocation(consumer domain.ConsumerLocation) (*domain.ConsumerLocation, error)
	DeleteConsumerLocation(consumerID int) error
	UpdateConsumerLocation(consumer domain.ConsumerLocation) (*domain.ConsumerLocation, error)
	GetConsumerLocation(id int) (*domain.ConsumerLocation, error)
}
