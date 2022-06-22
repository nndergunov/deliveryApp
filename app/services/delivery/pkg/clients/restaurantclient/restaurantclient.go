package restaurantclient

import (
	"github.com/nndergunov/deliveryApp/app/pkg/api/v1/restaurantapi"
)

type RestaurantClient struct {
	restaurantURL string
}

func NewRestaurantClient(url string) *RestaurantClient {
	return &RestaurantClient{restaurantURL: url}
}

func (a RestaurantClient) GetRestaurant(restaurantID int) (*restaurantapi.ReturnRestaurant, error) {
	//resp, err := http.Get(a.restaurantURL + "v1/restaurants/" + strconv.Itoa(restaurantID))
	//if err != nil {
	//	return nil, fmt.Errorf("sending request: %w", err)
	//}
	//
	//if resp.StatusCode != http.StatusOK {
	//	return nil, fmt.Errorf("not ok status: %v", resp.StatusCode)
	//}
	//
	//restaurantData := restaurantapi.ReturnRestaurant{}
	//if err = deliveryapi.DecodeJSON(resp.Body, &restaurantData); err != nil {
	//	return nil, fmt.Errorf("decoding : %w", err)
	//}
	//
	//if err := resp.Body.Close(); err != nil {
	//	return nil, err
	//}

	restaurantData := restaurantapi.ReturnRestaurant{
		ID:              1,
		Name:            "Balboo",
		AcceptingOrders: true,
		City:            "Istanbul",
		Address:         "Bagcilar",
		Longitude:       28.860093289472598,
		Latitude:        41.036234457409066,
	}
	return &restaurantData, nil
}
