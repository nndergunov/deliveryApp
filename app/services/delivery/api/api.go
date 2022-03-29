package api

import (
	"log"
	"net/http"
)

type API struct {
	mux *http.ServeMux
	log *log.Logger
}

func NewAPI(log *log.Logger) *API {
	mux := http.NewServeMux()

	a := &API{
		mux: mux,
		log: log,
	}

	a.handlerInit()

	return a
}

func (a *API) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	a.mux.ServeHTTP(w, r)
}
