// Package handler contains a small handler framework extension.
package handler

import (
	"courier/pkg/decoder"
	"database/sql"
	"net/http"
	"os"
	"strconv"
	"syscall"

	"courier/models"
	"courier/service"

	"github.com/gorilla/mux"
	"github.com/nndergunov/deliveryApp/app/pkg/logger"
)

type Params struct {
	Logger         *logger.Logger
	CourierService service.CourierService
	Route          *mux.Router
	Shutdown       chan os.Signal
}

// CourierHandler is the entrypoint into our application
type CourierHandler struct {
	rout           *mux.Router
	log            *logger.Logger
	courierService service.CourierService
	shutdown       chan os.Signal
}

// Handler is the interface for the carrier handler.
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
	p.Route.HandleFunc(version+courier+"/remove/{id}", handler.removeCourier).Methods(http.MethodPost)
	p.Route.HandleFunc(version+courier+"/update", handler.updateCourier).Methods(http.MethodPut)
	p.Route.HandleFunc(version+courier+"/get-all", handler.getAllCourier).Methods(http.MethodGet)
	p.Route.HandleFunc(version+courier+"/get/{id}", handler.getCourier).Methods(http.MethodGet)

	return handler
}

func (a *CourierHandler) insertNewCourier(rw http.ResponseWriter, r *http.Request) {
	var courier models.Courier

	if err := decoder.BindJson(r, &courier); err != nil {
		if err := Respond(rw, "incorrect input data", http.StatusBadRequest); err != nil {
			a.log.Println(err)
			a.SignalShutdown()
			return
		}
		a.log.Println(err)
		return
	}
	if len(courier.Username) < 4 || len(courier.Password) < 8 {
		if err := Respond(rw, "username or password don't meet requirement", http.StatusBadRequest); err != nil {
			a.log.Println(err)
			a.SignalShutdown()
			return
		}
		return
	}

	returnedCourier, err := a.courierService.GetCourier(0, courier.Username)
	if err != nil && err != sql.ErrNoRows {
		if err := Respond(rw, "", http.StatusInternalServerError); err != nil {
			a.log.Println(err)
			a.SignalShutdown()
			return
		}
		a.log.Println(err)
		return
	}

	if returnedCourier != nil {
		if err := Respond(rw, "courier with this username already exist", http.StatusBadRequest); err != nil {
			a.log.Println(err)
			a.SignalShutdown()
			return
		}
		return
	}

	newCourier, err := a.courierService.InsertCourier(courier)
	if err != nil {
		if err := Respond(rw, err, http.StatusBadRequest); err != nil {
			a.log.Println(err)
			a.SignalShutdown()
			return
		}
		a.log.Println("can't create courier", err)
		return
	}

	if err := Respond(rw, newCourier, http.StatusOK); err != nil {
		a.log.Println(err)
		a.SignalShutdown()
	}
}

func (a *CourierHandler) removeCourier(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		if err := Respond(rw, "id is missing in parameters", http.StatusInternalServerError); err != nil {
			a.log.Println(err)
			a.SignalShutdown()
			return
		}
		a.log.Println("id is missing in parameters")
		return
	}

	idUint, err := strconv.ParseUint(string(id), 10, 64)
	if err != nil {

		if err := Respond(rw, "wrong id type", http.StatusInternalServerError); err != nil {
			a.log.Println(err)
			a.SignalShutdown()
		}
		a.log.Println(err)
		return
	}

	if err = a.courierService.RemoveCourier(idUint); err != nil {
		if err := Respond(rw, err, http.StatusInternalServerError); err != nil {
			a.log.Println(err)
			a.SignalShutdown()
			return
		}
	}

	if err := Respond(rw, "Courier removed", http.StatusOK); err != nil {
		a.log.Println(err)
		a.SignalShutdown()
		return
	}
}

func (a *CourierHandler) updateCourier(rw http.ResponseWriter, r *http.Request) {
	var courier models.Courier

	if err := decoder.BindJson(r, &courier); err != nil {
		if err := Respond(rw, "incorrect input data", http.StatusBadRequest); err != nil {
			a.log.Println(err)
			a.SignalShutdown()
		}
		a.log.Println(err)
		return
	}

	updatedCourier, err := a.courierService.UpdateCourier(courier)
	if err != nil && err == sql.ErrNoRows {
		if err := Respond(rw, "this courier doesn't exist", http.StatusOK); err != nil {
			a.SignalShutdown()
			a.log.Println(err)
			return
		}
		return
	}

	if err != nil && err != sql.ErrNoRows {
		if err := Respond(rw, "", http.StatusInternalServerError); err != nil {
			a.SignalShutdown()
			a.log.Println(err)
			return
		}
		return
	}

	if err := Respond(rw, updatedCourier, http.StatusOK); err != nil {
		a.log.Println(err)
		return
	}
}

func (a *CourierHandler) getAllCourier(rw http.ResponseWriter, r *http.Request) {
	allCourier, err := a.courierService.GetAllCourier()
	if err != nil {
		if err := Respond(rw, err, http.StatusInternalServerError); err != nil {
			a.log.Println(err)
			return
		}
		a.log.Println(err)
		return
	}

	if err := Respond(rw, allCourier, http.StatusOK); err != nil {
		a.log.Println(err)
		return
	}
}

func (a *CourierHandler) getCourier(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		if err := Respond(rw, "id is missing in parameters", http.StatusInternalServerError); err != nil {
			a.log.Println(err)
			a.SignalShutdown()
			return
		}
		a.log.Println("id is missing in parameters")
		return
	}

	idUint, err := strconv.ParseUint(string(id), 10, 64)
	if err != nil {
		if err := Respond(rw, "wrong id type", http.StatusInternalServerError); err != nil {
			a.log.Println(err)
			a.SignalShutdown()
			return
		}
		a.log.Println(err)
		return
	}

	courier, err := a.courierService.GetCourier(idUint, "")
	if err != nil && err == sql.ErrNoRows {
		if err := Respond(rw, "no courier found", http.StatusOK); err != nil {
			a.log.Println(err)
			a.SignalShutdown()
			return
		}
		if err := Respond(rw, "", http.StatusInternalServerError); err != nil {
			a.log.Println(err)
			a.SignalShutdown()
			return
		}

		a.log.Println(err)
		return
	}

	if err := Respond(rw, courier, http.StatusOK); err != nil {
		a.log.Println(err)
		a.SignalShutdown()
		return
	}
}

// SignalShutdown is used to gracefully shut down the app when an integrity
// issue is identified.
func (a *CourierHandler) SignalShutdown() {
	a.shutdown <- syscall.SIGTERM
}
