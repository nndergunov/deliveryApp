package restaurantservice

import "accounting/pkg/domain"

// RestaurantStorage is the interface for the accounting storage.
type RestaurantStorage interface {
	InsertNewRestaurantAccount(account domain.RestaurantAccount) (*domain.RestaurantAccount, error)
	GetRestaurantAccountByID(id int) (*domain.RestaurantAccount, error)
	DeleteRestaurantAccount(restaurantID int) error

	AddToBalanceRestaurantAccount(account domain.RestaurantAccount) (*domain.RestaurantAccount, error)
	SubFromBalanceRestaurantAccount(account domain.RestaurantAccount) (*domain.RestaurantAccount, error)
}
