package handlers

import (
	"errors"
	"net/http"

	v1 "github.com/nndergunov/deliveryApp/app/pkg/api/v1"
	"github.com/nndergunov/deliveryApp/app/services/order/pkg/service"
)

var errNoIDInEndpoint = errors.New("id not found in endpoint")

func (e endpointHandler) handleError(err error, responseWriter http.ResponseWriter) {
	e.log.Println(err)

	switch {
	case errors.Is(err, service.ErrRestaurantOffline):
		err := v1.RespondWithError("restaurant is not accepting orders", http.StatusBadRequest, responseWriter)
		if err != nil {
			e.log.Println(err)
		}
	case errors.Is(err, service.ErrLowBalance):
		err := v1.RespondWithError("user has not enough balance to create this order", http.StatusBadRequest, responseWriter)
		if err != nil {
			e.log.Println(err)
		}
	default:
		responseWriter.WriteHeader(http.StatusInternalServerError)
	}
}
