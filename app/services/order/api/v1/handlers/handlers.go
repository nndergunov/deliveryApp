package handlers

import (
	"errors"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
	v1 "github.com/nndergunov/deliveryApp/app/pkg/api/v1"
	"github.com/nndergunov/deliveryApp/app/pkg/api/v1/orderapi"
	"github.com/nndergunov/deliveryApp/app/pkg/logger"
	"github.com/nndergunov/deliveryApp/app/services/order/pkg/domain"
	"github.com/nndergunov/deliveryApp/app/services/order/pkg/service"
)

const orderID = "orderID"

type endpointHandler struct {
	serviceInstance service.App
	serveMux        *mux.Router
	log             *logger.Logger
}

// NewEndpointHandler returns new http multiplexer with configured endpoints.
func NewEndpointHandler(serviceInstance service.App, log *logger.Logger) *mux.Router {
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
	parameters := domain.SearchParameters{
		FromRestaurantID: nil,
		Statuses:         nil,
		ExcludeStatuses:  nil,
	}

	req, err := ioutil.ReadAll(request.Body)
	if err != nil {
		e.log.Println(err)

		responseWriter.WriteHeader(http.StatusBadRequest)

		return
	}

	searchParams := new(orderapi.OrderFilters)

	if len(req) != 0 {
		err = v1.Decode(req, searchParams)
		if err != nil {
			e.log.Println(err)

			responseWriter.WriteHeader(http.StatusBadRequest)

			return
		}

		parameters = parseParameters(*searchParams)
	}

	orders, err := e.serviceInstance.ReturnOrderList(parameters)
	if err != nil {
		e.log.Println(err)

		responseWriter.WriteHeader(http.StatusInternalServerError)

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
	req, err := ioutil.ReadAll(request.Body)
	if err != nil {
		e.log.Println(err)

		responseWriter.WriteHeader(http.StatusBadRequest)

		return
	}

	orderData := new(orderapi.OrderData)

	err = v1.Decode(req, orderData)
	if err != nil {
		e.log.Println(err)

		responseWriter.WriteHeader(http.StatusBadRequest)

		return
	}

	order := requestToOrder(*orderData)

	createdOrder, err := e.serviceInstance.CreateOrder(order, order.UserAccount)
	if err != nil {
		if errors.Is(err, service.ErrRestaurantOffline) {
			err := v1.RespondWithError("restaurant is not accepting orders", http.StatusBadRequest, responseWriter)
			if err != nil {
				e.log.Println(err)

				responseWriter.WriteHeader(http.StatusInternalServerError)
			}

			return
		}

		if errors.Is(err, service.ErrLowBalance) {
			err := v1.RespondWithError("user has not enough balance to create this order", http.StatusBadRequest, responseWriter)
			if err != nil {
				e.log.Println(err)

				responseWriter.WriteHeader(http.StatusInternalServerError)
			}

			return
		}

		e.log.Println(err)

		responseWriter.WriteHeader(http.StatusInternalServerError)

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
	returnOrderID, err := getIDFromEndpoint(orderID, request)
	if err != nil {
		e.log.Println(err)

		responseWriter.WriteHeader(http.StatusBadRequest)

		return
	}

	order, err := e.serviceInstance.ReturnOrder(returnOrderID)
	if err != nil {
		e.log.Println(err)

		responseWriter.WriteHeader(http.StatusInternalServerError)

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
	updateOrderID, err := getIDFromEndpoint(orderID, request)
	if err != nil {
		e.log.Println(err)

		responseWriter.WriteHeader(http.StatusBadRequest)

		return
	}

	req, err := ioutil.ReadAll(request.Body)
	if err != nil {
		e.log.Println(err)

		responseWriter.WriteHeader(http.StatusBadRequest)

		return
	}

	orderData := new(orderapi.OrderData)

	err = v1.Decode(req, orderData)
	if err != nil {
		e.log.Println(err)

		responseWriter.WriteHeader(http.StatusBadRequest)

		return
	}

	order := requestToOrder(*orderData)

	order.OrderID = updateOrderID

	updatedOrder, err := e.serviceInstance.UpdateOrder(order)
	if err != nil {
		e.log.Println(err)

		responseWriter.WriteHeader(http.StatusInternalServerError)

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
	updateOrderID, err := getIDFromEndpoint(orderID, request)
	if err != nil {
		e.log.Println(err)

		responseWriter.WriteHeader(http.StatusBadRequest)

		return
	}

	req, err := ioutil.ReadAll(request.Body)
	if err != nil {
		e.log.Println(err)

		responseWriter.WriteHeader(http.StatusBadRequest)

		return
	}

	statusData := new(orderapi.OrderStatusData)

	err = v1.Decode(req, statusData)
	if err != nil {
		e.log.Println(err)

		responseWriter.WriteHeader(http.StatusBadRequest)

		return
	}

	status := requestToStatus(*statusData)

	status.OrderID = updateOrderID

	_, err = e.serviceInstance.UpdateStatus(status)
	if err != nil {
		e.log.Println(err)

		responseWriter.WriteHeader(http.StatusInternalServerError)

		return
	}

	responseWriter.WriteHeader(http.StatusOK)
}
