package service

import "delivery/pkg/domain"

// DeliveryStorage is the interface for the delivery storage.
type DeliveryStorage interface {
	AssignCourier(courierID int, orderID int) (*domain.AssignedCourier, error)
}
