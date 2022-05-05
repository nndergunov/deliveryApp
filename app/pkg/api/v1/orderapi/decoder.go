package orderapi

import (
	"encoding/json"
	"fmt"
)

func DecodeOrderData(data []byte) (*OrderData, error) {
	req := new(OrderData)

	err := json.Unmarshal(data, req)
	if err != nil {
		return nil, fmt.Errorf("DecodeOrderData: %w", err)
	}

	return req, nil
}

func DecodeOrderStatusData(data []byte) (*OrderStatusData, error) {
	req := new(OrderStatusData)

	err := json.Unmarshal(data, req)
	if err != nil {
		return nil, fmt.Errorf("DecodeOrderStatusData: %w", err)
	}

	return req, nil
}

func DecodeReturnOrder(data []byte) (*ReturnOrder, error) {
	req := new(ReturnOrder)

	err := json.Unmarshal(data, req)
	if err != nil {
		return nil, fmt.Errorf("DecodeReturnOrder: %w", err)
	}

	return req, nil
}

func DecodeReturnOrderList(data []byte) (*ReturnOrderList, error) {
	req := new(ReturnOrderList)

	err := json.Unmarshal(data, req)
	if err != nil {
		return nil, fmt.Errorf("DecodeReturnOrderList: %w", err)
	}

	return req, nil
}
