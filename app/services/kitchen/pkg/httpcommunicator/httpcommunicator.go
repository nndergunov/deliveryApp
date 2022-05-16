package httpcommunicator

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"

	v1 "github.com/nndergunov/deliveryApp/app/pkg/api/v1"
	"github.com/nndergunov/deliveryApp/app/pkg/api/v1/orderapi"
)

type HTTPCommunicator struct {
	orderServiceBaseURL string
}

func NewHTTPCommunicator(orderServiceBaseURL string) *HTTPCommunicator {
	return &HTTPCommunicator{orderServiceBaseURL: orderServiceBaseURL}
}

func (c HTTPCommunicator) GetIncompleteOrders(id int) (*orderapi.ReturnOrderList, error) {
	filters := orderapi.OrderFilters{
		FromRestaurantID: id,
		Statuses:         nil,
		ExcludeStatuses:  []string{orderapi.Complete},
	}

	requestData, err := v1.Encode(filters)
	if err != nil {
		return nil, fmt.Errorf("encoding HTTP request: %w", err)
	}

	req, err := http.NewRequest(http.MethodGet, c.orderServiceBaseURL+"/v1/orders", bytes.NewBuffer(requestData))
	if err != nil {
		return nil, fmt.Errorf("making HTTP request: %w", err)
	}

	client := http.DefaultClient

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("sending HTTP request: %w", err)
	}

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("reaking HTTP response: %w", err)
	}

	orders := new(orderapi.ReturnOrderList)

	err = v1.Decode(respBody, orders)
	if err != nil {
		return nil, fmt.Errorf("decoding HTTP response: %w", err)
	}

	return orders, nil
}
