package deliveryhandler

import (
	"delivery/api/v1/deliveryapi"
	"delivery/pkg/service/deliveryservice"
	"github.com/gorilla/mux"
	v1 "github.com/nndergunov/deliveryApp/app/pkg/api/v1"
	"github.com/nndergunov/deliveryApp/app/pkg/logger"
	"io"
	"net/http"
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

	c.serveMux.HandleFunc(version+"/delivery/time", c.getDeliveryTime).Methods(http.MethodGet)
	c.serveMux.HandleFunc(version+"/delivery/cost", c.getDeliveryCost).Methods(http.MethodGet)

	c.serveMux.HandleFunc(version+"/delivery/assign/courier/{"+orderIDKey+"}", c.assignCourierToOrder).Methods(http.MethodPost)
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

func (c *deliveryHandler) getDeliveryTime(rw http.ResponseWriter, r *http.Request) {
	var deliveryTimeRequest deliveryapi.DeliveryTimeRequest

	if err := deliveryapi.BindJson(r, &deliveryTimeRequest); err != nil {
		c.log.Println(err)
		if err := deliveryapi.Respond(rw, http.StatusBadRequest, errIncorrectInputData.Error()); err != nil {
			c.log.Println(err)
		}
		return
	}

	deliveryLocation := requestToDeliveryLocation(&deliveryTimeRequest)

	data, err := c.deliveryService.GetDeliveryTime(deliveryLocation)

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

	response := deliveryTimeToResponse(data)

	if err := deliveryapi.Respond(rw, http.StatusOK, response); err != nil {
		c.log.Println(err)
		return
	}

}

func (c *deliveryHandler) getDeliveryCost(rw http.ResponseWriter, r *http.Request) {
	var deliveryCostRequest deliveryapi.DeliveryCostRequest

	if err := deliveryapi.BindJson(r, &deliveryCostRequest); err != nil {
		c.log.Println(err)
		if err := deliveryapi.Respond(rw, http.StatusBadRequest, errIncorrectInputData.Error()); err != nil {
			c.log.Println(err)
		}
		return
	}

	deliveryLocation := requestToDeliveryCostLocation(&deliveryCostRequest)

	data, err := c.deliveryService.GetDeliveryCost(deliveryLocation)

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

	response := deliveryCostToResponse(data)

	if err := deliveryapi.Respond(rw, http.StatusOK, response); err != nil {
		c.log.Println(err)
		return
	}

}

func (c *deliveryHandler) assignCourierToOrder(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	orderID, ok := vars[orderIDKey]
	if !ok {
		if err := deliveryapi.Respond(rw, http.StatusBadRequest, errNoOrderIDParam); err != nil {
			c.log.Println(err)
		}
	}

	var assignOrderToCourierRequest deliveryapi.OrderRequest

	if err := deliveryapi.BindJson(r, &assignOrderToCourierRequest); err != nil {
		c.log.Println(err)
		if err := deliveryapi.Respond(rw, http.StatusBadRequest, errIncorrectInputData.Error()); err != nil {
			c.log.Println(err)
		}
		return
	}

	assignOrderToCourier := requestToOrder(&assignOrderToCourierRequest)

	data, err := c.deliveryService.AssignCourierToOrder(orderID, assignOrderToCourier)

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

	response := deliveryAssignedCourierResponse(data)

	if err := deliveryapi.Respond(rw, http.StatusOK, response); err != nil {
		c.log.Println(err)
		return
	}

}
