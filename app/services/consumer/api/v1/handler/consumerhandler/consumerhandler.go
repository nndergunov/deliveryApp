// Package handler contains a small handler framework extension.
package consumerhandler

import (
	"consumer/api/v1/consumerapi"
	"consumer/service"
	"github.com/gorilla/mux"
	"github.com/nndergunov/deliveryApp/app/pkg/logger"
	"net/http"
)

type Params struct {
	Logger          *logger.Logger
	ConsumerService service.ConsumerService
}

// consumerHandler is the entrypoint into our application
type consumerHandler struct {
	serveMux        *mux.Router
	log             *logger.Logger
	consumerService service.ConsumerService
}

// NewConsumerHandler returns new http multiplexer with configured endpoints.
func NewConsumerHandler(p Params) *mux.Router {
	serveMux := mux.NewRouter()

	handler := consumerHandler{
		serveMux:        serveMux,
		log:             p.Logger,
		consumerService: p.ConsumerService,
	}

	handler.handlerInit()

	return handler.serveMux
}

// NewConsumerHandler creates an consumerHandler value that handle a set of routes for the application.
func (c *consumerHandler) handlerInit() {

	const version = "/v1"
	const consumer = "/consumer"
	const consumerLocation = "/consumer-location"

	c.serveMux.HandleFunc(version+consumer+"/new", c.insertNewConsumer).Methods(http.MethodPost)
	c.serveMux.HandleFunc(version+consumer+"/delete/{id}", c.deleteConsumer).Methods(http.MethodPost)
	c.serveMux.HandleFunc(version+consumer+"/update/{id}", c.updateConsumer).Methods(http.MethodPut)
	c.serveMux.HandleFunc(version+consumer+"/get-all", c.getAllConsumer).Methods(http.MethodGet)
	c.serveMux.HandleFunc(version+consumer+"/get/{id}", c.getConsumer).Methods(http.MethodGet)

	c.serveMux.HandleFunc(version+consumer+"/update/{id}", c.updateConsumerLocation).Methods(http.MethodPut)
}

func (c *consumerHandler) insertNewConsumer(rw http.ResponseWriter, r *http.Request) {
	var consumerRequest consumerapi.NewConsumerRequest

	if err := consumerapi.BindJson(r, &consumerRequest); err != nil {
		c.log.Println(err)
		if err := consumerapi.Respond(rw, http.StatusBadRequest, "incorrect input data"); err != nil {
			c.log.Println(err)
		}
		return
	}

	consumer := requestToNewConsumer(&consumerRequest)

	data, err := c.consumerService.InsertConsumer(consumer)
	if err != nil {
		if err := consumerapi.Respond(rw, http.StatusBadRequest, err); err != nil {
			c.log.Println(err)
		}
		return
	}

	if data == nil {
		if err := consumerapi.Respond(rw, http.StatusInternalServerError, ""); err != nil {
			c.log.Println(err)
		}
		return
	}

	response := consumerToResponse(*data)

	if err := consumerapi.Respond(rw, http.StatusOK, response); err != nil {
		c.log.Println(err)
		return
	}

}

func (c *consumerHandler) deleteConsumer(rw http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		if err := consumerapi.Respond(rw, http.StatusBadRequest, "id param is not found"); err != nil {
			c.log.Println(err)
		}
	}

	data, err := c.consumerService.DeleteConsumer(id)
	if err != nil {
		if err := consumerapi.Respond(rw, http.StatusBadRequest, err); err != nil {
			c.log.Println(err)
		}
		return
	}

	if err := consumerapi.Respond(rw, http.StatusOK, data); err != nil {
		c.log.Println(err)
		return
	}
}

func (c *consumerHandler) updateConsumer(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		if err := consumerapi.Respond(rw, http.StatusBadRequest, "id param is not found"); err != nil {
			c.log.Println(err)
		}
	}

	var updateCourierRequest consumerapi.UpdateConsumerRequest

	if err := consumerapi.BindJson(r, &updateCourierRequest); err != nil {
		c.log.Println(err)
		if err := consumerapi.Respond(rw, http.StatusBadRequest, "incorrect input data"); err != nil {
			c.log.Println(err)
		}
		return
	}

	consumer := requestToUpdateConsumer(&updateCourierRequest)

	data, err := c.consumerService.UpdateConsumer(consumer, id)

	if err != nil {
		if err := consumerapi.Respond(rw, http.StatusBadRequest, err); err != nil {
			c.log.Println(err)
		}
		return
	}

	if data == nil {
		if err := consumerapi.Respond(rw, http.StatusInternalServerError, ""); err != nil {
			c.log.Println(err)
		}
		return
	}

	response := consumerToResponse(*data)

	if err := consumerapi.Respond(rw, http.StatusOK, response); err != nil {
		c.log.Println(err)
		return
	}
}

func (c *consumerHandler) getAllConsumer(rw http.ResponseWriter, r *http.Request) {
	data, err := c.consumerService.GetAllConsumer()

	if err != nil {
		if err := consumerapi.Respond(rw, http.StatusBadRequest, err); err != nil {
			c.log.Println(err)
		}
		return
	}

	if data == nil {
		if err := consumerapi.Respond(rw, http.StatusInternalServerError, ""); err != nil {
			c.log.Println(err)
		}
		return
	}

	response := consumerListToResponse(data)

	if err := consumerapi.Respond(rw, http.StatusOK, response); err != nil {
		c.log.Println(err)
		return
	}
}

func (c *consumerHandler) getConsumer(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		if err := consumerapi.Respond(rw, http.StatusBadRequest, "id param is not found"); err != nil {
			c.log.Println(err)
		}
	}

	data, err := c.consumerService.GetConsumer(id)

	if err != nil {
		if err := consumerapi.Respond(rw, http.StatusBadRequest, err); err != nil {
			c.log.Println(err)
		}
		return
	}

	if data == nil {
		if err := consumerapi.Respond(rw, http.StatusInternalServerError, ""); err != nil {
			c.log.Println(err)
		}
		return
	}

	response := consumerToResponse(*data)

	if err := consumerapi.Respond(rw, http.StatusOK, response); err != nil {
		c.log.Println(err)
		return
	}
}

func (c *consumerHandler) updateConsumerLocation(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	consumerID, ok := vars["id"]
	if !ok {
		if err := consumerapi.Respond(rw, http.StatusBadRequest, "consumer_id param is not found"); err != nil {
			c.log.Println(err)
		}
	}

	var updateConsumerLocationRequest consumerapi.UpdateConsumerLocationRequest

	if err := consumerapi.BindJson(r, &updateConsumerLocationRequest); err != nil {
		c.log.Println(err)
		if err := consumerapi.Respond(rw, http.StatusBadRequest, "incorrect input data"); err != nil {
			c.log.Println(err)
		}
		return
	}

	consumerLocation := requestToUpdateConsumerLocation(&updateConsumerLocationRequest)

	data, err := c.consumerService.UpdateConsumerLocation(consumerLocation, consumerID)

	if err != nil {
		if err := consumerapi.Respond(rw, http.StatusBadRequest, err); err != nil {
			c.log.Println(err)
		}
		return
	}

	if data == nil {
		if err := consumerapi.Respond(rw, http.StatusInternalServerError, ""); err != nil {
			c.log.Println(err)
		}
		return
	}

	response := consumerLocationToResponse(*data)

	if err := consumerapi.Respond(rw, http.StatusOK, response); err != nil {
		c.log.Println(err)
		return
	}
}
