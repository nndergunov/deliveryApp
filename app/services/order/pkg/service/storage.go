package service

import "github.com/nndergunov/deliveryApp/app/services/order/pkg/domain"

// Storage interface outlines storage repository layer.
type Storage interface {
	GetAllOrders(params *domain.SearchParameters) ([]domain.Order, error)

	InsertOrder(order domain.Order) (int, error)
	GetOrder(orderID int) (*domain.Order, error)
	UpdateOrder(order domain.Order) error
	DeleteOrder(orderID int) error

	UpdateOrderStatus(orderID int, status string) error
}
