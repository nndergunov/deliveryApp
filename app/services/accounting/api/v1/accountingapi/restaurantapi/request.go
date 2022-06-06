package Restaurantapi

type NewRestaurantAccountRequest struct {
	RestaurantID int `json:"restaurant_id" yaml:"restaurant_id"`
}

type AddRestaurantAccountRequest struct {
	Amount int `json:"amount" yaml:"amount"`
}

type SubBalanceRestaurantAccountRequest struct {
	Amount int `json:"amount" yaml:"amount"`
}
