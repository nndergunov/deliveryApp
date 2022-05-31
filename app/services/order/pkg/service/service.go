package service

import (
	"fmt"
	"strconv"

	"github.com/nndergunov/deliveryApp/app/pkg/messagebroker/messages"
	"github.com/nndergunov/deliveryApp/app/pkg/messagebroker/publisher"
	"github.com/nndergunov/deliveryApp/app/services/order/pkg/domain"
)

type App interface {
	ReturnOrderList(params domain.SearchParameters) ([]domain.Order, error)
	CreateOrder(order domain.Order) (*domain.Order, error)
	ReturnOrder(orderID int) (*domain.Order, error)
	UpdateOrder(order domain.Order) (*domain.Order, error)
	UpdateStatus(status domain.OrderStatus) (*domain.OrderStatus, error)
}

type Service struct {
	storage       Storage
	notificator   publisher.Publisher
	statusChecker AppStatusChecker
}

func NewService(storage Storage, notificator publisher.Publisher, statusChecker AppStatusChecker) *Service {
	return &Service{
		storage:       storage,
		notificator:   notificator,
		statusChecker: statusChecker,
	}
}

func (s Service) ReturnOrderList(params domain.SearchParameters) ([]domain.Order, error) {
	orders, err := s.storage.GetAllOrders(&params)
	if err != nil {
		return nil, fmt.Errorf("ReturnIncompleteOrderList: %w", err)
	}

	return orders, nil
}

func (s Service) CreateOrder(order domain.Order) (*domain.Order, error) {
	err := s.statusChecker.CheckStatuses()
	if err != nil {
		return nil, fmt.Errorf("checking initial service statuces: %w", err)
	}

	orderID, err := s.storage.InsertOrder(order)
	if err != nil {
		return nil, fmt.Errorf("CreateOrder: %w", err)
	}

	order.OrderID = orderID

	err = s.notificator.Publish("restaurant"+strconv.Itoa(order.RestaurantID), messages.OrderNotification{
		Data: messages.CreatedChange,
	})
	if err != nil {
		return nil, fmt.Errorf("sending notification: %w", err)
	}

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

	err = s.notificator.Publish("restaurant"+strconv.Itoa(order.RestaurantID), messages.OrderNotification{
		Data: messages.UpdatedChange,
	})
	if err != nil {
		return nil, fmt.Errorf("sending notification: %w", err)
	}

	return &order, nil
}

func (s Service) UpdateStatus(status domain.OrderStatus) (*domain.OrderStatus, error) {
	err := s.storage.UpdateOrderStatus(status.OrderID, status.Status)
	if err != nil {
		return nil, fmt.Errorf("CreateOrder: %w", err)
	}

	err = s.notificator.Publish("order"+strconv.Itoa(status.OrderID), messages.OrderNotification{
		Data: messages.StatusUpdatedChange,
	})
	if err != nil {
		return nil, fmt.Errorf("sending notification: %w", err)
	}

	return &status, nil
}
