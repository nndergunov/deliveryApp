package kitchen

import (
	"encoding/json"
	"fmt"
)

func DecodeCreateRestaurant(data []byte) (*CreateRestaurant, error) {
	var req *CreateRestaurant

	err := json.Unmarshal(data, req)
	if err != nil {
		return nil, fmt.Errorf("DecodeCreateRestaurant: %w", err)
	}

	return req, nil
}

func DecodeCreateMenu(data []byte) (*CreateMenu, error) {
	var req *CreateMenu

	err := json.Unmarshal(data, req)
	if err != nil {
		return nil, fmt.Errorf("DecodeCreateRestaurant: %w", err)
	}

	return req, nil
}
