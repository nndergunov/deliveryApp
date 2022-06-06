package service

import "delivery/pkg/domain"

// DeliveryStorage is the interface for the delivery storage.
type DeliveryStorage interface {
	AssignOrder(courierID int, orderID int) (*domain.AssignOrder, error)
}
