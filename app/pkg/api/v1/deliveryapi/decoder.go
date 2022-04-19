package deliveryapi

import (
	"encoding/json"
	"fmt"
)

func DecodeFindCourier(data []byte) (*FindCourier, error) {
	var req *FindCourier

	err := json.Unmarshal(data, req)
	if err != nil {
		return nil, fmt.Errorf("DecodeFindCourier: %w", err)
	}

	return req, nil
}
