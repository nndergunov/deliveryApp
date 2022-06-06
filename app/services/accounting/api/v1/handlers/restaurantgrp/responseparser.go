package restaurantgrp

import (
	"accounting/api/v1/accountingapi/restaurantapi"
	"accounting/pkg/domain"
)

func RestaurantAccountToResponse(restaurantAccount domain.RestaurantAccount) Restaurantapi.RestaurantAccountResponse {
	return Restaurantapi.RestaurantAccountResponse{
		RestaurantID: restaurantAccount.RestaurantID,
		Balance:      restaurantAccount.Balance,
	}
}
