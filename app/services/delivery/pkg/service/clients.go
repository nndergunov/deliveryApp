package service

import "delivery/pkg/domain"

type RestaurantClient interface {
	GetRestaurantLocation(restaurantID int) (*domain.Location, error)
}

type CourierClient interface {
	GetNearestCourier(location *domain.Location, radius int) (*domain.Courier, error)
	UpdateCourierAvailable(courierID int, available bool) (*domain.Courier, error)
}
