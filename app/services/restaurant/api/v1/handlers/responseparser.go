package handlers

import (
	"net/http"

	"github.com/nndergunov/deliveryApp/app/pkg/api/v1/restaurantapi"
	"github.com/nndergunov/deliveryApp/app/services/restaurant/pkg/domain"
	v1 "github.com/nndergunov/deliveryApp/service/pkg/api/v1"
)

func (e endpointHandler) respond(response any, w http.ResponseWriter) {
	data, err := v1.Encode(response)
	if err != nil {
		e.log.Println(err)

		w.WriteHeader(http.StatusInternalServerError)

		return
	}

	w.WriteHeader(http.StatusOK)

	_, err = w.Write(data)
	if err != nil {
		e.log.Println(err)

		w.WriteHeader(http.StatusInternalServerError)

		return
	}
}

func restaurantToResponse(restaurant domain.Restaurant) restaurantapi.ReturnRestaurant {
	return restaurantapi.ReturnRestaurant{
		ID:      restaurant.ID,
		Name:    restaurant.Name,
		City:    restaurant.City,
		Address: restaurant.Address,
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
		Items:        items,
	}
}

func menuItemToResponse(menuItem domain.MenuItem) restaurantapi.ReturnMenuItem {
	return restaurantapi.ReturnMenuItem{
		ID:     menuItem.ID,
		Name:   menuItem.Name,
		Course: menuItem.Course,
	}
}
