// Package handler contains a small handler framework extension.
package courierhandler

import (
	"courier/api/v1/courierapi"
	"courier/pkg/service/courierservice"
	"github.com/gorilla/mux"
	"github.com/nndergunov/deliveryApp/app/pkg/logger"
	"net/http"
)

type Params struct {
	Logger         *logger.Logger
	CourierService courierservice.CourierService
}

// courierHandler is the entrypoint into our application
type courierHandler struct {
	serveMux       *mux.Router
	log            *logger.Logger
	courierService courierservice.CourierService
}

// NewCourierHandler returns new http multiplexer with configured endpoints.
func NewCourierHandler(p Params) *mux.Router {
	serveMux := mux.NewRouter()

	handler := courierHandler{
		serveMux:       serveMux,
		log:            p.Logger,
		courierService: p.CourierService,
	}

	handler.handlerInit()

	return handler.serveMux
}

// NewCourierHandler creates an courierHandler value that handle a set of routes for the application.
func (c *courierHandler) handlerInit() {

	const version = "/v1"
	const courier = "/courier"

	c.serveMux.HandleFunc(version+courier+"/new", c.insertNewCourier).Methods(http.MethodPost)
	c.serveMux.HandleFunc(version+courier+"/remove", c.removeCourier).Methods(http.MethodPost)
	c.serveMux.HandleFunc(version+courier+"/update", c.updateCourier).Methods(http.MethodPut)
	c.serveMux.HandleFunc(version+courier+"/update-available", c.updateCourierAvailable).Methods(http.MethodPut)
	c.serveMux.HandleFunc(version+courier+"/get-all", c.getAllCourier).Methods(http.MethodGet)
	c.serveMux.HandleFunc(version+courier+"/get", c.getCourier).Methods(http.MethodGet)
}

func (c *courierHandler) insertNewCourier(rw http.ResponseWriter, r *http.Request) {
	var courierRequest courierapi.NewCourierRequest

	if err := courierapi.BindJson(r, &courierRequest); err != nil {
		c.log.Println(err)
		if err := courierapi.Respond(rw, http.StatusBadRequest, "incorrect input data"); err != nil {
			c.log.Println(err)
		}
		return
	}

	courier := requestToNewCourier(&courierRequest)

	data, err := c.courierService.InsertCourier(courier)
	if err != nil {
		if err := courierapi.Respond(rw, http.StatusBadRequest, err); err != nil {
			c.log.Println(err)
		}
		return
	}

	if data == nil {
		if err := courierapi.Respond(rw, http.StatusInternalServerError, ""); err != nil {
			c.log.Println(err)
		}
		return
	}

	response := courierToResponse(*data)

	if err := courierapi.Respond(rw, http.StatusOK, response); err != nil {
		c.log.Println(err)
		return
	}

}

func (c *courierHandler) removeCourier(rw http.ResponseWriter, r *http.Request) {

	queryParams := r.URL.Query()
	id := queryParams["id"][0]
	if id == "" {
		if err := courierapi.Respond(rw, http.StatusBadRequest, "id param is not found"); err != nil {
			c.log.Println(err)
		}
	}

	data, err := c.courierService.RemoveCourier(id)
	if err != nil {
		if err := courierapi.Respond(rw, http.StatusBadRequest, err); err != nil {
			c.log.Println(err)
		}
		return
	}

	if err := courierapi.Respond(rw, http.StatusOK, data); err != nil {
		c.log.Println(err)
		return
	}
}

func (c *courierHandler) updateCourier(rw http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()
	id := queryParams["id"][0]
	if id == "" {
		if err := courierapi.Respond(rw, http.StatusBadRequest, "id param is not found"); err != nil {
			c.log.Println(err)
		}
	}

	var updateCourierRequest courierapi.UpdateCourierRequest

	if err := courierapi.BindJson(r, &updateCourierRequest); err != nil {
		c.log.Println(err)
		if err := courierapi.Respond(rw, http.StatusBadRequest, "incorrect input data"); err != nil {
			c.log.Println(err)
		}
		return
	}

	courier := requestToUpdateCourier(&updateCourierRequest)

	data, err := c.courierService.UpdateCourier(courier, id)

	if err != nil {
		if err := courierapi.Respond(rw, http.StatusBadRequest, err); err != nil {
			c.log.Println(err)
		}
		return
	}

	if data == nil {
		if err := courierapi.Respond(rw, http.StatusInternalServerError, ""); err != nil {
			c.log.Println(err)
		}
		return
	}

	response := courierToResponse(*data)

	if err := courierapi.Respond(rw, http.StatusOK, response); err != nil {
		c.log.Println(err)
		return
	}
}

func (c *courierHandler) updateCourierAvailable(rw http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()
	id := queryParams["id"][0]
	if id == "" {
		if err := courierapi.Respond(rw, http.StatusBadRequest, "id param is not found"); err != nil {
			c.log.Println(err)
		}
	}

	available := queryParams["available"][0]
	if available == "" {
		if err := courierapi.Respond(rw, http.StatusBadRequest, "available param is not found"); err != nil {
			c.log.Println(err)
		}
	}

	data, err := c.courierService.UpdateCourierAvailable(id, available)
	if err != nil {
		if err := courierapi.Respond(rw, http.StatusBadRequest, err); err != nil {
			c.log.Println(err)
		}
		return
	}

	if data == nil {
		if err := courierapi.Respond(rw, http.StatusInternalServerError, ""); err != nil {
			c.log.Println(err)
		}
		return
	}

	response := courierToResponse(*data)
	if err := courierapi.Respond(rw, http.StatusOK, response); err != nil {
		c.log.Println(err)
		return
	}
}

func (c *courierHandler) getAllCourier(rw http.ResponseWriter, r *http.Request) {
	data, err := c.courierService.GetAllCourier()

	if err != nil {
		if err := courierapi.Respond(rw, http.StatusBadRequest, err); err != nil {
			c.log.Println(err)
		}
		return
	}

	if data == nil {
		if err := courierapi.Respond(rw, http.StatusInternalServerError, ""); err != nil {
			c.log.Println(err)
		}
		return
	}

	response := courierListToResponse(data)

	if err := courierapi.Respond(rw, http.StatusOK, response); err != nil {
		c.log.Println(err)
		return
	}
}

func (c *courierHandler) getCourier(rw http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()
	id := queryParams["id"][0]
	if id == "" {
		if err := courierapi.Respond(rw, http.StatusBadRequest, "id param is not found"); err != nil {
			c.log.Println(err)
		}
	}

	data, err := c.courierService.GetCourier(id)

	if err != nil {
		if err := courierapi.Respond(rw, http.StatusBadRequest, err); err != nil {
			c.log.Println(err)
		}
		return
	}

	if data == nil {
		if err := courierapi.Respond(rw, http.StatusInternalServerError, ""); err != nil {
			c.log.Println(err)
		}
		return
	}

	response := courierToResponse(*data)

	if err := courierapi.Respond(rw, http.StatusOK, response); err != nil {
		c.log.Println(err)
		return
	}
}
