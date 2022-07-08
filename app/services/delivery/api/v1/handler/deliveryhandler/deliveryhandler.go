package deliveryhandler

import (
	"errors"
	"io"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/nndergunov/deliveryApp/app/pkg/api/v1"
	"github.com/nndergunov/deliveryApp/app/pkg/logger"

	"github.com/nndergunov/deliveryApp/app/services/delivery/api/v1/deliveryapi"

	"github.com/nndergunov/deliveryApp/app/services/delivery/pkg/service/deliveryservice"
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

	c.serveMux.HandleFunc(version+"/estimate", c.getEstimateDeliveryValues).Methods(http.MethodGet)
	c.serveMux.HandleFunc(version+"/orders/{"+orderIDKey+"}/assign", c.assignOrder).Methods(http.MethodPost)
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

func (c *deliveryHandler) getEstimateDeliveryValues(rw http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()

	consumerIDList := queryParams["consumer_id"]
	restaurantIDList := queryParams["restaurant_id"]

	if consumerIDList == nil || restaurantIDList == nil {
		if err := deliveryapi.Respond(rw, http.StatusBadRequest, "consumer_id or restaurant_id param not found"); err != nil {
			c.log.Println(err)
			return
		}
	}

	consumerID := consumerIDList[0]
	if consumerID == "" {
		if err := deliveryapi.Respond(rw, http.StatusBadRequest, "wrong consumer_id"); err != nil {
			c.log.Println(err)
		}
		return
	}

	restaurantID := consumerIDList[0]
	if restaurantID == "" {
		if err := deliveryapi.Respond(rw, http.StatusBadRequest, "wrong restaurant_id"); err != nil {
			c.log.Println(err)
		}
		return
	}

	data, err := c.deliveryService.GetEstimateDelivery(consumerID, restaurantID)
	if err != nil {

		if errors.Is(err, systemErr) {
			if err := deliveryapi.Respond(rw, http.StatusInternalServerError, ""); err != nil {
				c.log.Println(err)
			}
			return
		}

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
	if err != nil {

		if errors.Is(err, systemErr) {
			if err := deliveryapi.Respond(rw, http.StatusInternalServerError, ""); err != nil {
				c.log.Println(err)
			}
			return
		}

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
