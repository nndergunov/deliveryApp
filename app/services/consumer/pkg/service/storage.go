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

	InsertLocation(location domain.Location) (*domain.Location, error)
	DeleteLocation(userID int) error
	UpdateLocation(location domain.Location) (*domain.Location, error)
	GetLocation(userID int) (*domain.Location, error)
}
