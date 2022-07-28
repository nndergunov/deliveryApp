package clients

import (
	pbConsumer "github.com/nndergunov/deliveryApp/app/services/consumer/api/v1/grpc/proto"
	pbCourier "github.com/nndergunov/deliveryApp/app/services/courier/api/v1/grpc/proto"
	pbRes "github.com/nndergunov/deliveryApp/app/services/restaurant/api/v1/grpclogic/pb"
)

type RestaurantClient interface {
	GetRestaurant(restaurantID int) (*pbRes.RestaurantResponse, error)
}

type CourierClient interface {
	GetLocation(city string) (*pbCourier.LocationList, error)
	UpdateCourierAvailable(courierID string, available string) (*pbCourier.CourierResponse, error)
}

type ConsumerClient interface {
	GetLocation(id int64) (*pbConsumer.Location, error)
}
