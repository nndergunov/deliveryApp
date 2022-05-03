// Package handler contains a small handler framework extension.
package handler

import (
	"courier/models"
	"courier/service"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/nndergunov/deliveryApp/app/pkg/logger"
	"io"
	"net/http"
	"os"
	"strconv"
	"syscall"
)

type Params struct {
	Logger         *logger.Logger
	CourierService service.CourierService
	Srv            *mux.Router
	Shutdown       chan os.Signal
}

// App is the entrypoint into our application
type App struct {
	srv            *mux.Router
	log            *logger.Logger
	courierService service.CourierService
	shutdown       chan os.Signal
}

// NewCarrierHandler creates an App value that handle a set of routes for the application.
func NewCarrierHandler(p Params) *App {
	app := &App{
		log:            p.Logger,
		srv:            p.Srv,
		courierService: p.CourierService,
		shutdown:       p.Shutdown,
	}
	const version = "/v1"
	const courier = "/courier"

	app.srv.HandleFunc(version+courier+"/new", app.new).Methods(http.MethodPost)
	app.srv.HandleFunc(version+courier+"/remove/{id}", app.remove).Methods(http.MethodPost)
	app.srv.HandleFunc(version+courier+"/update", app.update).Methods(http.MethodPut)
	app.srv.HandleFunc(version+courier+"/get-all", app.getAll).Methods(http.MethodGet)
	app.srv.HandleFunc(version+courier+"/get/{id}", app.get).Methods(http.MethodGet)
	app.srv.HandleFunc(version+courier+"/status", app.status).Methods(http.MethodGet)

	return app
}

// ServeHTTP implements the http.Handler interface.
func (a *App) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	a.srv.ServeHTTP(w, r)
}

func (a *App) new(rw http.ResponseWriter, r *http.Request) {

	var courier models.Courier

	if err := BindJson(r, &courier); err != nil {
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

	returnedCourier, err := a.courierService.Get(0, courier.Username)
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

	newCourier, err := a.courierService.Insert(courier)
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

func (a *App) remove(rw http.ResponseWriter, r *http.Request) {
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

	if err = a.courierService.Remove(idUint); err != nil {
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

func (a *App) update(rw http.ResponseWriter, r *http.Request) {
	var courier models.Courier

	if err := BindJson(r, &courier); err != nil {
		if err := Respond(rw, "incorrect input data", http.StatusBadRequest); err != nil {
			a.log.Println(err)
			a.SignalShutdown()
		}
		a.log.Println(err)
		return
	}

	updatedCourier, err := a.courierService.Update(courier)
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

func (a *App) getAll(rw http.ResponseWriter, r *http.Request) {

	allCourier, err := a.courierService.GetAll()
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
func (a *App) get(rw http.ResponseWriter, r *http.Request) {
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

	courier, err := a.courierService.Get(idUint, "")
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

func (a *App) status(rw http.ResponseWriter, r *http.Request) {
	rw.WriteHeader(http.StatusOK)
}

func BindJson(req *http.Request, obj interface{}) error {
	if req == nil || req.Body == nil {
		return fmt.Errorf("invalid request")
	}
	return decodeJSON(req.Body, obj)
}

func decodeJSON(r io.Reader, obj interface{}) error {
	decoder := json.NewDecoder(r)
	decoder.DisallowUnknownFields()

	if err := decoder.Decode(obj); err != nil {
		return err
	}
	return nil
}

// SignalShutdown is used to gracefully shut down the app when an integrity
// issue is identified.
func (a *App) SignalShutdown() {
	a.shutdown <- syscall.SIGTERM
}
