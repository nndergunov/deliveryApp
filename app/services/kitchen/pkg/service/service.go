package service

import "github.com/nndergunov/deliveryApp/app/services/kitchen/pkg/domain"

type App interface{}

// Service is a main service logic.
type Service struct {
	communicator Communicator
}

func NewService(communicator Communicator) *Service {
	return &Service{
		communicator: communicator,
	}
}

func (s Service) GetTasks(kitchenID int) domain.Tasks {
	orders, err := s.communicator.GetRestaurantIncompleteOrders(kitchenID)

	// TODO
	panic("TODO")
}
