package restaurantclient

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/nndergunov/deliveryApp/app/services/restaurant/api/v1/restaurantapi"

	"github.com/nndergunov/deliveryApp/app/services/delivery/api/v1/rest/deliveryapi"
)

type RestaurantClient struct {
	restaurantURL string
}

func NewRestaurantClient(url string) *RestaurantClient {
	return &RestaurantClient{restaurantURL: url}
}

func (a RestaurantClient) GetRestaurant(restaurantID int) (*restaurantapi.ReturnRestaurant, error) {
	resp, err := http.Get(a.restaurantURL + "/v1/restaurants/" + strconv.Itoa(restaurantID))
	if err != nil {
		return nil, fmt.Errorf("sending request: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("not ok status: %v", resp.StatusCode)
	}

	restaurantData := restaurantapi.ReturnRestaurant{}
	if err = deliveryapi.DecodeJSON(resp.Body, &restaurantData); err != nil {
		return nil, fmt.Errorf("decoding : %w", err)
	}

	if err := resp.Body.Close(); err != nil {
		return nil, err
	}

	return &restaurantData, nil
}
