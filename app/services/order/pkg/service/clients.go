package service

import "github.com/nndergunov/deliveryApp/app/services/order/pkg/domain"

type AccountingClient interface {
	CheckIfEnoughBalance(order domain.Order) (bool, error)
}

type RestaurantClient interface {
	CheckIfAvailable(restaurantID int) (bool, error)
}
