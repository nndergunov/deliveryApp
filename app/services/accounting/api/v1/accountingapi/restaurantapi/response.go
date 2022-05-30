package Restaurantapi

type RestaurantAccountResponse struct {
	RestaurantID uint64 `json:"restaurant_id,omitempty" yaml:"restaurant_id,omitempty"`
	Balance      int64  `json:"balance" yaml:"balance"`
}

type ReturnAccountingResponseList struct {
	AccountingResponseList []RestaurantAccountResponse
}
