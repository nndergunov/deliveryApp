package restaurantgrp

import (
	"accounting/api/v1/accountingapi/restaurantapi"
	"accounting/pkg/domain"
)

func requestToNewRestaurantAccount(req *Restaurantapi.NewRestaurantAccountRequest) domain.RestaurantAccount {
	return domain.RestaurantAccount{
		RestaurantID: req.RestaurantID,
	}
}

func requestToAddBalanceRestaurantAccount(req *Restaurantapi.AddRestaurantAccountRequest) domain.RestaurantAccount {
	return domain.RestaurantAccount{
		RestaurantID: req.RestaurantID,
		Balance:      req.Amount,
	}
}

func requestToSubBalanceRestaurantAccount(req *Restaurantapi.SubBalanceRestaurantAccountRequest) domain.RestaurantAccount {
	return domain.RestaurantAccount{
		RestaurantID: req.RestaurantID,
		Balance:      req.Amount,
	}
}
