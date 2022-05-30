package Restaurantapi

type NewRestaurantAccountRequest struct {
	RestaurantID uint64 `json:"restaurant_id" yaml:"restaurant_id"`
}

type AddRestaurantAccountRequest struct {
	RestaurantID uint64 `json:"restaurant_id" yaml:"restaurant_id"`
	Amount       int64  `json:"amount" yaml:"amount"`
}

type SubBalanceRestaurantAccountRequest struct {
	RestaurantID uint64 `json:"restaurant_id" yaml:"restaurant_id"`
	Amount       int64  `json:"amount" yaml:"amount"`
}
