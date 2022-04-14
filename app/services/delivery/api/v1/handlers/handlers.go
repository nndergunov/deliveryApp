package handlers

import (
	"io"
	"net/http"

	"github.com/gorilla/mux"
	v1 "github.com/nndergunov/deliveryApp/app/pkg/api/v1"
	"github.com/nndergunov/deliveryApp/app/pkg/logger"
	"github.com/nndergunov/deliveryApp/app/services/delivery/pkg/app"
)

type endpointHandler struct {
	serveMux *mux.Router
	log      *logger.Logger
}

// NewEndpointHandler returns new http multiplexer with configured endpoints.
func NewEndpointHandler(log *logger.Logger) *mux.Router {
	serveMux := mux.NewRouter()

	handler := endpointHandler{
		serveMux: serveMux,
		log:      log,
	}

	handler.handlerInit()

	return handler.serveMux
}

func (e *endpointHandler) handlerInit() {
	e.serveMux.HandleFunc("/status", e.statusHandler)
	e.serveMux.HandleFunc("/v1/couriers/courier-for-delivery", e.findCourier).Methods(http.MethodGet)
	e.serveMux.HandleFunc("/v1/deliveries/cost", e.returnDeliveryCost).Methods(http.MethodGet)
}

func (e endpointHandler) statusHandler(responseWriter http.ResponseWriter, _ *http.Request) {
	data := v1.Status{
		ServiceName: "delivery",
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

func (e endpointHandler) findCourier(w http.ResponseWriter, r *http.Request) {
	// TODO logic.
}

func (e endpointHandler) returnDeliveryCost(w http.ResponseWriter, r *http.Request) {
	data, err := v1.Encode(app.CalculateDeliveryPrice())
	if err != nil {
		e.log.Println(err)
	}

	_, err = w.Write(data)
	if err != nil {
		e.log.Println(err)
	}
}
