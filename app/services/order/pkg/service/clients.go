package service

import "github.com/nndergunov/deliveryApp/app/services/order/pkg/domain"

// AccountingClient interface outlines accounting client functionality.
type AccountingClient interface {
	CreateTransaction(accountID, restaurantID int, orderPrice float64) (bool, error)
}

// RestaurantClient interface outlines restaurant client functionality.
type RestaurantClient interface {
	CheckIfAvailable(restaurantID int) (bool, error)
	CalculateOrderPrice(order domain.Order) (float64, error)
}
