package service

import "github.com/nndergunov/deliveryApp/app/services/order/pkg/domain"

type Storage interface {
	GetAllOrders() ([]domain.Order, error)
	GetAllIncompleteOrdersFromRestaurant(restaurantID int) ([]domain.Order, error)

	InsertOrder(order domain.Order) (int, error)
	GetOrder(orderID int) (*domain.Order, error)
	UpdateOrder(order domain.Order) error
	DeleteOrder(orderID int) error

	UpdateOrderStatus(orderID int, status string) error
}
