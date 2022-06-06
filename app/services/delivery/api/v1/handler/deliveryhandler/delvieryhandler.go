package deliveryhandler

import (
	"github.com/gorilla/mux"

	v1 "github.com/nndergunov/deliveryApp/app/pkg/api/v1"
	"github.com/nndergunov/deliveryApp/app/pkg/logger"

	"io"
	"net/http"

	"delivery/api/v1/deliveryapi"
	"delivery/pkg/service/deliveryservice"
)

type Params struct {
	Logger          *logger.Logger
	DeliveryService deliveryservice.DeliveryService
}

// deliveryHandler is the entrypoint into our application
type deliveryHandler struct {
	serveMux        *mux.Router
	log             *logger.Logger
	deliveryService deliveryservice.DeliveryService
}

// NewDeliveryHandler returns new http multiplexer with configured endpoints.
func NewDeliveryHandler(p Params) *mux.Router {
	serveMux := mux.NewRouter()

	handler := deliveryHandler{
		serveMux:        serveMux,
		log:             p.Logger,
		deliveryService: p.DeliveryService,
	}

	handler.handlerInit()

	return handler.serveMux
}

const orderIDKey = "order_id"

// NewDeliveryHandler creates an deliveryHandler value that handle a set of routes for the application.
func (c *deliveryHandler) handlerInit() {

	version := "/v1"
	c.serveMux.HandleFunc("/status", c.statusHandler).Methods(http.MethodPost)

	c.serveMux.HandleFunc(version+"/estimate-delivery", c.getEstimateDelivery).Methods(http.MethodGet)
	c.serveMux.HandleFunc(version+"/order/{"+orderIDKey+"}/assign", c.assignOrder).Methods(http.MethodPost)
}

func (c *deliveryHandler) statusHandler(responseWriter http.ResponseWriter, _ *http.Request) {
	data := v1.Status{
		ServiceName: "delivery",
		IsUp:        "up",
	}

	status, err := v1.EncodeIndent(data, "", " ")
	if err != nil {
		c.log.Println(err)

		responseWriter.WriteHeader(http.StatusInternalServerError)

		return
	}

	_, err = io.WriteString(responseWriter, string(status))
	if err != nil {
		c.log.Println(err)

		responseWriter.WriteHeader(http.StatusInternalServerError)

		return
	}

	responseWriter.WriteHeader(http.StatusOK)

	c.log.Printf("gave status %s", data.IsUp)
}

func (c *deliveryHandler) getEstimateDelivery(rw http.ResponseWriter, r *http.Request) {
	var estimateDeliveryRequest deliveryapi.EstimateDeliveryRequest

	if err := deliveryapi.BindJson(r, &estimateDeliveryRequest); err != nil {
		c.log.Println(err)
		if err := deliveryapi.Respond(rw, http.StatusBadRequest, errIncorrectInputData.Error()); err != nil {
			c.log.Println(err)
		}
		return
	}

	estimateDelivery := requestToEstimateDelivery(&estimateDeliveryRequest)

	data, err := c.deliveryService.GetEstimateDelivery(estimateDelivery)

	if err != nil && err == systemErr {
		c.log.Println(err)
		if err := deliveryapi.Respond(rw, http.StatusInternalServerError, ""); err != nil {
			c.log.Println(err)
		}
		return
	}

	if err != nil {
		if err := deliveryapi.Respond(rw, http.StatusBadRequest, err.Error()); err != nil {
			c.log.Println(err)
		}
		return
	}

	response := estimateDeliveryToResponse(data)

	if err := deliveryapi.Respond(rw, http.StatusOK, response); err != nil {
		c.log.Println(err)
		return
	}

}

func (c *deliveryHandler) assignOrder(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	orderID, ok := vars[orderIDKey]
	if !ok {
		if err := deliveryapi.Respond(rw, http.StatusBadRequest, errNoOrderIDParam); err != nil {
			c.log.Println(err)
		}
	}

	var assignOrderRequest deliveryapi.AssignOrderRequest

	if err := deliveryapi.BindJson(r, &assignOrderRequest); err != nil {
		c.log.Println(err)
		if err := deliveryapi.Respond(rw, http.StatusBadRequest, errIncorrectInputData.Error()); err != nil {
			c.log.Println(err)
		}
		return
	}

	order := requestToOrder(&assignOrderRequest)

	data, err := c.deliveryService.AssignOrder(orderID, order)

	if err != nil && err == systemErr {
		c.log.Println(err)
		if err := deliveryapi.Respond(rw, http.StatusInternalServerError, ""); err != nil {
			c.log.Println(err)
		}
		return
	}

	if err != nil {
		if err := deliveryapi.Respond(rw, http.StatusBadRequest, err.Error()); err != nil {
			c.log.Println(err)
		}
		return
	}

	response := assignOrderResponse(data)

	if err := deliveryapi.Respond(rw, http.StatusOK, response); err != nil {
		c.log.Println(err)
		return
	}

}