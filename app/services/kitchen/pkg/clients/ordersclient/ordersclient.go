package ordersclient

import (
	"bytes"
	"fmt"
	"net/http"

	v1 "github.com/nndergunov/deliveryApp/app/pkg/api/v1"
	"github.com/nndergunov/deliveryApp/app/pkg/api/v1/orderapi"
)

type OrdersClient struct {
	orderServiceBaseURL string
}

func NewOrdersClient(orderServiceBaseURL string) *OrdersClient {
	return &OrdersClient{orderServiceBaseURL: orderServiceBaseURL}
}

func (c OrdersClient) GetIncompleteOrders(id int) (*orderapi.ReturnOrderList, error) {
	filters := orderapi.OrderFilters{
		FromRestaurantID: &id,
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

	orders := new(orderapi.ReturnOrderList)

	err = v1.DecodeResponse(resp, orders)
	if err != nil {
		return nil, fmt.Errorf("decoding HTTP response: %w", err)
	}

	return orders, nil
}
