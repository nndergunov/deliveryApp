package clients

import (
	"github.com/nndergunov/deliveryApp/app/services/consumer/api/v1/rest/consumerapi"
	"github.com/nndergunov/deliveryApp/app/services/courier/api/v1/rest/courierapi"
	"github.com/nndergunov/deliveryApp/app/services/restaurant/api/v1/restaurantapi"
)

type RestaurantClient interface {
	GetRestaurant(restaurantID int) (*restaurantapi.ReturnRestaurant, error)
}

type CourierClient interface {
	GetCourier(courierID int) (*courierapi.CourierResponse, error)
	GetLocation(city string) (*courierapi.LocationResponseList, error)
	UpdateCourierAvailable(courierID int, available string) (*courierapi.CourierResponse, error)
}

type ConsumerClient interface {
	GetLocation(consumerID int) (*consumerapi.LocationResponse, error)
}
