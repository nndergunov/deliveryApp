package orderapi

import (
	"encoding/json"
	"fmt"
)

func DecodeCreateOrder(data []byte) (*CreateOrder, error) {
	var req *CreateOrder

	err := json.Unmarshal(data, req)
	if err != nil {
		return nil, fmt.Errorf("DecodeCreateOrder: %w", err)
	}

	return req, nil
}

func DecodeOrderStatusUpdate(data []byte) (*OrderStatusUpdate, error) {
	var req *OrderStatusUpdate

	err := json.Unmarshal(data, req)
	if err != nil {
		return nil, fmt.Errorf("DecodeCreateOrder: %w", err)
	}

	return req, nil
}

func DecodeReturnOrders(data []byte) (*ReturnOrders, error) {
	var req *ReturnOrders

	err := json.Unmarshal(data, req)
	if err != nil {
		return nil, fmt.Errorf("DecodeReturnOrders: %w", err)
	}

	return req, nil
}

func DecodeReturnOrderStatus(data []byte) (*ReturnOrderStatus, error) {
	var req *ReturnOrderStatus

	err := json.Unmarshal(data, req)
	if err != nil {
		return nil, fmt.Errorf("DecodeReturnOrderStatus: %w", err)
	}

	return req, nil
}
