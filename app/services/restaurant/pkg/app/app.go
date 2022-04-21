package app

import (
	"fmt"

	"github.com/nndergunov/deliveryApp/app/services/restaurant/pkg/domain"
)

// App is a main service logic.
type App struct {
	restaurants map[int]*domain.Restaurant
	menus       map[int]*domain.Menu
}

func NewApp() *App {
	rests := make(map[int]*domain.Restaurant)
	menus := make(map[int]*domain.Menu)

	return &App{
		restaurants: rests,
		menus:       menus,
	}
}

func (a *App) ReturnAllRestaurants() []domain.Restaurant {
	restaurants := make([]domain.Restaurant, 0, len(a.restaurants))

	for _, restaurant := range a.restaurants {
		restaurants = append(restaurants, *restaurant)
	}

	return restaurants
}

func (a *App) CreateNewRestaurant(restaurant domain.Restaurant) error {
	currID := 123 // To be changed to database id.

	a.restaurants[currID] = &restaurant

	return nil
}

func (a *App) UpdateRestaurant(restaurant domain.Restaurant) error {
	if _, ok := a.restaurants[restaurant.ID]; !ok {
		return fmt.Errorf("%w: id: %d", ErrIsNotInMap, restaurant.ID)
	}

	a.restaurants[restaurant.ID] = &restaurant

	return nil
}

func (a *App) ReturnMenu(restaurantID int) (*domain.Menu, error) {
	if _, ok := a.restaurants[restaurantID]; !ok {
		return nil, fmt.Errorf("%w: id: %d", ErrIsNotInMap, restaurantID)
	}

	return a.menus[restaurantID], nil
}

func (a *App) CreateMenu(menu domain.Menu) error {
	if _, ok := a.restaurants[menu.RestaurantID]; !ok {
		return fmt.Errorf("%w: id: %d", ErrIsNotInMap, menu.RestaurantID)
	}

	a.menus[menu.RestaurantID] = &menu

	return nil
}

func (a *App) AddMenuItem(restaurantID int, menuItem domain.MenuItem) error {
	if _, ok := a.restaurants[restaurantID]; !ok {
		return fmt.Errorf("%w: id: %d", ErrIsNotInMap, restaurantID)
	}

	next := 123 // To be replaced by the database index.

	a.menus[restaurantID].Items[next] = menuItem

	return nil
}

func (a *App) UpdateMenuItem(restaurantID int, updatedMenuItem domain.MenuItem) error {
	if _, ok := a.restaurants[restaurantID]; !ok {
		return fmt.Errorf("%w: id: %d", ErrIsNotInMap, restaurantID)
	}

	if _, ok := a.menus[restaurantID].Items[updatedMenuItem.ID]; !ok {
		return fmt.Errorf("%w: id: %d", ErrIsNotInMap, updatedMenuItem.ID)
	}

	a.menus[restaurantID].Items[updatedMenuItem.ID] = updatedMenuItem

	return nil
}

func (a *App) DeleteMenuItem(restaurantID, menuItemID int) error {
	if _, ok := a.restaurants[restaurantID]; !ok {
		return fmt.Errorf("%w: id: %d", ErrIsNotInMap, restaurantID)
	}

	if _, ok := a.menus[restaurantID].Items[menuItemID]; !ok {
		return fmt.Errorf("%w: id: %d", ErrIsNotInMap, menuItemID)
	}

	delete(a.menus[restaurantID].Items, menuItemID)

	return nil
}
