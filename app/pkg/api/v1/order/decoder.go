package order

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
