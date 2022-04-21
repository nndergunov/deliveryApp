package handlers

import (
	"errors"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
	v1 "github.com/nndergunov/deliveryApp/app/pkg/api/v1"
	"github.com/nndergunov/deliveryApp/app/pkg/api/v1/restaurantapi"
	"github.com/nndergunov/deliveryApp/app/pkg/logger"
	"github.com/nndergunov/deliveryApp/app/services/kitchen/pkg/app"
)

const (
	restaurantIDKey = "restaurantID"
	menuIDKey       = "menuID"
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

	e.serveMux.HandleFunc("/v1/restaurants/{"+restaurantIDKey+"}/menu", e.returnMenu).Methods(http.MethodGet)
	e.serveMux.HandleFunc("/v1/restaurants/{"+restaurantIDKey+"}/menu", e.createMenu).Methods(http.MethodPost)
	e.serveMux.HandleFunc("/v1/restaurants/{"+restaurantIDKey+"}/menu", e.addMenuItem).Methods(http.MethodPut)

	e.serveMux.HandleFunc(
		"/v1/restaurants/{"+restaurantIDKey+"}/menu/{"+menuIDKey+"}", e.updateMenuItem).Methods(http.MethodPatch)
	e.serveMux.HandleFunc(
		"/v1/restaurants/{"+restaurantIDKey+"}/menu/{"+menuIDKey+"}", e.updateMenuItem).Methods(http.MethodDelete)
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

func (e *endpointHandler) createRestaurant(w http.ResponseWriter, r *http.Request) {
	req, err := ioutil.ReadAll(r.Body)
	if err != nil {
		e.log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
	}

	restaurantData, err := restaurantapi.DecodeRestaurantData(req)
	if err != nil {
		e.log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
	}

	rest := requestToRestaurant(0, restaurantData)

	err = e.app.CreateNewRestaurant(rest)
	if err != nil {
		e.log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func (e *endpointHandler) updateRestaurant(w http.ResponseWriter, r *http.Request) {
	restaurantID, err := getIDFromEndpoint(restaurantIDKey, r)
	if err != nil {
		e.log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
	}

	req, err := ioutil.ReadAll(r.Body)
	if err != nil {
		e.log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
	}

	restaurantData, err := restaurantapi.DecodeRestaurantData(req)
	if err != nil {
		e.log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
	}

	rest := requestToRestaurant(restaurantID, restaurantData)

	err = e.app.UpdateRestaurant(rest)
	if err != nil {
		e.log.Println(err)

		if errors.Is(err, app.ErrIsNotInMap) {
			w.WriteHeader(http.StatusBadRequest)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}
	}
}

func (e endpointHandler) returnMenu(w http.ResponseWriter, r *http.Request) {
	// TODO logic.
}

func (e *endpointHandler) createMenu(w http.ResponseWriter, r *http.Request) {
	restaurantID, err := getIDFromEndpoint(restaurantIDKey, r)
	if err != nil {
		e.log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
	}

	req, err := ioutil.ReadAll(r.Body)
	if err != nil {
		e.log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
	}

	menuData, err := restaurantapi.DecodeMenuData(req)
	if err != nil {
		e.log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
	}

	menu := requestToMenu(restaurantID, menuData)

	err = e.app.CreateMenu(menu)
	if err != nil {
		e.log.Println(err)

		if errors.Is(err, app.ErrIsNotInMap) {
			w.WriteHeader(http.StatusBadRequest)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}
	}
}

func (e *endpointHandler) addMenuItem(w http.ResponseWriter, r *http.Request) {
	restaurantID, err := getIDFromEndpoint(restaurantIDKey, r)
	if err != nil {
		e.log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
	}

	req, err := ioutil.ReadAll(r.Body)
	if err != nil {
		e.log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
	}

	menuItemData, err := restaurantapi.DecodeMenuItem(req)
	if err != nil {
		e.log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
	}

	menuItem := requestToMenuItem(0, menuItemData)

	err = e.app.AddMenuItem(restaurantID, menuItem)
	if err != nil {
		e.log.Println(err)

		if errors.Is(err, app.ErrIsNotInMap) {
			w.WriteHeader(http.StatusBadRequest)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}
	}
}

func (e *endpointHandler) updateMenuItem(w http.ResponseWriter, r *http.Request) {
	restaurantID, err := getIDFromEndpoint(restaurantIDKey, r)
	if err != nil {
		e.log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
	}

	menuItemID, err := getIDFromEndpoint(menuIDKey, r)
	if err != nil {
		e.log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
	}

	req, err := ioutil.ReadAll(r.Body)
	if err != nil {
		e.log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
	}

	menuItemData, err := restaurantapi.DecodeMenuItem(req)
	if err != nil {
		e.log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
	}

	menuItem := requestToMenuItem(menuItemID, menuItemData)

	err = e.app.UpdateMenuItem(restaurantID, menuItem)
	if err != nil {
		e.log.Println(err)

		if errors.Is(err, app.ErrIsNotInMap) {
			w.WriteHeader(http.StatusBadRequest)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}
	}
}

func (e *endpointHandler) deleteMenuItem(w http.ResponseWriter, r *http.Request) {
	restaurantID, err := getIDFromEndpoint(restaurantIDKey, r)
	if err != nil {
		e.log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
	}

	menuItemID, err := getIDFromEndpoint(menuIDKey, r)
	if err != nil {
		e.log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
	}

	err = e.app.DeleteMenuItem(restaurantID, menuItemID)
	if err != nil {
		e.log.Println(err)

		if errors.Is(err, app.ErrIsNotInMap) {
			w.WriteHeader(http.StatusBadRequest)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}
	}
}
