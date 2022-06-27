package handlers

import (
	"errors"
	"net/http"

	v1 "github.com/nndergunov/deliveryApp/app/pkg/api/v1"
	"github.com/nndergunov/deliveryApp/app/services/restaurant/pkg/service"
)

var errNoIDInEndpoint = errors.New("id not found in endpoint")

func (e endpointHandler) handleError(err error, responseWriter http.ResponseWriter) {
	e.log.Println(err)

	switch {
	case errors.Is(err, service.ErrItemIsNotInMenu):
		err := v1.RespondWithError("requested item is not in menu", http.StatusBadRequest, responseWriter)
		if err != nil {
			e.log.Println(err)
		}
	default:
		responseWriter.WriteHeader(http.StatusBadRequest)
	}
}
