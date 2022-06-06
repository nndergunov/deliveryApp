package courierservice

import "accounting/pkg/domain"

// CourierStorage is the interface for the accounting storage.
type CourierStorage interface {
	InsertNewCourierAccount(account domain.CourierAccount) (*domain.CourierAccount, error)
	GetCourierAccountByID(id int) (*domain.CourierAccount, error)
	DeleteCourierAccount(courierID int) error

	AddToBalanceCourierAccount(account domain.CourierAccount) (*domain.CourierAccount, error)
	SubFromBalanceCourierAccount(account domain.CourierAccount) (*domain.CourierAccount, error)
}
