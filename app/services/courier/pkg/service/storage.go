package service

import (
	"courier/pkg/domain"
)

// CourierStorage is the interface for the courier storage.
type CourierStorage interface {
	InsertCourier(courier domain.Courier) (*domain.Courier, error)
	DeleteCourier(id int) error
	UpdateCourier(courier domain.Courier) (*domain.Courier, error)
	UpdateCourierAvailable(id int, available bool) (*domain.Courier, error)
	GetCourierList(param domain.SearchParam) ([]domain.Courier, error)
	GetCourierByID(id int) (*domain.Courier, error)
	GetCourierDuplicateByParam(param domain.SearchParam) (*domain.Courier, error)

	CleanCourierTable() error

	InsertLocation(courier domain.Location) (*domain.Location, error)
	DeleteLocation(courierID int) error
	UpdateLocation(courier domain.Location) (*domain.Location, error)
	GetLocation(userID int) (*domain.Location, error)
	GetLocationList(param domain.SearchParam) ([]domain.Location, error)

	CleanLocationTable() error
}
