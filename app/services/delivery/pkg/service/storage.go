package service

import "github.com/nndergunov/deliveryApp/app/services/delivery/pkg/domain"

// DeliveryStorage is the interface for the delivery storage.
type DeliveryStorage interface {
	AssignOrder(courierID int, orderID int) (*domain.AssignOrder, error)
}
