package handlers

import (
	"io"
	"net/http"

	"github.com/gorilla/mux"
	v1 "github.com/nndergunov/deliveryApp/app/pkg/api/v1"
	"github.com/nndergunov/deliveryApp/app/pkg/logger"
	"github.com/nndergunov/deliveryApp/app/services/order/pkg/app"
)

type endpointHandler struct {
	appInstance *app.App
	serveMux    *mux.Router
	log         *logger.Logger
}

// NewEndpointHandler returns new http multiplexer with configured endpoints.
func NewEndpointHandler(appInstance *app.App, log *logger.Logger) *mux.Router {
	serveMux := mux.NewRouter()

	handler := endpointHandler{
		appInstance: appInstance,
		serveMux:    serveMux,
		log:         log,
	}

	handler.handlerInit()

	return handler.serveMux
}

func (e *endpointHandler) handlerInit() {
	e.serveMux.HandleFunc("/status", e.statusHandler)
	e.serveMux.HandleFunc("/v1/orders", e.createOrder).Methods(http.MethodPost)
	e.serveMux.HandleFunc("/v1/orders/{id}/status", e.returnOrderStatus).Methods(http.MethodGet)
	e.serveMux.HandleFunc("/v1/orders/{id}/status", e.updateOrderList).Methods(http.MethodPut)
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

func (e endpointHandler) createOrder(w http.ResponseWriter, r *http.Request) {
	// TODO logic.
}

func (e endpointHandler) returnOrderStatus(w http.ResponseWriter, r *http.Request) {
	// TODO logic.
}

func (e endpointHandler) updateOrderList(w http.ResponseWriter, r *http.Request) {
	// TODO logic.
}
