package handlers

import (
	"github.com/nndergunov/deliveryApp/app/pkg/api/v1/restaurantapi"
	"github.com/nndergunov/deliveryApp/app/services/restaurant/pkg/domain"
)

func restaurantToResponse(restaurant domain.Restaurant) restaurantapi.ReturnRestaurant {
	return restaurantapi.ReturnRestaurant{
		ID:              restaurant.ID,
		Name:            restaurant.Name,
		AcceptingOrders: restaurant.AcceptingOrders,
		City:            restaurant.City,
		Address:         restaurant.Address,
		Longitude:       restaurant.Longitude,
		Altitude:        restaurant.Altitude,
	}
}

func restaurantListToResponse(restaurants []domain.Restaurant) restaurantapi.ReturnRestaurantList {
	list := make([]restaurantapi.ReturnRestaurant, 0, len(restaurants))

	for _, restaurant := range restaurants {
		currElement := restaurantToResponse(restaurant)

		list = append(list, currElement)
	}

	return restaurantapi.ReturnRestaurantList{
		List: list,
	}
}

func menuToResponse(menu domain.Menu) restaurantapi.ReturnMenu {
	items := make([]restaurantapi.ReturnMenuItem, 0, len(menu.Items))

	for _, menuItem := range menu.Items {
		items = append(items, menuItemToResponse(menuItem))
	}

	return restaurantapi.ReturnMenu{
		RestaurantID: menu.RestaurantID,
		MenuItems:    items,
	}
}

func menuItemToResponse(menuItem domain.MenuItem) restaurantapi.ReturnMenuItem {
	return restaurantapi.ReturnMenuItem{
		ID:     menuItem.ID,
		Name:   menuItem.Name,
		Price:  menuItem.Price,
		Course: menuItem.Course,
	}
}
