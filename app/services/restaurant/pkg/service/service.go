package service

import (
	"fmt"

	"github.com/nndergunov/deliveryApp/app/services/restaurant/pkg/domain"
)

type AppService interface {
	ReturnAllRestaurants() ([]domain.Restaurant, error)
	CreateNewRestaurant(restaurant domain.Restaurant) (*domain.Restaurant, error)
	UpdateRestaurant(restaurant domain.Restaurant) (*domain.Restaurant, error)
	DeleteRestaurant(restaurantID int) error
	ReturnMenu(restaurantID int) (*domain.Menu, error)
	CreateMenu(menu domain.Menu) (*domain.Menu, error)
	AddMenuItem(restaurantID int, menuItem domain.MenuItem) (*domain.MenuItem, error)
	UpdateMenuItem(restaurantID int, menuItem domain.MenuItem) (*domain.MenuItem, error)
	DeleteMenuItem(restaurantID int, menuItemID int) error
}

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

func (s Service) UpdateMenuItem(restaurantID int, menuItem domain.MenuItem) (*domain.MenuItem, error) {
	isInMenu, err := s.checkIfItemIsInMenu(restaurantID, menuItem.ID)
	if err != nil {
		return nil, fmt.Errorf("UpdateMenuItem: %w", err)
	}

	if !isInMenu {
		return nil, fmt.Errorf("UpdateMenuItem: %w, restaurantID: %d, menuItemID: %d",
			ErrItemIsNotInMenu, restaurantID, menuItem.ID)
	}

	err = s.storage.UpdateMenuItem(menuItem)
	if err != nil {
		return nil, fmt.Errorf("UpdateMenuItem: %w", err)
	}

	return &menuItem, nil
}

func (s Service) DeleteMenuItem(restaurantID, menuItemID int) error {
	isInMenu, err := s.checkIfItemIsInMenu(restaurantID, menuItemID)
	if err != nil {
		return fmt.Errorf("DeleteMenuItem: %w", err)
	}

	if !isInMenu {
		return fmt.Errorf("DeleteMenuItem: %w, restaurantID: %d, menuItemID: %d",
			ErrItemIsNotInMenu, restaurantID, menuItemID)
	}

	err = s.storage.DeleteMenuItem(menuItemID)
	if err != nil {
		return fmt.Errorf("DeleteMenuItem: %w", err)
	}

	return nil
}

func (s Service) checkIfItemIsInMenu(restaurantID, menuItemID int) (bool, error) {
	menu, err := s.storage.GetMenu(restaurantID)
	if err != nil {
		return false, fmt.Errorf("DeleteMenuItem: %w", err)
	}

	for _, menuItem := range menu.Items {
		if menuItem.ID == menuItemID {
			return true, nil
		}
	}

	return false, nil
}
