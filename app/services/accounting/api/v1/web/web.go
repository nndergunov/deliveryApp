package web

import (
	"github.com/gorilla/mux"
	"github.com/nndergunov/deliveryApp/app/pkg/logger"

	"accounting/pkg/service/consumerservice"
	"accounting/pkg/service/courierservice"
	"accounting/pkg/service/restaurantrservice"
)

// AppMux is the entrypoint into our application
type AppMux struct {
	ServeMux          *mux.Router
	Log               *logger.Logger
	ConsumerService   consumerservice.ConsumerService
	CourierService    courierservice.CourierService
	RestaurantService restaurantservice.RestaurantService
}
