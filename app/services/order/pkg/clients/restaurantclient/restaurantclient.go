package restaurantclient

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	v1 "github.com/nndergunov/deliveryApp/app/pkg/api/v1"
	"github.com/nndergunov/deliveryApp/app/pkg/api/v1/restaurantapi"
	"github.com/nndergunov/deliveryApp/app/services/order/pkg/domain"
)

type RestaurantClient struct {
	restaurantURL string
}

func NewRestaurantClient(url string) *RestaurantClient {
	return &RestaurantClient{restaurantURL: url}
}

func (r RestaurantClient) CheckIfAvailable(restaurantID int) (bool, error) {
	resp, err := http.Get(r.restaurantURL + "/v1/restaurants/" + strconv.Itoa(restaurantID))
	if err != nil {
		return false, fmt.Errorf("getting restaurant: %w", err)
	}

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return false, fmt.Errorf("getting response body: %w", err)
	}

	err = resp.Body.Close()
	if err != nil {
		return false, fmt.Errorf("closing response body: %w", err)
	}

	restaurant := new(restaurantapi.ReturnRestaurant)

	err = v1.Decode(respBody, restaurant)
	if err != nil {
		return false, fmt.Errorf("decoding response: %w", err)
	}

	return restaurant.AcceptingOrders, nil
}

func (r RestaurantClient) CalculateOrderPrice(order domain.Order) (float64, error) {
	resp, err := http.Get(r.restaurantURL + "/v1/restaurants/" + strconv.Itoa(order.RestaurantID) + "/menu")
	if err != nil {
		return 0, fmt.Errorf("getting menu: %w", err)
	}

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return 0, fmt.Errorf("getting response body: %w", err)
	}

	err = resp.Body.Close()
	if err != nil {
		return 0, fmt.Errorf("closing response body: %w", err)
	}

	menu := new(restaurantapi.ReturnMenu)

	err = v1.Decode(respBody, menu)
	if err != nil {
		return 0, fmt.Errorf("decoding response: %w", err)
	}

	var price float64

	for itemID, quantity := range order.OrderItems {
		price += menu.MenuItems[itemID].Price * float64(quantity)
	}

	return price, nil
}
