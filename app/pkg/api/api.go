package api

import (
	"net/http"

	"github.com/nndergunov/deliveryApp/app/pkg/logger"
)

type Multiplexer interface {
	ServeHTTP(w http.ResponseWriter, r *http.Request)
}

// API is main server handler.
type API struct {
	serveMux Multiplexer
	log      *logger.Logger
}

// NewAPI returns new instance of api.API.
func NewAPI(endpointHandler Multiplexer, log *logger.Logger) *API {
	api := &API{
		serveMux: endpointHandler,
		log:      log,
	}

	return api
}

// ServeHTTP method satisfies http.Handler interface.
func (a *API) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	a.serveMux.ServeHTTP(w, r)
}
