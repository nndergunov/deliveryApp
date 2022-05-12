// Package handlers contains a small handlers framework extension.
package courierhandler

import (
	"courier/internal/handlers"
	"fmt"
	"net/http"
	"os"
	"syscall"

	"courier/internal/models"
	"courier/internal/services"
	"courier/pkg/decoder"

	"github.com/gorilla/mux"
	"github.com/nndergunov/deliveryApp/app/pkg/logger"
)

type Params struct {
	Logger         *logger.Logger
	CourierService services.CourierService
	Route          *mux.Router
	Shutdown       chan os.Signal
}

// CourierHandler is the entrypoint into our application
type CourierHandler struct {
	rout           *mux.Router
	log            *logger.Logger
	courierService services.CourierService
	shutdown       chan os.Signal
}

// Handler is the interface for the carrier handlers.
type Handler interface {
	insertNewCourier(rw http.ResponseWriter, r *http.Request)
	removeCourier(rw http.ResponseWriter, r *http.Request)
	updateCourier(rw http.ResponseWriter, r *http.Request)
	getAllCourier(rw http.ResponseWriter, r *http.Request)
	getCourier(rw http.ResponseWriter, r *http.Request)
}

// NewCourierHandler creates an CourierHandler value that handle a set of routes for the application.
func NewCourierHandler(p Params) Handler {
	handler := &CourierHandler{
		log:            p.Logger,
		rout:           p.Route,
		courierService: p.CourierService,
		shutdown:       p.Shutdown,
	}
	const version = "/v1"
	const courier = "/courier"

	p.Route.HandleFunc(version+courier+"/new", handler.insertNewCourier).Methods(http.MethodPost)
	p.Route.HandleFunc(version+courier+"/remove", handler.removeCourier).Methods(http.MethodPost)
	p.Route.HandleFunc(version+courier+"/update", handler.updateCourier).Methods(http.MethodPut)
	p.Route.HandleFunc(version+courier+"/update-available", handler.updateCourierAvailable).Methods(http.MethodPut)
	p.Route.HandleFunc(version+courier+"/get-all", handler.getAllCourier).Methods(http.MethodGet)
	p.Route.HandleFunc(version+courier+"/get", handler.getCourier).Methods(http.MethodGet)

	return handler
}

func (a *CourierHandler) insertNewCourier(rw http.ResponseWriter, r *http.Request) {
	var courier models.NewCourierRequest

	if err := decoder.BindJson(r, &courier); err != nil {
		a.log.Println(err)
		if err := handlers.Respond(rw, nil, fmt.Errorf("incorrect input data")); err != nil {
			a.log.Println(err)
		}
		return
	}

	data, err := a.courierService.InsertNewCourier(courier)
	if err := handlers.Respond(rw, data, err); err != nil {
		a.log.Println(err)
		return
	}
}

func (a *CourierHandler) removeCourier(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		if err := handlers.Respond(rw, nil, fmt.Errorf("id is not indicated")); err != nil {
			a.log.Println(err)
		}
		return
	}

	data, err := a.courierService.RemoveCourier(id)
	if err := handlers.Respond(rw, data, err); err != nil {
		a.log.Println(err)
		return
	}
}

func (a *CourierHandler) updateCourier(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		if err := handlers.Respond(rw, nil, fmt.Errorf("id param is not found")); err != nil {
			a.log.Println(err)
		}
		return
	}

	var courier models.UpdateCourierRequest

	if err := decoder.BindJson(r, &courier); err != nil {
		a.log.Println(err)
		if err := handlers.Respond(rw, nil, fmt.Errorf("incorrect input data")); err != nil {
			a.log.Println(err)
		}
		return
	}

	data, err := a.courierService.UpdateCourier(courier, id)
	if err := handlers.Respond(rw, data, err); err != nil {
		a.log.Println(err)
		return
	}
}

func (a *CourierHandler) updateCourierAvailable(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		if err := handlers.Respond(rw, nil, fmt.Errorf("id param is not found")); err != nil {
			a.log.Println(err)
		}
		return
	}

	available, ok := vars["available"]
	if !ok {
		if err := handlers.Respond(rw, nil, fmt.Errorf("available param is not found")); err != nil {
			a.log.Println(err)
		}
		return
	}

	data, err := a.courierService.UpdateCourierAvailable(id, available)
	if err := handlers.Respond(rw, data, err); err != nil {
		a.log.Println(err)
		return
	}
}

func (a *CourierHandler) getAllCourier(rw http.ResponseWriter, r *http.Request) {
	data, err := a.courierService.GetAllCourier()
	if err := handlers.Respond(rw, data, err); err != nil {
		a.log.Println(err)
		return
	}
}

func (a *CourierHandler) getCourier(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		if err := handlers.Respond(rw, nil, fmt.Errorf("id param is not found")); err != nil {
			a.log.Println(err)
		}
		return
	}

	data, err := a.courierService.GetCourier(id)
	if err := handlers.Respond(rw, data, err); err != nil {
		a.log.Println(err)
		return
	}
}

// SignalShutdown is used to gracefully shut down the app when an integrity
// issue is identified.
func (a *CourierHandler) SignalShutdown() {
	a.shutdown <- syscall.SIGTERM
}
