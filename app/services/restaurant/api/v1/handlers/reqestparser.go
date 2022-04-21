package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/nndergunov/deliveryApp/app/pkg/api/v1/restaurantapi"
	"github.com/nndergunov/deliveryApp/app/services/kitchen/pkg/domain"
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
		Address: req.Street,
	}
}

func requestToMenu(restaurantID int, req *restaurantapi.MenuData) domain.Menu {
	menuItems := make(map[int]domain.MenuItem)

	for _, item := range req.MenuItems {
		menuItems[item.ID] = domain.MenuItem{
			ID:     item.ID,
			Name:   item.Name,
			Course: item.Course,
		}
	}

	return domain.Menu{
		RestaurantID: restaurantID,
		Items:        menuItems,
	}
}

func requestToMenuItem(itemID int, req *restaurantapi.MenuItemData) domain.MenuItem {
	return domain.MenuItem{
		ID:     itemID,
		Name:   req.Name,
		Course: req.Course,
	}
}
