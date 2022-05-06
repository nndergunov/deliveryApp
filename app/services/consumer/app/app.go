package app

import (
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/nndergunov/deliveryApp/app/pkg/configreader"
	"github.com/nndergunov/deliveryApp/app/pkg/logger"
)

// Params is the input parameter struct for the handlers module.
type Params struct {
	Logger   *logger.Logger
	Shutdown chan os.Signal
}

// NewHandlerServer create new mux and server
func NewHandlerServer(p Params) (router *mux.Router, server *http.Server, err error) {
	router = mux.NewRouter()

	port := configreader.GetString("Server.dev.port")
	if port == "" {
		port = ":7070"
	}

	server = &http.Server{
		Addr:           port,
		Handler:        router,
		ReadTimeout:    20 * time.Second,
		WriteTimeout:   20 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	return router, server, nil
}
