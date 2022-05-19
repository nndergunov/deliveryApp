package service

import "courier/domain"

// CourierService is the interface for the user service.
type CourierService interface {
	InsertCourier(courier domain.Courier) (*domain.Courier, error)
	RemoveCourier(id string) (data any, err error)
	UpdateCourier(courier domain.Courier, id string) (*domain.Courier, error)
	UpdateCourierAvailable(id, available string) (*domain.Courier, error)
	GetAllCourier() ([]domain.Courier, error)
	GetCourier(id string) (*domain.Courier, error)
}

// CourierStorage is the interface for the courier storage.
type CourierStorage interface {
	InsertCourier(courier domain.Courier) (*domain.Courier, error)
	RemoveCourier(id uint64) error
	UpdateCourier(courier domain.Courier) (*domain.Courier, error)
	UpdateCourierAvailable(id uint64, available bool) (*domain.Courier, error)
	GetAllCourier() ([]domain.Courier, error)
	GetCourier(id uint64, username, status string) (*domain.Courier, error)
	CleanDB() error
}
