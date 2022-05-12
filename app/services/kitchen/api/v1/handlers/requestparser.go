package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/nndergunov/deliveryApp/app/pkg/api/v1/restaurantapi"
	"github.com/nndergunov/deliveryApp/app/services/restaurant/pkg/domain"
)

func getIDFromEndpoint(code string, r *http.Request) (int, error) {
	vars := mux.Vars(r)

	idVar := vars[code]
	if idVar == "" {
		return 0, fmt.Errorf("getIDFromEndpoint: %w", errNoIDInEndpoint)
	}

	id, err := strconv.Atoi(idVar)
	if err != nil {
		return 0, fmt.Errorf("getIDFromEndpoint: %w", err)
	}

	return id, nil
}

func requestToRestaurant(restID int, req *restaurantapi.RestaurantData) domain.Restaurant {
	return domain.Restaurant{
		ID:      restID,
		Name:    req.Name,
		City:    req.City,
		Address: req.Address,
	}
}

func requestToMenu(restaurantID int, req *restaurantapi.MenuData) domain.Menu {
	menuItems := make([]domain.MenuItem, 0, len(req.MenuItems))

	for _, item := range req.MenuItems {
		currItem := domain.MenuItem{
			ID:     item.ID,
			MenuID: 0,
			Name:   item.Name,
			Course: item.Course,
		}

		menuItems = append(menuItems, currItem)
	}

	return domain.Menu{
		RestaurantID: restaurantID,
		Items:        menuItems,
	}
}

func requestToMenuItem(itemID int, req *restaurantapi.MenuItemData) domain.MenuItem {
	return domain.MenuItem{
		ID:     itemID,
		MenuID: 0,
		Name:   req.Name,
		Course: req.Course,
	}
}
