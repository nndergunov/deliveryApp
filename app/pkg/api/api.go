package api

import (
	"net/http"

	"github.com/nndergunov/deliveryApp/app/pkg/logger"
)

// API is main server handler.
type API struct {
	mux *http.ServeMux
	log *logger.Logger
}

// NewAPI returns new instance of api.API.
func NewAPI(endpointHandler *http.ServeMux, log *logger.Logger) *API {
	api := &API{
		mux: endpointHandler,
		log: log,
	}

	return api
}

// ServeHTTP method satisfies http.Handler interface.
func (a *API) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	a.mux.ServeHTTP(w, r)
}
