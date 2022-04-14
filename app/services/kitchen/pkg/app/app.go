package app

import (
	"fmt"

	domain2 "github.com/nndergunov/deliveryApp/app/services/kitchen/pkg/domain"
)

// App is a simple struct to store restaurants and menus,
// will be replaced with something with more functional down the line of development.
type App struct {
	restaurants []domain2.Restaurant
}

func NewApp() *App {
	var rests []domain2.Restaurant

	return &App{
		restaurants: rests,
	}
}

func (a *App) CreateNewRestaurant(rest domain2.Restaurant) {
	currID := len(a.restaurants)
	rest.Id = currID

	a.restaurants = append(a.restaurants, rest)
}

func (a *App) ReturnAllRestaurants() []domain2.Restaurant {
	return a.restaurants
}

func (a *App) UpdateMenu(restaurantID int, newMenu domain2.Menu) {
	a.restaurants[restaurantID].Menu = newMenu
}

func (a *App) ReturnMenu(restaurantID int) (*domain2.Menu, error) {
	lastArrInd := len(a.restaurants) - 1

	if restaurantID < 0 || restaurantID > lastArrInd {
		return nil, fmt.Errorf("%w: id: %d, possible values: 0 - %d", ErrOutOfRange, restaurantID, lastArrInd)
	}

	return &a.restaurants[restaurantID].Menu, nil
}
