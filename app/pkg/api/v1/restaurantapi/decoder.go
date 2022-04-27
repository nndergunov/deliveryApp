package restaurantapi

import (
	"encoding/json"
	"fmt"
)

func DecodeRestaurantData(data []byte) (*RestaurantData, error) {
	req := new(RestaurantData)

	err := json.Unmarshal(data, req)
	if err != nil {
		return nil, fmt.Errorf("DecodeCreateRestaurant: %w", err)
	}

	return req, nil
}

func DecodeMenuData(data []byte) (*MenuData, error) {
	req := new(MenuData)

	err := json.Unmarshal(data, req)
	if err != nil {
		return nil, fmt.Errorf("DecodeCreateMenu: %w", err)
	}

	return req, nil
}

func DecodeMenuItem(data []byte) (*MenuItemData, error) {
	req := new(MenuItemData)

	err := json.Unmarshal(data, req)
	if err != nil {
		return nil, fmt.Errorf("DecodeCreateMenu: %w", err)
	}

	return req, nil
}

func DecodeReturnRestaurant(data []byte) (*ReturnRestaurant, error) {
	req := new(ReturnRestaurant)

	err := json.Unmarshal(data, req)
	if err != nil {
		return nil, fmt.Errorf("DecodeCreateMenu: %w", err)
	}

	return req, nil
}

func DecodeReturnRestaurantList(data []byte) (*ReturnRestaurantList, error) {
	req := new(ReturnRestaurantList)

	err := json.Unmarshal(data, req)
	if err != nil {
		return nil, fmt.Errorf("DecodeCreateMenu: %w", err)
	}

	return req, nil
}

func DecodeReturnMenu(data []byte) (*ReturnMenu, error) {
	req := new(ReturnMenu)

	err := json.Unmarshal(data, req)
	if err != nil {
		return nil, fmt.Errorf("DecodeCreateMenu: %w", err)
	}

	return req, nil
}

func DecodeReturnMenuItem(data []byte) (*ReturnMenuItem, error) {
	req := new(ReturnMenuItem)

	err := json.Unmarshal(data, req)
	if err != nil {
		return nil, fmt.Errorf("DecodeCreateMenu: %w", err)
	}

	return req, nil
}
