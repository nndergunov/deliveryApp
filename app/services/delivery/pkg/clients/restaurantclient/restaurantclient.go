package restaurantclient

import (
	"fmt"
	"github.com/nndergunov/deliveryApp/app/pkg/api/v1/deliveryapi"
	"github.com/nndergunov/deliveryApp/app/pkg/api/v1/restaurantapi"
	"net/http"
	"strconv"
)

type RestaurantClient struct {
	restaurantURL string
}

func NewRestaurantClient(url string) *RestaurantClient {
	return &RestaurantClient{restaurantURL: url}
}

func (a RestaurantClient) GetRestaurant(restaurantID int) (*restaurantapi.ReturnRestaurant, error) {
	resp, err := http.Get(a.restaurantURL + "v1/restaurants/" + strconv.Itoa(restaurantID))
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
