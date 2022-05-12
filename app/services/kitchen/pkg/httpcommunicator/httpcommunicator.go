package httpcommunicator

import "github.com/nndergunov/deliveryApp/app/pkg/api/v1/orderapi"

type HTTPCommunicator struct {
	orderServiceBaseURL string
}

func NewHTTPCommunicator(orderServiceBaseURL string) *HTTPCommunicator {
	return &HTTPCommunicator{orderServiceBaseURL: orderServiceBaseURL}
}

func (c HTTPCommunicator) GetRestaurantIncompleteOrders(restaurantID int) (orderapi.ReturnOrderList, error) {

}
