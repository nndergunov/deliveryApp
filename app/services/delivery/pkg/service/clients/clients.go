package clients

import (
	"github.com/nndergunov/deliveryApp/app/services/consumer/api/v1/rest/consumerapi"
	pb "github.com/nndergunov/deliveryApp/app/services/courier/api/v1/grpc/proto"
	"github.com/nndergunov/deliveryApp/app/services/restaurant/api/v1/restaurantapi"
)

type RestaurantClient interface {
	GetRestaurant(restaurantID int) (*restaurantapi.ReturnRestaurant, error)
}

type CourierClient interface {
	GetLocation(city string) (*pb.LocationList, error)
	UpdateCourierAvailable(courierID int, available string) (*pb.CourierResponse, error)
}

type ConsumerClient interface {
	GetLocation(consumerID int) (*consumerapi.LocationResponse, error)
}
