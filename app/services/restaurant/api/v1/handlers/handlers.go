package handlers

import (
	"io"
	"net/http"

	"github.com/gorilla/mux"
	v1 "github.com/nndergunov/deliveryApp/app/pkg/api/v1"
	"github.com/nndergunov/deliveryApp/app/pkg/logger"
	"github.com/nndergunov/deliveryApp/app/services/kitchen/pkg/app"
)

type endpointHandler struct {
	app      *app.App
	serveMux *mux.Router
	log      *logger.Logger
}

// NewEndpointHandler returns new http multiplexer with configured endpoints.
func NewEndpointHandler(appInstance *app.App, log *logger.Logger) *mux.Router {
	serveMux := mux.NewRouter()

	handler := endpointHandler{
		app:      appInstance,
		serveMux: serveMux,
		log:      log,
	}

	handler.handlerInit()

	return handler.serveMux
}

func (e *endpointHandler) handlerInit() {
	e.serveMux.HandleFunc("/status", e.statusHandler)
	e.serveMux.HandleFunc("/v1/restaurants", e.returnRestaurantList).Methods(http.MethodGet)
	e.serveMux.HandleFunc("/v1/restaurants", e.createRestaurant).Methods(http.MethodPost)
	e.serveMux.HandleFunc("/v1/restaurants", e.updateRestaurant).Methods(http.MethodPut)
	e.serveMux.HandleFunc("/v1/restaurants/{id}/menu", e.returnMenu).Methods(http.MethodGet)
	e.serveMux.HandleFunc("/v1/restaurants/{id}/menu", e.createMenu).Methods(http.MethodPost)
	e.serveMux.HandleFunc("/v1/restaurants/{id}/menu", e.updateMenu).Methods(http.MethodPut)
}

func (e endpointHandler) statusHandler(responseWriter http.ResponseWriter, _ *http.Request) {
	data := v1.Status{
		ServiceName: "restaurant",
		IsUp:        "up",
	}

	status, err := v1.EncodeIndent(data, "", " ")
	if err != nil {
		e.log.Println(err)
	}

	_, err = io.WriteString(responseWriter, string(status))
	if err != nil {
		e.log.Printf("status write: %v", err)

		return
	}

	e.log.Printf("gave status %s", data.IsUp)
}

func (e endpointHandler) returnRestaurantList(w http.ResponseWriter, r *http.Request) {
	// TODO logic.
}

func (e endpointHandler) createRestaurant(w http.ResponseWriter, r *http.Request) {
	// TODO logic.
}

func (e endpointHandler) updateRestaurant(w http.ResponseWriter, r *http.Request) {
	// TODO logic.
}

func (e endpointHandler) returnMenu(w http.ResponseWriter, r *http.Request) {
	// TODO logic.
}

func (e endpointHandler) createMenu(w http.ResponseWriter, r *http.Request) {
	// TODO logic.
}

func (e endpointHandler) updateMenu(w http.ResponseWriter, r *http.Request) {
	// TODO logic.
}

func (e endpointHandler) returnIncompleteOrderList(w http.ResponseWriter, r *http.Request) {
	// TODO logic.
}
