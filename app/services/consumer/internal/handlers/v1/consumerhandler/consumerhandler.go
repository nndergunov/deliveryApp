// Package handlers contains a small handlers framework extension.
package consumerhandler

import (
	"consumer/internal/handlers"
	"net/http"
	"os"
	"syscall"

	"consumer/internal/models"
	"consumer/internal/services"
	"consumer/pkg/decoder"

	"github.com/gorilla/mux"
	"github.com/nndergunov/deliveryApp/app/pkg/logger"
)

type Params struct {
	Logger          *logger.Logger
	ConsumerService services.ConsumerService
	Route           *mux.Router
	Shutdown        chan os.Signal
}

// ConsumerHandler is the entrypoint into our application
type ConsumerHandler struct {
	rout            *mux.Router
	log             *logger.Logger
	consumerService services.ConsumerService
	shutdown        chan os.Signal
}

// Handler is the interface for the carrier handlers.
type Handler interface {
	insertNewConsumer(rw http.ResponseWriter, r *http.Request)
	removeConsumer(rw http.ResponseWriter, r *http.Request)
	updateConsumer(rw http.ResponseWriter, r *http.Request)
	getAllConsumer(rw http.ResponseWriter, r *http.Request)
	getConsumer(rw http.ResponseWriter, r *http.Request)

	updateConsumerLocation(rw http.ResponseWriter, r *http.Request)
}

// NewConsumerHandler creates an ConsumerHandler value that handle a set of routes for the application.
func NewConsumerHandler(p Params) Handler {
	handler := &ConsumerHandler{
		log:             p.Logger,
		rout:            p.Route,
		consumerService: p.ConsumerService,
		shutdown:        p.Shutdown,
	}
	const apiVersion = "/v1"
	const consumer = "/consumer"
	const consumerLocation = "/consumer-location"

	p.Route.HandleFunc(apiVersion+consumer+"/new", handler.insertNewConsumer).Methods(http.MethodPost)
	p.Route.HandleFunc(apiVersion+consumer+"/remove/{id}", handler.removeConsumer).Methods(http.MethodPost)
	p.Route.HandleFunc(apiVersion+consumer+"/update", handler.updateConsumer).Methods(http.MethodPut)
	p.Route.HandleFunc(apiVersion+consumer+"/get-all", handler.getAllConsumer).Methods(http.MethodGet)
	p.Route.HandleFunc(apiVersion+consumer+"/get/{id}", handler.getConsumer).Methods(http.MethodGet)

	p.Route.HandleFunc(apiVersion+consumerLocation+"/update", handler.updateConsumerLocation).Methods(http.MethodPut)

	return handler
}

func (a *ConsumerHandler) insertNewConsumer(rw http.ResponseWriter, r *http.Request) {
	var consumer models.Consumer

	if err := decoder.BindJson(r, &consumer); err != nil {
		a.log.Println(err)
		if err := handlers.Respond(rw, "incorrect input data", http.StatusBadRequest); err != nil {
			a.log.Println(err)
			a.SignalShutdown()
		}
		return
	}

	data, statusCode := a.consumerService.InsertNewConsumer(consumer)
	if err := handlers.Respond(rw, data, statusCode); err != nil {
		a.log.Println(err)
		a.SignalShutdown()
		return
	}
}

func (a *ConsumerHandler) removeConsumer(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		if err := handlers.Respond(rw, "id is not indicated", http.StatusInternalServerError); err != nil {
			a.log.Println(err)
			a.SignalShutdown()
		}
		return
	}

	data, statusCode := a.consumerService.RemoveConsumer(id)
	if err := handlers.Respond(rw, data, statusCode); err != nil {
		a.log.Println(err)
		a.SignalShutdown()
		return
	}
}

func (a *ConsumerHandler) updateConsumer(rw http.ResponseWriter, r *http.Request) {
	var consumer models.Consumer

	if err := decoder.BindJson(r, &consumer); err != nil {
		a.log.Println(err)
		if err := handlers.Respond(rw, "incorrect input data", http.StatusBadRequest); err != nil {
			a.log.Println(err)
			a.SignalShutdown()
		}
		return
	}

	data, statusCode := a.consumerService.UpdateConsumer(consumer)
	if err := handlers.Respond(rw, data, statusCode); err != nil {
		a.log.Println(err)
		a.SignalShutdown()
		return
	}
}

func (a *ConsumerHandler) getAllConsumer(rw http.ResponseWriter, r *http.Request) {
	data, statusCode := a.consumerService.GetAllConsumer()
	if err := handlers.Respond(rw, data, statusCode); err != nil {
		a.log.Println(err)
		a.SignalShutdown()
		return
	}
}

func (a *ConsumerHandler) getConsumer(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		if err := handlers.Respond(rw, "id is missing in parameters", http.StatusInternalServerError); err != nil {
			a.log.Println(err)
			a.SignalShutdown()
			return
		}
		return
	}

	data, statusCode := a.consumerService.GetConsumer(id)
	if err := handlers.Respond(rw, data, statusCode); err != nil {
		a.log.Println(err)
		a.SignalShutdown()
		return
	}
}

func (a *ConsumerHandler) updateConsumerLocation(rw http.ResponseWriter, r *http.Request) {
	var consumerLocation models.ConsumerLocation

	if err := decoder.BindJson(r, &consumerLocation); err != nil {
		a.log.Println(err)
		if err := handlers.Respond(rw, "incorrect input data", http.StatusBadRequest); err != nil {
			a.log.Println(err)
			a.SignalShutdown()
		}
		return
	}

	data, statusCode := a.consumerService.UpdateConsumerLocation(consumerLocation)
	if err := handlers.Respond(rw, data, statusCode); err != nil {
		a.log.Println(err)
		a.SignalShutdown()
		return
	}
}

// SignalShutdown is used to gracefully shut down the app when an integrity
// issue is identified.
func (a *ConsumerHandler) SignalShutdown() {
	a.shutdown <- syscall.SIGTERM
}
