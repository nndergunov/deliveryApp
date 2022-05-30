package handlers

import (
	"accounting/api/v1/handlers/routs"
	"accounting/api/v1/web"
	"accounting/pkg/service/consumerservice"
	"accounting/pkg/service/courierservice"
	"accounting/pkg/service/restaurantrservice"
	"github.com/gorilla/mux"
	"github.com/nndergunov/deliveryApp/app/pkg/logger"
)

type Params struct {
	Logger            *logger.Logger
	ConsumerService   consumerservice.ConsumerService
	CourierService    courierservice.CourierService
	RestaurantService restaurantservice.RestaurantService
}

// NewAPIMux returns new http multiplexer with configured endpoints.
func NewAPIMux(p Params) *mux.Router {

	serveMux := mux.NewRouter()

	app := &web.AppMux{
		ServeMux:          serveMux,
		Log:               p.Logger,
		ConsumerService:   p.ConsumerService,
		CourierService:    p.CourierService,
		RestaurantService: p.RestaurantService,
	}

	routs.InitRoutes(app)

	return app.ServeMux
}
