package service

import "github.com/nndergunov/deliveryApp/app/services/restaurant/pkg/domain"

type Storage interface {
	GetAllRestaurants() ([]domain.Restaurant, error)

	InsertRestaurant(restaurant domain.Restaurant) (int, error)
	GetRestaurant(restaurantID int) (*domain.Restaurant, error)
	UpdateRestaurant(restaurant domain.Restaurant) error
	DeleteRestaurant(restaurantID int) error

	InsertMenu(menu domain.Menu) (int, []domain.MenuItem, error)
	GetMenu(restaurantID int) (*domain.Menu, error)
	UpdateMenu(menu domain.Menu) error
	DeleteMenu(restaurantID int) error

	InsertMenuItem(restaurantID int, menuItem domain.MenuItem) (int, error)
	GetMenuItem(menuItemID int) (*domain.MenuItem, error)
	UpdateMenuItem(menuItem domain.MenuItem) error
	DeleteMenuItem(menuItemID int) error
}
