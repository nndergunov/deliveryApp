package Restaurantapi

type NewRestaurantAccountRequest struct {
	RestaurantID uint64 `json:"restaurant_id" yaml:"restaurant_id"`
}

type AddRestaurantAccountRequest struct {
	Amount int64 `json:"amount" yaml:"amount"`
}

type SubBalanceRestaurantAccountRequest struct {
	Amount int64 `json:"amount" yaml:"amount"`
}
