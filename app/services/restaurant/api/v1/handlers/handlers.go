package handlers

import (
	"io"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
	v1 "github.com/nndergunov/deliveryApp/app/pkg/api/v1"
	"github.com/nndergunov/deliveryApp/app/pkg/logger"
	"github.com/nndergunov/deliveryApp/app/services/restaurant/api/v1/communication"
	"github.com/nndergunov/deliveryApp/app/services/restaurant/pkg/service"
)

const (
	restaurantIDKey = "restaurantID"
	menuIDKey       = "menuID"
)

type endpointHandler struct {
	service  service.AppService
	serveMux *mux.Router
	log      *logger.Logger
}

// NewEndpointHandler returns new http multiplexer with configured endpoints.
func NewEndpointHandler(serviceInstance service.AppService, log *logger.Logger) *mux.Router {
	serveMux := mux.NewRouter()

	handler := endpointHandler{
		service:  serviceInstance,
		serveMux: serveMux,
		log:      log,
	}

	handler.handlerInit()

	return handler.serveMux
}

func (e *endpointHandler) handlerInit() {
	e.serveMux.HandleFunc("/status", e.statusHandler)

	e.serveMux.HandleFunc("/v1/restaurants", e.returnRestaurantList).Methods(http.MethodGet)
	e.serveMux.HandleFunc("/v1/restaurants/{"+restaurantIDKey+"}", e.returnRestaurant).Methods(http.MethodGet)
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
		e.handleError(err, responseWriter)

		return
	}

	_, err = io.WriteString(responseWriter, string(status))
	if err != nil {
		e.handleError(err, responseWriter)

		return
	}

	responseWriter.WriteHeader(http.StatusOK)

	e.log.Printf("gave status %s", data.IsUp)
}

func (e endpointHandler) returnRestaurantList(w http.ResponseWriter, _ *http.Request) {
	// swagger:operation GET /restaurants returnRestaurantList
	//
	// Returns restaurant list
	//
	// ---
	// produces:
	// - application/json
	// responses:
	//   '200':
	//     description: restaurant list
	restaurants, err := e.service.ReturnAllRestaurants()
	if err != nil {
		e.handleError(err, w)

		return
	}

	response := restaurantListToResponse(restaurants)

	err = v1.Respond(response, w)
	if err != nil {
		e.log.Println(err)
	}
}

func (e endpointHandler) returnRestaurant(w http.ResponseWriter, r *http.Request) {
	// swagger:operation GET /restaurants/{id} returnRestaurant
	//
	// Returns requested restaurant restaurant
	//
	// ---
	// produces:
	// - application/json
	// responses:
	//   '200':
	//     description: requested restaurtant
	restaurantID, err := getIDFromEndpoint(restaurantIDKey, r)
	if err != nil {
		e.handleError(err, w)

		return
	}

	restaurant, err := e.service.ReturnRestaurant(restaurantID)
	if err != nil {
		e.handleError(err, w)

		return
	}

	response := restaurantToResponse(*restaurant)

	err = v1.Respond(response, w)
	if err != nil {
		e.log.Println(err)
	}
}

func (e endpointHandler) returnMenu(w http.ResponseWriter, r *http.Request) {
	// swagger:operation GET /restaurants/{id}/menu returnMenu
	//
	// Returns menu of the requested restaurant
	//
	// ---
	// produces:
	// - application/json
	// responses:
	//   '200':
	//     description: menu of the requested restaurtant
	restaurantID, err := getIDFromEndpoint(restaurantIDKey, r)
	if err != nil {
		e.handleError(err, w)

		return
	}

	menu, err := e.service.ReturnMenu(restaurantID)
	if err != nil {
		e.handleError(err, w)

		return
	}

	response := menuToResponse(*menu)

	err = v1.Respond(response, w)
	if err != nil {
		e.log.Println(err)
	}
}

func (e *endpointHandler) createRestaurant(w http.ResponseWriter, r *http.Request) {
	// swagger:operation POST /restaurants createRestaurant
	//
	// Returns menu of the requested restaurant
	//
	// ---
	// produces:
	// - application/json
	// parameters:
	//   - name: Body
	//     in: body
	//     description: restaurant data
	//     required: true
	// responses:
	//   '200':
	//     description: created restaurant
	req, err := ioutil.ReadAll(r.Body)
	if err != nil {
		e.handleError(err, w)

		return
	}

	restaurantData := new(communication.RestaurantData)

	err = v1.Decode(req, restaurantData)
	if err != nil {
		e.handleError(err, w)

		return
	}

	rest := requestToRestaurant(0, restaurantData)

	createdRest, err := e.service.CreateNewRestaurant(rest)
	if err != nil {
		e.handleError(err, w)

		return
	}

	response := restaurantToResponse(*createdRest)

	err = v1.Respond(response, w)
	if err != nil {
		e.log.Println(err)
	}
}

func (e *endpointHandler) updateRestaurant(w http.ResponseWriter, r *http.Request) {
	// swagger:operation PUT /restaurants/{id} updateRestaurant
	//
	// Updates restaurant data
	//
	// ---
	// produces:
	// - application/json
	// parameters:
	//   - name: Body
	//     in: body
	//     description: updated restaurant data
	//     required: true
	// responses:
	//   '200':
	//     description: updated restaurant data
	restaurantID, err := getIDFromEndpoint(restaurantIDKey, r)
	if err != nil {
		e.handleError(err, w)

		return
	}

	req, err := ioutil.ReadAll(r.Body)
	if err != nil {
		e.handleError(err, w)

		return
	}

	restaurantData := new(communication.RestaurantData)

	err = v1.Decode(req, restaurantData)
	if err != nil {
		e.handleError(err, w)

		return
	}

	rest := requestToRestaurant(restaurantID, restaurantData)

	updatedRestaurant, err := e.service.UpdateRestaurant(rest)
	if err != nil {
		e.handleError(err, w)

		return
	}

	response := restaurantToResponse(*updatedRestaurant)

	err = v1.Respond(response, w)
	if err != nil {
		e.log.Println(err)
	}
}

func (e *endpointHandler) deleteRestaurant(w http.ResponseWriter, r *http.Request) {
	// swagger:operation DELETE /restaurants/{id} deleteRestaurant
	//
	// Deletes restaurant data
	//
	// ---
	// produces:
	// - application/json
	// responses:
	//   '200':
	restaurantID, err := getIDFromEndpoint(restaurantIDKey, r)
	if err != nil {
		e.handleError(err, w)

		return
	}

	err = e.service.DeleteRestaurant(restaurantID)
	if err != nil {
		e.handleError(err, w)

		return
	}

	err = v1.Respond(nil, w)
	if err != nil {
		e.log.Println(err)
	}
}

func (e *endpointHandler) createMenu(w http.ResponseWriter, r *http.Request) {
	// swagger:operation POST /restaurants/{id}/menu createMenu
	//
	// Creates menu in the restaurant
	//
	// ---
	// produces:
	// - application/json
	// parameters:
	//   - name: Body
	//     in: body
	//     description: menu data
	//     required: true
	// responses:
	//   '200':
	//     description: created menu
	restaurantID, err := getIDFromEndpoint(restaurantIDKey, r)
	if err != nil {
		e.handleError(err, w)

		return
	}

	req, err := ioutil.ReadAll(r.Body)
	if err != nil {
		e.handleError(err, w)

		return
	}

	menuData := new(communication.MenuData)

	err = v1.Decode(req, menuData)
	if err != nil {
		e.handleError(err, w)

		return
	}

	menu := requestToMenu(restaurantID, menuData)

	createdMenu, err := e.service.CreateMenu(menu)
	if err != nil {
		e.handleError(err, w)

		return
	}

	response := menuToResponse(*createdMenu)

	err = v1.Respond(response, w)
	if err != nil {
		e.log.Println(err)
	}
}

func (e *endpointHandler) addMenuItem(w http.ResponseWriter, r *http.Request) {
	// swagger:operation PUT /restaurants/{id}/menu addMenuItem
	//
	// Adds new menu item in the restaurant
	//
	// ---
	// produces:
	// - application/json
	// parameters:
	//   - name: Body
	//     in: body
	//     description: menu item data
	//     required: true
	// responses:
	//   '200':
	//     description: created menu item
	restaurantID, err := getIDFromEndpoint(restaurantIDKey, r)
	if err != nil {
		e.handleError(err, w)

		return
	}

	req, err := ioutil.ReadAll(r.Body)
	if err != nil {
		e.handleError(err, w)

		return
	}

	menuItemData := new(communication.MenuItemData)

	err = v1.Decode(req, menuItemData)
	if err != nil {
		e.handleError(err, w)

		return
	}

	menuItem := requestToMenuItem(0, menuItemData)

	addedMenuItem, err := e.service.AddMenuItem(restaurantID, menuItem)
	if err != nil {
		e.handleError(err, w)

		return
	}

	response := menuItemToResponse(*addedMenuItem)

	err = v1.Respond(response, w)
	if err != nil {
		e.log.Println(err)
	}
}

func (e *endpointHandler) updateMenuItem(w http.ResponseWriter, r *http.Request) {
	// swagger:operation PATCH /restaurants/{id}/menu/{itemid} updateMenuItem
	//
	// Updates menu item in the restaurant
	//
	// ---
	// produces:
	// - application/json
	// parameters:
	//   - name: Body
	//     in: body
	//     description: updated menu item data
	//     required: true
	// responses:
	//   '200':
	//     description: updated menu item
	restaurantID, err := getIDFromEndpoint(restaurantIDKey, r)
	if err != nil {
		e.handleError(err, w)

		return
	}

	menuItemID, err := getIDFromEndpoint(menuIDKey, r)
	if err != nil {
		e.handleError(err, w)

		return
	}

	req, err := ioutil.ReadAll(r.Body)
	if err != nil {
		e.handleError(err, w)

		return
	}

	menuItemData := new(communication.MenuItemData)

	err = v1.Decode(req, menuItemData)
	if err != nil {
		e.handleError(err, w)

		return
	}

	menuItem := requestToMenuItem(menuItemID, menuItemData)

	updatedMenuItem, err := e.service.UpdateMenuItem(restaurantID, menuItem)
	if err != nil {
		e.handleError(err, w)

		return
	}

	response := menuItemToResponse(*updatedMenuItem)

	err = v1.Respond(response, w)
	if err != nil {
		e.log.Println(err)
	}
}

func (e *endpointHandler) deleteMenuItem(w http.ResponseWriter, r *http.Request) {
	// swagger:operation DELETE /restaurants/{id}/menu/{itemid} deleteMenuItem
	//
	// Deletes menu item in the restaurant
	//
	// ---
	// produces:
	// - application/json
	// responses:
	//   '200':
	restaurantID, err := getIDFromEndpoint(restaurantIDKey, r)
	if err != nil {
		e.handleError(err, w)

		return
	}

	menuItemID, err := getIDFromEndpoint(menuIDKey, r)
	if err != nil {
		e.handleError(err, w)

		return
	}

	err = e.service.DeleteMenuItem(restaurantID, menuItemID)
	if err != nil {
		e.handleError(err, w)

		return
	}

	w.WriteHeader(http.StatusOK)
}
