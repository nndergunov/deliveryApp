package restaurantservice

import "accounting/pkg/domain"

// RestaurantStorage is the interface for the accounting storage.
type RestaurantStorage interface {
	InsertNewRestaurantAccount(account domain.RestaurantAccount) (*domain.RestaurantAccount, error)
	GetRestaurantAccountByID(id uint64) (*domain.RestaurantAccount, error)
	DeleteRestaurantAccount(restaurantID uint64) error

	AddToBalanceRestaurantAccount(account domain.RestaurantAccount) (*domain.RestaurantAccount, error)
	SubFromBalanceRestaurantAccount(account domain.RestaurantAccount) (*domain.RestaurantAccount, error)
}
