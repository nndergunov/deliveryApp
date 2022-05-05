package service

import (
	"fmt"

	"github.com/nndergunov/deliveryApp/app/services/order/pkg/domain"
)

type Service struct {
	storage Storage
}

func NewService(storage Storage) *Service {
	return &Service{
		storage: storage,
	}
}

func (s Service) CreateOrder(order domain.Order) (*domain.Order, error) {
	order.Status = "Order placed"

	orderID, err := s.storage.InsertOrder(order)
	if err != nil {
		return nil, fmt.Errorf("CreateOrder: %w", err)
	}

	order.OrderID = orderID

	return &order, nil
}

func (s Service) ReturnOrder(orderID int) (*domain.Order, error) {
	order, err := s.storage.GetOrder(orderID)
	if err != nil {
		return nil, fmt.Errorf("ReturnOrder: %w", err)
	}

	return order, nil
}

func (s Service) UpdateOrder(order domain.Order) (*domain.Order, error) {
	err := s.storage.UpdateOrder(order)
	if err != nil {
		return nil, fmt.Errorf("CreateOrder: %w", err)
	}

	return &order, nil
}

func (s Service) ReturnIncompleteOrderList(restaurantID int) ([]domain.Order, error) {
	orders, err := s.storage.GetAllIncompleteOrdersFromRestaurant(restaurantID)
	if err != nil {
		return nil, fmt.Errorf("ReturnIncompleteOrderList: %w", err)
	}

	return orders, nil
}
