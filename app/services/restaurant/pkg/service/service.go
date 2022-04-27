package service

import (
	"fmt"

	"github.com/nndergunov/deliveryApp/app/services/restaurant/pkg/domain"
)

// Service is a main service logic.
type Service struct {
	storage Storage
}

func NewService(storage Storage) *Service {
	return &Service{
		storage: storage,
	}
}

func (s Service) ReturnAllRestaurants() ([]domain.Restaurant, error) {
	restaurants, err := s.storage.GetAllRestaurants()
	if err != nil {
		return nil, fmt.Errorf("ReturnAllRestaurants: %w", err)
	}

	return restaurants, nil
}

func (s Service) CreateNewRestaurant(restaurant domain.Restaurant) (*domain.Restaurant, error) {
	id, err := s.storage.InsertRestaurant(restaurant)
	if err != nil {
		return nil, fmt.Errorf("InsertRestaurant: %w", err)
	}

	restaurant.ID = id

	return &restaurant, nil
}

func (s Service) UpdateRestaurant(restaurant domain.Restaurant) (*domain.Restaurant, error) {
	err := s.storage.UpdateRestaurant(restaurant)
	if err != nil {
		return nil, fmt.Errorf("UpdateRestaurant: %w", err)
	}

	return &restaurant, nil
}

func (s Service) DeleteRestaurant(restaurantID int) error {
	err := s.storage.DeleteRestaurant(restaurantID)
	if err != nil {
		return fmt.Errorf("DeleteRestaurant: %w", err)
	}

	return nil
}

func (s Service) ReturnMenu(restaurantID int) (*domain.Menu, error) {
	menu, err := s.storage.GetMenu(restaurantID)
	if err != nil {
		return nil, fmt.Errorf("ReturnMenu: %w", err)
	}

	return menu, nil
}

func (s Service) CreateMenu(menu domain.Menu) (*domain.Menu, error) {
	_, menuItems, err := s.storage.InsertMenu(menu)
	if err != nil {
		return nil, fmt.Errorf("InsertMenu: %w", err)
	}

	menu.Items = menuItems

	return &menu, nil
}

func (s Service) AddMenuItem(restaurantID int, menuItem domain.MenuItem) (*domain.MenuItem, error) {
	id, err := s.storage.InsertMenuItem(restaurantID, menuItem)
	if err != nil {
		return nil, fmt.Errorf("InsertMenuItem: %w", err)
	}

	menuItem.ID = id

	return &menuItem, nil
}

func (s Service) UpdateMenuItem(menuItem domain.MenuItem) (*domain.MenuItem, error) {
	err := s.storage.UpdateMenuItem(menuItem)
	if err != nil {
		return nil, fmt.Errorf("UpdateMenuItem: %w", err)
	}

	return &menuItem, nil
}

func (s Service) DeleteMenuItem(menuItemID int) error {
	err := s.storage.DeleteMenuItem(menuItemID)
	if err != nil {
		return fmt.Errorf("DeleteMenuItem: %w", err)
	}

	return nil
}
