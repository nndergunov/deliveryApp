package restaurantclient

import (
	"fmt"
	"net/http"
	"strconv"

	"delivery/pkg/domain"
)

type RestaurantClient struct {
	restaurantURL string
}

func NewRestaurantClient(url string) *RestaurantClient {
	return &RestaurantClient{restaurantURL: url}
}

func (a RestaurantClient) GetRestaurant(restaurantID int) (*domain.Restaurant, error) {
	_, err := http.Get(a.restaurantURL + "v1/restaurants/" + strconv.Itoa(restaurantID))
	if err != nil {
		return nil, fmt.Errorf("sending request: %w", err)
	}
	// todo when restaurant add this rout

	return &domain.Restaurant{}, nil
}
