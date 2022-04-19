package restaurantapi

import (
	"encoding/json"
	"fmt"
)

func DecodeRestaurantData(data []byte) (*RestaurantData, error) {
	var req *RestaurantData

	err := json.Unmarshal(data, req)
	if err != nil {
		return nil, fmt.Errorf("DecodeCreateRestaurant: %w", err)
	}

	return req, nil
}

func DecodeMenuData(data []byte) (*MenuData, error) {
	var req *MenuData

	err := json.Unmarshal(data, req)
	if err != nil {
		return nil, fmt.Errorf("DecodeCreateMenu: %w", err)
	}

	return req, nil
}

func DecodeMenuItem(data []byte) (*MenuItemData, error) {
	var req *MenuItemData

	err := json.Unmarshal(data, req)
	if err != nil {
		return nil, fmt.Errorf("DecodeCreateMenu: %w", err)
	}

	return req, nil
}
