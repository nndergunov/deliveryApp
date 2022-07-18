// Package handlers contains endpoint api layer of the application.
package handlers

import (
	"io"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
	v1 "github.com/nndergunov/deliveryApp/app/pkg/api/v1"
	"github.com/nndergunov/deliveryApp/app/pkg/logger"
	"github.com/nndergunov/deliveryApp/app/services/order/api/v1/orderapi"
	"github.com/nndergunov/deliveryApp/app/services/order/pkg/domain"
	"github.com/nndergunov/deliveryApp/app/services/order/pkg/service"
)

const orderID = "orderID"

type endpointHandler struct {
	serviceInstance service.AppService
	serveMux        *mux.Router
	log             *logger.Logger
}

// NewEndpointHandler returns new http multiplexer with configured endpoints.
func NewEndpointHandler(serviceInstance service.AppService, log *logger.Logger) *mux.Router {
	serveMux := mux.NewRouter()

	handler := endpointHandler{
		serviceInstance: serviceInstance,
		serveMux:        serveMux,
		log:             log,
	}

	handler.handlerInit()

	return handler.serveMux
}

func (e *endpointHandler) handlerInit() {
	e.serveMux.HandleFunc("/status", e.statusHandler)

	e.serveMux.HandleFunc("/v1/orders", e.returnAllOrders).Methods(http.MethodGet)
	e.serveMux.HandleFunc("/v1/orders", e.createOrder).Methods(http.MethodPost)
	e.serveMux.HandleFunc("/v1/orders/{"+orderID+"}", e.returnOrder).Methods(http.MethodGet)
	e.serveMux.HandleFunc("/v1/orders/{"+orderID+"}", e.updateOrder).Methods(http.MethodPut)

	e.serveMux.HandleFunc("/v1/admin/orders/{"+orderID+"}/status", e.updateOrderStatus).Methods(http.MethodPut)
}

func (e endpointHandler) statusHandler(responseWriter http.ResponseWriter, _ *http.Request) {
	data := v1.Status{
		ServiceName: "order",
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

func (e endpointHandler) returnAllOrders(responseWriter http.ResponseWriter, request *http.Request) {
	// swagger:operation GET /orders returnAllOrders
	//
	// Returns all orders from the order service
	//
	// ---
	// produces:
	// - application/json
	// parameters:
	// - name: Body
	//   in: body
	//   description: order search filters
	//   schema:
	//     $ref: "#/definitions/OrderFilters"
	//   required: false
	// responses:
	//   '200':
	//     description: order list response
	//     schema:
	//       $ref: "#/definitions/ReturnOrderList"
	parameters := domain.SearchParameters{
		FromRestaurantID: nil,
		Statuses:         nil,
		ExcludeStatuses:  nil,
	}

	req, err := ioutil.ReadAll(request.Body)
	if err != nil {
		e.handleError(err, responseWriter)

		return
	}

	searchParams := new(orderapi.OrderFilters)

	if len(req) != 0 {
		err = v1.Decode(req, searchParams)
		if err != nil {
			e.handleError(err, responseWriter)

			return
		}

		parameters = parseParameters(*searchParams)
	} // if request body is empty we do not need to parse searchParams as it would result in err

	orders, err := e.serviceInstance.ReturnOrderList(parameters)
	if err != nil {
		e.handleError(err, responseWriter)

		return
	}

	response := orderListToResponse(orders)

	err = v1.Respond(response, responseWriter)
	if err != nil {
		e.log.Println(err)

		return
	}
}

func (e endpointHandler) createOrder(responseWriter http.ResponseWriter, request *http.Request) {
	// swagger:operation POST /orders createOrder
	//
	// Creates new order in the order service
	//
	// ---
	// produces:
	// - application/json
	// parameters:
	// - name: Body
	//   in: body
	//   description: order data
	//   schema:
	//     $ref: "#/definitions/OrderData"
	//   required: true
	// responses:
	//   '200':
	//     description: created order data
	//     schema:
	//       $ref: "#/definitions/ReturnOrder"
	req, err := ioutil.ReadAll(request.Body)
	if err != nil {
		e.handleError(err, responseWriter)

		return
	}

	orderData := new(orderapi.OrderData)

	err = v1.Decode(req, orderData)
	if err != nil {
		e.handleError(err, responseWriter)

		return
	}

	order := requestToOrder(*orderData)

	createdOrder, err := e.serviceInstance.CreateOrder(order, order.FromUserID)
	if err != nil {
		e.handleError(err, responseWriter)

		return
	}

	response := orderToResponse(*createdOrder)

	err = v1.Respond(response, responseWriter)
	if err != nil {
		e.log.Println(err)

		return
	}
}

func (e endpointHandler) returnOrder(responseWriter http.ResponseWriter, request *http.Request) {
	// swagger:operation GET /orders/{id} returnOrder
	//
	// Returns specified order data
	//
	// ---
	// produces:
	// - application/json
	// responses:
	//   '200':
	//     description: requested order data
	//     schema:
	//       $ref: "#/definitions/ReturnOrder"
	returnOrderID, err := getIDFromEndpoint(orderID, request)
	if err != nil {
		e.handleError(err, responseWriter)

		return
	}

	order, err := e.serviceInstance.ReturnOrder(returnOrderID)
	if err != nil {
		e.handleError(err, responseWriter)

		return
	}

	response := orderToResponse(*order)

	err = v1.Respond(response, responseWriter)
	if err != nil {
		e.log.Println(err)

		return
	}
}

func (e endpointHandler) updateOrder(responseWriter http.ResponseWriter, request *http.Request) {
	// swagger:operation PUT /orders/{id} updateOrder
	//
	// Updates data of the specified order
	//
	// ---
	// produces:
	// - application/json
	// parameters:
	// - name: Body
	//   in: body
	//   description: order data
	//   schema:
	//     $ref: "#/definitions/OrderData"
	//   required: true
	// responses:
	//   '200':
	//     description: updated order data
	//     schema:
	//       $ref: "#/definitions/ReturnOrder"
	updateOrderID, err := getIDFromEndpoint(orderID, request)
	if err != nil {
		e.handleError(err, responseWriter)

		return
	}

	req, err := ioutil.ReadAll(request.Body)
	if err != nil {
		e.handleError(err, responseWriter)

		return
	}

	orderData := new(orderapi.OrderData)

	err = v1.Decode(req, orderData)
	if err != nil {
		e.handleError(err, responseWriter)

		return
	}

	order := requestToOrder(*orderData)

	order.OrderID = updateOrderID

	updatedOrder, err := e.serviceInstance.UpdateOrder(order)
	if err != nil {
		e.handleError(err, responseWriter)

		return
	}

	response := orderToResponse(*updatedOrder)

	err = v1.Respond(response, responseWriter)
	if err != nil {
		e.log.Println(err)

		return
	}
}

func (e *endpointHandler) updateOrderStatus(responseWriter http.ResponseWriter, request *http.Request) {
	// swagger:operation PUT /admin/orders/{id}/status updateOrderStatus
	//
	// Updates status of the specified order
	//
	// ---
	// produces:
	// - application/json
	// parameters:
	// - name: Body
	//   in: body
	//   description: order status data
	//   schema:
	//     $ref: "#/definitions/OrderStatusData"
	//   required: true
	// responses:
	//   '200':
	updateOrderID, err := getIDFromEndpoint(orderID, request)
	if err != nil {
		e.handleError(err, responseWriter)

		return
	}

	req, err := ioutil.ReadAll(request.Body)
	if err != nil {
		e.handleError(err, responseWriter)

		return
	}

	statusData := new(orderapi.OrderStatusData)

	err = v1.Decode(req, statusData)
	if err != nil {
		e.handleError(err, responseWriter)

		return
	}

	status := requestToStatus(*statusData)

	status.OrderID = updateOrderID

	_, err = e.serviceInstance.UpdateStatus(status)
	if err != nil {
		e.handleError(err, responseWriter)

		return
	}

	responseWriter.WriteHeader(http.StatusOK)
}
