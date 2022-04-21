package app

import (
	"fmt"

	"github.com/nndergunov/deliveryApp/app/services/restaurant/pkg/db"
	"github.com/nndergunov/deliveryApp/app/services/restaurant/pkg/domain"
)

// App is a main service logic.
type App struct {
	database *db.Database
}

func NewApp(database *db.Database) *App {
	return &App{
		database: database,
	}
}

func (a App) ReturnAllRestaurants() ([]domain.Restaurant, error) {
	restaurants, err := a.database.ReturnAllRestaurants()
	if err != nil {
		return nil, fmt.Errorf("ReturnAllRestaurants: %w", err)
	}

	return restaurants, nil
}

func (a App) CreateNewRestaurant(restaurant domain.Restaurant) error {
	err := a.database.CreateNewRestaurant(restaurant)
	if err != nil {
		return fmt.Errorf("CreateNewRestaurant: %w", err)
	}

	return nil
}

func (a App) UpdateRestaurant(restaurant domain.Restaurant) error {
	err := a.database.UpdateRestaurant(restaurant)
	if err != nil {
		return fmt.Errorf("UpdateRestaurant: %w", err)
	}

	return nil
}

func (a App) ReturnMenu(restaurantID int) (*domain.Menu, error) {
	menu, err := a.database.ReturnMenu(restaurantID)
	if err != nil {
		return nil, fmt.Errorf("ReturnMenu: %w", err)
	}

	return menu, nil
}

func (a App) CreateMenu(menu domain.Menu) error {
	err := a.database.CreateMenu(menu)
	if err != nil {
		return fmt.Errorf("CreateMenu: %w", err)
	}

	return nil
}

func (a App) AddMenuItem(restaurantID int, menuItem domain.MenuItem) error {
	err := a.database.AddMenuItem(restaurantID, menuItem)
	if err != nil {
		return fmt.Errorf("AddMenuItem: %w", err)
	}

	return nil
}

func (a App) UpdateMenuItem(menuItem domain.MenuItem) error {
	err := a.database.UpdateMenuItem(menuItem)
	if err != nil {
		return fmt.Errorf("UpdateMenuItem: %w", err)
	}

	return nil
}

func (a App) DeleteMenuItem(menuItemID int) error {
	err := a.database.DeleteMenuItem(menuItemID)
	if err != nil {
		return fmt.Errorf("DeleteMenuItem: %w", err)
	}

	return nil
}
