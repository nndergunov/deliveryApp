package accounting

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
