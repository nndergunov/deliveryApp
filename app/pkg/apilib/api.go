package apilib

import (
	"net/http"

	"github.com/nndergunov/deliveryApp/app/pkg/logger"
)

type API struct {
	mux *http.ServeMux
	log *logger.Logger
}

func NewAPI(endpointHandler *http.ServeMux, log *logger.Logger) *API {
	api := &API{
		mux: endpointHandler,
		log: log,
	}

	return api
}

func (a *API) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	a.mux.ServeHTTP(w, r)
}
