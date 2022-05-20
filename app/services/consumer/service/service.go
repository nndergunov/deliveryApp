package service

import "consumer/domain"

// ConsumerService is the interface for the user service.
type ConsumerService interface {
	InsertConsumer(consumer domain.Consumer) (*domain.Consumer, error)
	DeleteConsumer(id string) (data any, err error)
	UpdateConsumer(consumer domain.Consumer, id string) (*domain.Consumer, error)
	GetAllConsumer() ([]domain.Consumer, error)
	GetConsumer(id string) (*domain.Consumer, error)

	UpdateConsumerLocation(consumer domain.ConsumerLocation, id string) (*domain.ConsumerLocation, error)
}

// ConsumerStorage is the interface for the consumer storage.
type ConsumerStorage interface {
	InsertConsumer(consumer domain.Consumer) (*domain.Consumer, error)
	DeleteConsumer(id uint64) error
	UpdateConsumer(consumer domain.Consumer) (*domain.Consumer, error)
	GetAllConsumer() ([]domain.Consumer, error)
	GetConsumer(id uint64, phone, email string) (*domain.Consumer, error)
	CleanConsumerTable() error

	UpdateConsumerLocation(consumer domain.ConsumerLocation) (*domain.ConsumerLocation, error)
	CleanConsumerLocationTable() error
}
