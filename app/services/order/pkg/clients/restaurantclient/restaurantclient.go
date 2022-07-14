// Package restaurantclient implements communication with restaurant service.
package restaurantclient

import (
	"fmt"
	"net/http"
	"strconv"

	v1 "github.com/nndergunov/deliveryApp/app/pkg/api/v1"
	"github.com/nndergunov/deliveryApp/app/services/order/pkg/domain"
	"github.com/nndergunov/deliveryApp/app/services/restaurant/api/v1/restaurantapi"
)

// RestaurantClient is responsible for communicating with restaurant service.
type RestaurantClient struct {
	restaurantURL string
}

// NewRestaurantClient returns new RestaurantClient instance.
func NewRestaurantClient(url string) *RestaurantClient {
	return &RestaurantClient{restaurantURL: url}
}

// CheckIfAvailable returns whether the restaurant can accept orders.
func (r RestaurantClient) CheckIfAvailable(restaurantID int) (bool, error) {
	resp, err := http.Get(r.restaurantURL + "/v1/restaurants/" + strconv.Itoa(restaurantID))
	if err != nil {
		return false, fmt.Errorf("getting restaurant: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return false, fmt.Errorf("%w: response code %d", ErrRestaurantFail, resp.StatusCode)
	}

	restaurant := new(restaurantapi.ReturnRestaurant)

	err = v1.DecodeResponse(resp, restaurant)
	if err != nil {
		return false, fmt.Errorf("decoding response: %w", err)
	}

	return restaurant.AcceptingOrders, nil
}

// CalculateOrderPrice returns the price of the order.
func (r RestaurantClient) CalculateOrderPrice(order domain.Order) (float64, error) {
	resp, err := http.Get(r.restaurantURL + "/v1/restaurants/" + strconv.Itoa(order.RestaurantID) + "/menu")
	if err != nil {
		return 0, fmt.Errorf("getting menu: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return 0, fmt.Errorf("%w: response code %d", ErrRestaurantFail, resp.StatusCode)
	}

	menu := new(restaurantapi.ReturnMenu)

	err = v1.DecodeResponse(resp, menu)
	if err != nil {
		return 0, fmt.Errorf("decoding response: %w", err)
	}

	var price float64

	for _, itemID := range order.OrderItems {
		for _, menuItem := range menu.MenuItems {
			if menuItem.ID != itemID {
				continue
			}

			price += menuItem.Price
		}
	}

	return price, nil
}
