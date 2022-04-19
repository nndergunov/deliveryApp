package accountingapi

import (
	"encoding/json"
	"fmt"
)

func DecodeReturnPaymentSum(data []byte) (*ReturnPaymentSum, error) {
	var req *ReturnPaymentSum

	err := json.Unmarshal(data, req)
	if err != nil {
		return nil, fmt.Errorf("DecodeReturnPaymentSum: %w", err)
	}

	return req, nil
}

func DecodePaymentSum(data []byte) (*PaymentSum, error) {
	var req *PaymentSum

	err := json.Unmarshal(data, req)
	if err != nil {
		return nil, fmt.Errorf("DecodePaymentSum: %w", err)
	}

	return req, nil
}
