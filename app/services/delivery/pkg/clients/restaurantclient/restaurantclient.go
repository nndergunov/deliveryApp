package restaurantclient

import (
	"delivery/pkg/domain"
	"fmt"
	"net/http"
	"strconv"
)

type RestaurantClient struct {
	restaurantURL string
}

func NewRestaurantClient(url string) *RestaurantClient {
	return &RestaurantClient{restaurantURL: url}
}

func (a RestaurantClient) GetRestaurantLocation(restaurantID int) (*domain.Location, error) {
	_, err := http.Get(a.restaurantURL + "v1/restaurants/location/" + strconv.Itoa(restaurantID))
	if err != nil {
		return nil, fmt.Errorf("sending request: %w", err)
	}
	//todo when restaurant add this rout

	return &domain.Location{}, nil
}
