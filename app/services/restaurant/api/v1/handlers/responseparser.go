package handlers

import (
	"github.com/nndergunov/deliveryApp/app/pkg/api/v1/restaurantapi"
	"github.com/nndergunov/deliveryApp/app/services/restaurant/pkg/domain"
)

func RestaurantListToResponse(restaurants []domain.Restaurant) restaurantapi.ReturnRestaurantList {
	list := make([]struct {
		ID      int
		Name    string
		City    string
		Address string
	}, 0, len(restaurants))

	for _, restaurant := range restaurants {
		list = append(list, struct {
			ID      int
			Name    string
			City    string
			Address string
		}{
			ID:      restaurant.ID,
			Name:    restaurant.Name,
			City:    restaurant.City,
			Address: restaurant.Address,
		})
	}

	return restaurantapi.ReturnRestaurantList{
		List: list,
	}
}
