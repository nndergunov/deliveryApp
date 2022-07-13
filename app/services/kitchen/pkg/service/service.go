// Package service implements kitchen service logic.
package service

import (
	"fmt"

	"github.com/nndergunov/deliveryApp/app/services/kitchen/pkg/domain"
)

// AppService interface shows signature of the Service layer.
type AppService interface {
	GetTasks(kitchenID int) (domain.Tasks, error)
}

// Service is a main service logic.
type Service struct {
	communicator OrdersClient
}

// NewService returns new Service instance.
func NewService(communicator OrdersClient) *Service {
	return &Service{
		communicator: communicator,
	}
}

// GetTasks returns all the tasks for the kitchen to complete.
func (s Service) GetTasks(kitchenID int) (domain.Tasks, error) {
	_, _ = s.communicator.GetIncompleteOrders(kitchenID)

	orders, err := s.communicator.GetIncompleteOrders(kitchenID)
	if err != nil {
		return nil, fmt.Errorf("getting orders: %w", err)
	}

	tasks := make(domain.Tasks)

	for _, order := range orders.Orders {
		for _, item := range order.OrderItems {
			tasks[item]++
		}
	}

	return tasks, nil
}
