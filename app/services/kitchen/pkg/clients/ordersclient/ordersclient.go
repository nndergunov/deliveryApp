package ordersclient

import (
	"bytes"
	"fmt"
	"net/http"

	v1 "github.com/nndergunov/deliveryApp/app/pkg/api/v1"
	"github.com/nndergunov/deliveryApp/app/services/order/api/v1/communication"
)

type OrdersClient struct {
	orderServiceBaseURL string
}

func NewOrdersClient(orderServiceBaseURL string) *OrdersClient {
	return &OrdersClient{orderServiceBaseURL: orderServiceBaseURL}
}

func (c OrdersClient) GetIncompleteOrders(id int) (*communication.ReturnOrderList, error) {
	filters := communication.OrderFilters{
		FromRestaurantID: &id,
		Statuses:         nil,
		ExcludeStatuses:  []string{communication.Complete},
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

	orders := new(communication.ReturnOrderList)

	err = v1.DecodeResponse(resp, orders)
	if err != nil {
		return nil, fmt.Errorf("decoding HTTP response: %w", err)
	}

	return orders, nil
}
