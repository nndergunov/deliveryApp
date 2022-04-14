package kitchenapi

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
		return nil, fmt.Errorf("DecodeCreateMenu: %w", err)
	}

	return req, nil
}

func DecodeRestaurantList(data []byte) (*RestaurantList, error) {
	var req *RestaurantList

	err := json.Unmarshal(data, req)
	if err != nil {
		return nil, fmt.Errorf("DecodeRestaurantList: %w", err)
	}

	return req, nil
}

func DecodeMenu(data []byte) (*Menu, error) {
	var req *Menu

	err := json.Unmarshal(data, req)
	if err != nil {
		return nil, fmt.Errorf("DecodeMenu: %w", err)
	}

	return req, nil
}
