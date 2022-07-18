// Package ordersclient implements client that is used to send requests to the
// order service.
package ordersclient

import (
	"bytes"
	"fmt"
	"net/http"

	v1 "github.com/nndergunov/deliveryApp/app/pkg/api/v1"
	"github.com/nndergunov/deliveryApp/app/services/order/api/v1/orderapi"
)

// OrdersClient is used to comunicate with the order service.
type OrdersClient struct {
	orderServiceBaseURL string
}

// NewOrdersClient creates new OrdersClient instance.
func NewOrdersClient(orderServiceBaseURL string) *OrdersClient {
	return &OrdersClient{orderServiceBaseURL: orderServiceBaseURL}
}

// GetIncompleteOrders returns all orders that are not marked as complete.
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

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("%w: http status: %d", ErrOrderServiceFailed, resp.StatusCode)
	}

	orders := new(orderapi.ReturnOrderList)

	err = v1.DecodeResponse(resp, orders)
	if err != nil {
		return nil, fmt.Errorf("decoding HTTP response: %w", err)
	}

	return orders, nil
}
