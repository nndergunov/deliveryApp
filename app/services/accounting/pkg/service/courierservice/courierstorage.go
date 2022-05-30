package courierservice

import "accounting/pkg/domain"

// CourierStorage is the interface for the accounting storage.
type CourierStorage interface {
	InsertNewCourierAccount(account domain.CourierAccount) (*domain.CourierAccount, error)
	GetCourierAccountByID(id uint64) (*domain.CourierAccount, error)
	DeleteCourierAccount(courierID uint64) error

	AddToBalanceCourierAccount(account domain.CourierAccount) (*domain.CourierAccount, error)
	SubFromBalanceCourierAccount(account domain.CourierAccount) (*domain.CourierAccount, error)
}
