package handlers

import (
	"github.com/nndergunov/deliveryApp/app/services/restaurant/api/v1/communication"
	"github.com/nndergunov/deliveryApp/app/services/restaurant/pkg/domain"
)

func restaurantToResponse(restaurant domain.Restaurant) communication.ReturnRestaurant {
	return communication.ReturnRestaurant{
		ID:              restaurant.ID,
		Name:            restaurant.Name,
		AcceptingOrders: restaurant.AcceptingOrders,
		City:            restaurant.City,
		Address:         restaurant.Address,
		Longitude:       restaurant.Longitude,
		Latitude:        restaurant.Latitude,
	}
}

func restaurantListToResponse(restaurants []domain.Restaurant) communication.ReturnRestaurantList {
	list := make([]communication.ReturnRestaurant, 0, len(restaurants))

	for _, restaurant := range restaurants {
		currElement := restaurantToResponse(restaurant)

		list = append(list, currElement)
	}

	return communication.ReturnRestaurantList{
		List: list,
	}
}

func menuToResponse(menu domain.Menu) communication.ReturnMenu {
	items := make([]communication.ReturnMenuItem, 0, len(menu.Items))

	for _, menuItem := range menu.Items {
		items = append(items, menuItemToResponse(menuItem))
	}

	return communication.ReturnMenu{
		RestaurantID: menu.RestaurantID,
		MenuItems:    items,
	}
}

func menuItemToResponse(menuItem domain.MenuItem) communication.ReturnMenuItem {
	return communication.ReturnMenuItem{
		ID:     menuItem.ID,
		Name:   menuItem.Name,
		Price:  menuItem.Price,
		Course: menuItem.Course,
	}
}
