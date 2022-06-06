package Restaurantapi

type RestaurantAccountResponse struct {
	RestaurantID int `json:"restaurant_id,omitempty" yaml:"restaurant_id,omitempty"`
	Balance      int `json:"balance" yaml:"balance"`
}

type ReturnAccountingResponseList struct {
	AccountingResponseList []RestaurantAccountResponse
}
