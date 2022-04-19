package app

import (
	"fmt"

	"github.com/nndergunov/deliveryApp/app/services/kitchen/pkg/domain"
)

// App is a main service logic.
type App struct {
	restaurants []domain.Restaurant
}

func NewApp() *App {
	var rests []domain.Restaurant

	return &App{
		restaurants: rests,
	}
}

func (a *App) ReturnAllRestaurants() []domain.Restaurant {
	return a.restaurants
}

func (a *App) CreateNewRestaurant(rest domain.Restaurant) {
	currID := len(a.restaurants)
	rest.ID = currID

	a.restaurants = append(a.restaurants, rest)
}

func (a *App) UpdateNewRestaurant(rest domain.Restaurant) {
	currID := len(a.restaurants)
	rest.ID = currID

	a.restaurants = append(a.restaurants, rest)
}

func (a *App) UpdateMenu(restaurantID int, newMenu domain.Menu) {
	a.restaurants[restaurantID].Menu = newMenu
}

func (a *App) ReturnMenu(restaurantID int) (*domain.Menu, error) {
	lastArrInd := len(a.restaurants) - 1

	if restaurantID < 0 || restaurantID > lastArrInd {
		return nil, fmt.Errorf("%w: id: %d, possible values: 0 - %d", ErrOutOfRange, restaurantID, lastArrInd)
	}

	return &a.restaurants[restaurantID].Menu, nil
}
