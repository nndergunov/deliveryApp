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
	"github.com/nndergunov/deliveryApp/app/services/restaurant/pkg/service"
)

const (
	restaurantIDKey = "restaurantID"
	menuIDKey       = "menuID"
)

type endpointHandler struct {
	app      *service.Service
	serveMux *mux.Router
	log      *logger.Logger
}

// NewEndpointHandler returns new http multiplexer with configured endpoints.
func NewEndpointHandler(appInstance *service.Service, log *logger.Logger) *mux.Router {
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
	e.serveMux.HandleFunc("/v1/admin/restaurants/{"+restaurantIDKey+"}",
		e.updateRestaurant).Methods(http.MethodPut)
	e.serveMux.HandleFunc("/v1/admin/restaurants/{"+restaurantIDKey+"}",
		e.deleteRestaurant).Methods(http.MethodDelete)

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

	createdRest, err := e.app.CreateNewRestaurant(rest)
	if err != nil {
		e.log.Println(err)

		w.WriteHeader(http.StatusInternalServerError)

		return
	}

	resp := restaurantToResponse(*createdRest)

	e.respond(resp, w)
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

	updatedRestaurant, err := e.app.UpdateRestaurant(rest)
	if err != nil {
		e.log.Println(err)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)

			return
		}
	}

	resp := restaurantToResponse(*updatedRestaurant)

	e.respond(resp, w)
}

func (e *endpointHandler) deleteRestaurant(w http.ResponseWriter, r *http.Request) {
	restaurantID, err := getIDFromEndpoint(restaurantIDKey, r)
	if err != nil {
		e.log.Println(err)

		w.WriteHeader(http.StatusBadRequest)

		return
	}

	err = e.app.DeleteRestaurant(restaurantID)
	if err != nil {
		e.log.Println(err)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)

			return
		}
	}

	e.respond(nil, w)
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

	createdMenu, err := e.app.CreateMenu(menu)
	if err != nil {
		e.log.Println(err)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)

			return
		}
	}

	resp := menuToResponse(*createdMenu)

	e.respond(resp, w)
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

	addedMenuItem, err := e.app.AddMenuItem(restaurantID, menuItem)
	if err != nil {
		e.log.Println(err)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)

			return
		}
	}

	e.respond(addedMenuItem, w)
}

func (e *endpointHandler) updateMenuItem(w http.ResponseWriter, r *http.Request) {
	restaurantID, err := getIDFromEndpoint(restaurantIDKey, r)
	if err != nil {
		e.log.Println(err)

		w.WriteHeader(http.StatusBadRequest)

		return
	}

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

	updatedMenuItem, err := e.app.UpdateMenuItem(restaurantID, menuItem)
	if err != nil {
		e.log.Println(err)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)

			return
		}
	}

	resp := menuItemToResponse(*updatedMenuItem)

	e.respond(resp, w)
}

func (e *endpointHandler) deleteMenuItem(w http.ResponseWriter, r *http.Request) {
	restaurantID, err := getIDFromEndpoint(restaurantIDKey, r)
	if err != nil {
		e.log.Println(err)

		w.WriteHeader(http.StatusBadRequest)

		return
	}

	menuItemID, err := getIDFromEndpoint(menuIDKey, r)
	if err != nil {
		e.log.Println(err)

		w.WriteHeader(http.StatusBadRequest)

		return
	}

	err = e.app.DeleteMenuItem(restaurantID, menuItemID)
	if err != nil {
		e.log.Println(err)

		if errors.Is(err, service.ErrItemIsNotInMenu) {
			w.WriteHeader(http.StatusBadRequest)

			return
		}

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)

			return
		}
	}

	w.WriteHeader(http.StatusOK)
}
