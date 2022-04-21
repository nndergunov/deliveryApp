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
	"github.com/nndergunov/deliveryApp/app/services/restaurant/pkg/app"
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

	e.serveMux.HandleFunc("/v1/restaurants/{"+restaurantIDKey+"}/menu", e.returnMenu).Methods(http.MethodGet)

	e.serveMux.HandleFunc("/v1/admin/restaurants", e.createRestaurant).Methods(http.MethodPost)
	e.serveMux.HandleFunc("/v1/admin/restaurants", e.updateRestaurant).Methods(http.MethodPut)

	e.serveMux.HandleFunc(
		"/v1/admin/restaurants/{"+restaurantIDKey+"}/menu", e.createMenu).Methods(http.MethodPost)
	e.serveMux.HandleFunc(
		"/v1/admin/restaurants/{"+restaurantIDKey+"}/menu", e.addMenuItem).Methods(http.MethodPut)

	e.serveMux.HandleFunc(
		"/v1/admin/restaurants/{"+restaurantIDKey+"}/menu/{"+menuIDKey+"}",
		e.updateMenuItem).Methods(http.MethodPatch)
	e.serveMux.HandleFunc(
		"/v1/admin/restaurants/{"+restaurantIDKey+"}/menu/{"+menuIDKey+"}",
		e.deleteMenuItem).Methods(http.MethodDelete)
}

func (e endpointHandler) statusHandler(responseWriter http.ResponseWriter, _ *http.Request) {
	data := v1.Status{
		ServiceName: "restaurant",
		IsUp:        "up",
	}

	status, err := v1.EncodeIndent(data, "", " ")
	if err != nil {
		e.log.Println(err)

		responseWriter.WriteHeader(http.StatusInternalServerError)

		return
	}

	_, err = io.WriteString(responseWriter, string(status))
	if err != nil {
		e.log.Println(err)

		responseWriter.WriteHeader(http.StatusInternalServerError)

		return
	}

	responseWriter.WriteHeader(http.StatusOK)

	e.log.Printf("gave status %s", data.IsUp)
}

func (e endpointHandler) returnRestaurantList(w http.ResponseWriter, _ *http.Request) {
	restaurants, err := e.app.ReturnAllRestaurants()
	if err != nil {
		e.log.Println(err)

		w.WriteHeader(http.StatusInternalServerError)

		return
	}

	response := restaurantListToResponse(restaurants)

	e.respond(response, w)
}

func (e endpointHandler) returnMenu(w http.ResponseWriter, r *http.Request) {
	restaurantID, err := getIDFromEndpoint(restaurantIDKey, r)
	if err != nil {
		e.log.Println(err)

		w.WriteHeader(http.StatusInternalServerError)

		return
	}

	menu, err := e.app.ReturnMenu(restaurantID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)

		return
	}

	response := menuToResponse(*menu)

	e.respond(response, w)
}

func (e *endpointHandler) createRestaurant(w http.ResponseWriter, r *http.Request) {
	req, err := ioutil.ReadAll(r.Body)
	if err != nil {
		e.log.Println(err)

		w.WriteHeader(http.StatusBadRequest)

		return
	}

	restaurantData, err := restaurantapi.DecodeRestaurantData(req)
	if err != nil {
		e.log.Println(err)

		w.WriteHeader(http.StatusBadRequest)

		return
	}

	rest := requestToRestaurant(0, restaurantData)

	err = e.app.CreateNewRestaurant(rest)
	if err != nil {
		e.log.Println(err)

		w.WriteHeader(http.StatusInternalServerError)

		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (e *endpointHandler) updateRestaurant(w http.ResponseWriter, r *http.Request) {
	restaurantID, err := getIDFromEndpoint(restaurantIDKey, r)
	if err != nil {
		e.log.Println(err)

		w.WriteHeader(http.StatusBadRequest)

		return
	}

	req, err := ioutil.ReadAll(r.Body)
	if err != nil {
		e.log.Println(err)

		w.WriteHeader(http.StatusBadRequest)

		return
	}

	restaurantData, err := restaurantapi.DecodeRestaurantData(req)
	if err != nil {
		e.log.Println(err)

		w.WriteHeader(http.StatusBadRequest)

		return
	}

	rest := requestToRestaurant(restaurantID, restaurantData)

	err = e.app.UpdateRestaurant(rest)
	if err != nil {
		e.log.Println(err)

		if errors.Is(err, app.ErrIsNotInDatabase) {
			w.WriteHeader(http.StatusBadRequest)

			return
		} else {
			w.WriteHeader(http.StatusInternalServerError)

			return
		}
	}

	w.WriteHeader(http.StatusOK)
}

func (e *endpointHandler) createMenu(w http.ResponseWriter, r *http.Request) {
	restaurantID, err := getIDFromEndpoint(restaurantIDKey, r)
	if err != nil {
		e.log.Println(err)

		w.WriteHeader(http.StatusBadRequest)

		return
	}

	req, err := ioutil.ReadAll(r.Body)
	if err != nil {
		e.log.Println(err)

		w.WriteHeader(http.StatusBadRequest)

		return
	}

	menuData, err := restaurantapi.DecodeMenuData(req)
	if err != nil {
		e.log.Println(err)

		w.WriteHeader(http.StatusBadRequest)

		return
	}

	menu := requestToMenu(restaurantID, menuData)

	err = e.app.CreateMenu(menu)
	if err != nil {
		e.log.Println(err)

		if errors.Is(err, app.ErrIsNotInDatabase) {
			w.WriteHeader(http.StatusBadRequest)

			return
		} else {
			w.WriteHeader(http.StatusInternalServerError)

			return
		}
	}

	w.WriteHeader(http.StatusCreated)
}

func (e *endpointHandler) addMenuItem(w http.ResponseWriter, r *http.Request) {
	restaurantID, err := getIDFromEndpoint(restaurantIDKey, r)
	if err != nil {
		e.log.Println(err)

		w.WriteHeader(http.StatusBadRequest)

		return
	}

	req, err := ioutil.ReadAll(r.Body)
	if err != nil {
		e.log.Println(err)

		w.WriteHeader(http.StatusBadRequest)

		return
	}

	menuItemData, err := restaurantapi.DecodeMenuItem(req)
	if err != nil {
		e.log.Println(err)

		w.WriteHeader(http.StatusBadRequest)

		return
	}

	menuItem := requestToMenuItem(0, menuItemData)

	err = e.app.AddMenuItem(restaurantID, menuItem)
	if err != nil {
		e.log.Println(err)

		if errors.Is(err, app.ErrIsNotInDatabase) {
			w.WriteHeader(http.StatusBadRequest)

			return
		} else {
			w.WriteHeader(http.StatusInternalServerError)

			return
		}
	}

	w.WriteHeader(http.StatusCreated)
}

func (e *endpointHandler) updateMenuItem(w http.ResponseWriter, r *http.Request) {
	menuItemID, err := getIDFromEndpoint(menuIDKey, r)
	if err != nil {
		e.log.Println(err)

		w.WriteHeader(http.StatusBadRequest)

		return
	}

	req, err := ioutil.ReadAll(r.Body)
	if err != nil {
		e.log.Println(err)

		w.WriteHeader(http.StatusBadRequest)

		return
	}

	menuItemData, err := restaurantapi.DecodeMenuItem(req)
	if err != nil {
		e.log.Println(err)

		w.WriteHeader(http.StatusBadRequest)

		return
	}

	menuItem := requestToMenuItem(menuItemID, menuItemData)

	err = e.app.UpdateMenuItem(menuItem)
	if err != nil {
		e.log.Println(err)

		if errors.Is(err, app.ErrIsNotInDatabase) {
			w.WriteHeader(http.StatusBadRequest)

			return
		} else {
			w.WriteHeader(http.StatusInternalServerError)

			return
		}
	}

	w.WriteHeader(http.StatusOK)
}

func (e *endpointHandler) deleteMenuItem(w http.ResponseWriter, r *http.Request) {
	menuItemID, err := getIDFromEndpoint(menuIDKey, r)
	if err != nil {
		e.log.Println(err)

		w.WriteHeader(http.StatusBadRequest)

		return
	}

	err = e.app.DeleteMenuItem(menuItemID)
	if err != nil {
		e.log.Println(err)

		if errors.Is(err, app.ErrIsNotInDatabase) {
			w.WriteHeader(http.StatusBadRequest)

			return
		} else {
			w.WriteHeader(http.StatusInternalServerError)

			return
		}
	}

	w.WriteHeader(http.StatusOK)
}
