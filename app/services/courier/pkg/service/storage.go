package service

import (
	"courier/pkg/domain"
)

// CourierStorage is the interface for the courier storage.
type CourierStorage interface {
	InsertCourier(courier domain.Courier) (*domain.Courier, error)
	DeleteCourier(id uint64) error
	UpdateCourier(courier domain.Courier) (*domain.Courier, error)
	UpdateCourierAvailable(id uint64, available bool) (*domain.Courier, error)
	GetAllCourier(param domain.SearchParam) ([]domain.Courier, error)
	GetCourierByID(id uint64) (*domain.Courier, error)
	GetCourierDuplicateByParam(param domain.SearchParam) (*domain.Courier, error)

	CleanCourierTable() error

	InsertCourierLocation(courier domain.CourierLocation) (*domain.CourierLocation, error)
	DeleteCourierLocation(courierID uint64) error
	UpdateCourierLocation(courier domain.CourierLocation) (*domain.CourierLocation, error)
	GetCourierLocation(id uint64) (*domain.CourierLocation, error)
}
