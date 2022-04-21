package handlers

import (
	"net/http"

	v1 "github.com/nndergunov/deliveryApp/app/pkg/api/v1"
	"github.com/nndergunov/deliveryApp/app/pkg/api/v1/restaurantapi"
	"github.com/nndergunov/deliveryApp/app/services/restaurant/pkg/domain"
)

func (e endpointHandler) respond(response any, w http.ResponseWriter) {
	data, err := v1.Encode(response)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)

		return
	}

	_, err = w.Write(data)

	w.WriteHeader(http.StatusOK)
}

func restaurantListToResponse(restaurants []domain.Restaurant) restaurantapi.ReturnRestaurantList {
	list := make([]restaurantapi.ReturnRestaurant, 0, len(restaurants))

	for _, restaurant := range restaurants {
		currElement := restaurantapi.ReturnRestaurant{
			ID:      restaurant.ID,
			Name:    restaurant.Name,
			City:    restaurant.City,
			Address: restaurant.Address,
		}

		list = append(list, currElement)
	}

	return restaurantapi.ReturnRestaurantList{
		List: list,
	}
}

func menuToResponse(menu domain.Menu) restaurantapi.ReturnMenu {
	items := make([]restaurantapi.ReturnMenuItem, 0, len(menu.Items))

	for _, menuItem := range menu.Items {
		currItem := restaurantapi.ReturnMenuItem{
			ID:   menuItem.ID,
			Name: menuItem.Name,
		}

		items = append(items, currItem)
	}

	return restaurantapi.ReturnMenu{
		RestaurantID: menu.RestaurantID,
		Items:        items,
	}
}
