package service

import "github.com/nndergunov/deliveryApp/app/services/delivery/pkg/domain"

type RestaurantClient interface {
	GetRestaurant(restaurantID int) (*domain.Restaurant, error)
}

type CourierClient interface {
	GetNearestCourier(location *domain.Location, radiusKm int) (*domain.Courier, error)
	UpdateCourierAvailable(courierID int, available bool) (*domain.Courier, error)
}

type ConsumerClient interface {
	GetConsumerLocation(consumerID int) (*domain.ConsumerLocation, error)
}
