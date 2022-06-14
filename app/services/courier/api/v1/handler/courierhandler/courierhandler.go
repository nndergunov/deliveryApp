package courierhandler

import (
	"io"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/nndergunov/deliveryApp/app/pkg/api/v1"
	"github.com/nndergunov/deliveryApp/app/pkg/logger"

	"courier/api/v1/courierapi"
	"courier/pkg/domain"
	"courier/pkg/service/courierservice"
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

const courierIDKey = "courier_id"

// NewCourierHandler creates an courierHandler value that handle a set of routes for the application.
func (c *courierHandler) handlerInit() {
	c.serveMux.HandleFunc("/status", c.statusHandler).Methods(http.MethodPost)

	c.serveMux.HandleFunc("/v1/couriers", c.insertNewCourier).Methods(http.MethodPost)
	c.serveMux.HandleFunc("/v1/couriers", c.getAllCourier).Methods(http.MethodGet)
	c.serveMux.HandleFunc("/v1/couriers/{"+courierIDKey+"}", c.deleteCourier).Methods(http.MethodDelete)
	c.serveMux.HandleFunc("/v1/couriers/{"+courierIDKey+"}", c.updateCourier).Methods(http.MethodPut)
	c.serveMux.HandleFunc("/v1/couriers/{"+courierIDKey+"}", c.getCourier).Methods(http.MethodGet)
	c.serveMux.HandleFunc("/v1/couriers-available/{"+courierIDKey+"}", c.updateCourierAvailable).Methods(http.MethodPut)

	c.serveMux.HandleFunc("/v1/couriers/{"+courierIDKey+"}/location", c.insertNewLocation).Methods(http.MethodPost)
	c.serveMux.HandleFunc("/v1/couriers/{"+courierIDKey+"}/location", c.updateLocation).Methods(http.MethodPut)
	c.serveMux.HandleFunc("/v1/couriers/{"+courierIDKey+"}/location", c.getLocation).Methods(http.MethodGet)
}

func (c *courierHandler) statusHandler(responseWriter http.ResponseWriter, _ *http.Request) {
	data := v1.Status{
		ServiceName: "courier",
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

func (c *courierHandler) insertNewCourier(rw http.ResponseWriter, r *http.Request) {
	var courierRequest courierapi.NewCourierRequest

	if err := courierapi.BindJson(r, &courierRequest); err != nil {
		c.log.Println(err)
		if err := courierapi.Respond(rw, http.StatusBadRequest, errIncorrectInputData.Error()); err != nil {
			c.log.Println(err)
		}
		return
	}

	courier := requestToNewCourier(&courierRequest)

	data, err := c.courierService.InsertCourier(courier)

	if err != nil && err == systemErr {
		c.log.Println(err)
		if err := courierapi.Respond(rw, http.StatusInternalServerError, ""); err != nil {
			c.log.Println(err)
		}
		return
	}

	if err != nil {
		c.log.Println(err)
		if err := courierapi.Respond(rw, http.StatusBadRequest, err.Error()); err != nil {
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

func (c *courierHandler) deleteCourier(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, ok := vars[courierIDKey]
	if !ok {
		if err := courierapi.Respond(rw, http.StatusBadRequest, errNoCourierIDParam.Error()); err != nil {
			c.log.Println(err)
		}
	}

	data, err := c.courierService.DeleteCourier(id)

	if err != nil && err == systemErr {
		c.log.Println(err)
		if err := courierapi.Respond(rw, http.StatusInternalServerError, ""); err != nil {
			c.log.Println(err)
		}
		return
	}

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
	vars := mux.Vars(r)
	id, ok := vars[courierIDKey]
	if !ok {
		if err := courierapi.Respond(rw, http.StatusBadRequest, errNoCourierIDParam.Error()); err != nil {
			c.log.Println(err)
		}
	}

	var updateCourierRequest courierapi.UpdateCourierRequest

	if err := courierapi.BindJson(r, &updateCourierRequest); err != nil {
		c.log.Println(err)
		if err := courierapi.Respond(rw, http.StatusBadRequest, errIncorrectInputData.Error()); err != nil {
			c.log.Println(err)
		}
		return
	}

	courier := requestToUpdateCourier(&updateCourierRequest)

	data, err := c.courierService.UpdateCourier(courier, id)

	if err != nil && err == systemErr {
		c.log.Println(err)
		if err := courierapi.Respond(rw, http.StatusInternalServerError, ""); err != nil {
			c.log.Println(err)
		}
		return
	}

	if err != nil {
		if err := courierapi.Respond(rw, http.StatusBadRequest, err); err != nil {
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
	vars := mux.Vars(r)
	id, ok := vars[courierIDKey]
	if !ok {
		if err := courierapi.Respond(rw, http.StatusBadRequest, errNoCourierIDParam.Error()); err != nil {
			c.log.Println(err)
		}
	}

	queryParams := r.URL.Query()
	queryParamsList := queryParams["available"]

	if queryParamsList == nil {
		if err := courierapi.Respond(rw, http.StatusBadRequest, errNoAvailableParam.Error()); err != nil {
			c.log.Println(err)
		}
	}
	available := queryParamsList[0]

	data, err := c.courierService.UpdateCourierAvailable(id, available)

	if err != nil && err == systemErr {
		c.log.Println(err)
		if err := courierapi.Respond(rw, http.StatusInternalServerError, ""); err != nil {
			c.log.Println(err)
		}
		return
	}

	if err != nil {
		c.log.Println(err)
		if err := courierapi.Respond(rw, http.StatusBadRequest, err); err != nil {
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
	param := domain.SearchParam{}

	var available string
	queryParams := r.URL.Query()

	availableList := queryParams["available"]
	if availableList != nil {
		available = availableList[0]
	}

	// latitude := queryParams["latitude"][0]
	// longitude := queryParams["longitude"][0]
	// radius := queryParams["radius"][0]

	if available != "" {
		param["available"] = available
	}

	//if latitude != "" && longitude != "" && radius != "" {
	//	param["latitude"] = longitude
	//	param["longitude"] = longitude
	//	param["radius"] = radius
	//}

	data, err := c.courierService.GetAllCourier(param)

	if err != nil && err == systemErr {
		c.log.Println(err)
		if err := courierapi.Respond(rw, http.StatusInternalServerError, err.Error()); err != nil {
			c.log.Println(err)
		}
		return
	}

	if err != nil {
		c.log.Println(err)
		if err := courierapi.Respond(rw, http.StatusBadRequest, err.Error()); err != nil {
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
	vars := mux.Vars(r)
	id, ok := vars[courierIDKey]
	if !ok {
		if err := courierapi.Respond(rw, http.StatusBadRequest, errNoCourierIDParam); err != nil {
			c.log.Println(err)
		}
	}

	data, err := c.courierService.GetCourier(id)

	if err != nil && err == systemErr {
		c.log.Println(err)
		if err := courierapi.Respond(rw, http.StatusInternalServerError, ""); err != nil {
			c.log.Println(err)
		}
		return
	}

	if err != nil {
		c.log.Println(err)
		if err := courierapi.Respond(rw, http.StatusBadRequest, err); err != nil {
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

func (c *courierHandler) insertNewLocation(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	courierID, ok := vars[courierIDKey]
	if !ok {
		if err := courierapi.Respond(rw, http.StatusBadRequest, errNoCourierIDParam.Error()); err != nil {
			c.log.Println(err)
		}
	}

	var locationRequest courierapi.NewLocationRequest

	if err := courierapi.BindJson(r, &locationRequest); err != nil {
		c.log.Println(err)
		if err := courierapi.Respond(rw, http.StatusBadRequest, errIncorrectInputData.Error()); err != nil {
			c.log.Println(err)
		}
		return
	}

	location := requestToNewLocation(&locationRequest)

	data, err := c.courierService.InsertLocation(location, courierID)

	if err != nil && err == systemErr {
		c.log.Println(err)
		if err := courierapi.Respond(rw, http.StatusInternalServerError, ""); err != nil {
			c.log.Println(err)
		}
		return
	}

	if err != nil {
		c.log.Println(err)
		if err := courierapi.Respond(rw, http.StatusBadRequest, err); err != nil {
			c.log.Println(err)
		}
		return
	}

	response := locationToResponse(*data)

	if err := courierapi.Respond(rw, http.StatusOK, response); err != nil {
		c.log.Println(err)
		return
	}
}

func (c *courierHandler) updateLocation(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	courierID, ok := vars[courierIDKey]
	if !ok {
		if err := courierapi.Respond(rw, http.StatusBadRequest, errNoCourierIDParam.Error()); err != nil {
			c.log.Println(err)
		}
	}

	var updateLocationRequest courierapi.UpdateLocationRequest

	if err := courierapi.BindJson(r, &updateLocationRequest); err != nil {
		c.log.Println(err)
		if err := courierapi.Respond(rw, http.StatusBadRequest, errIncorrectInputData.Error()); err != nil {
			c.log.Println(err)
		}
		return
	}

	location := requestToUpdateLocation(&updateLocationRequest)

	data, err := c.courierService.UpdateLocation(location, courierID)

	if err != nil && err == systemErr {
		c.log.Println(err)
		if err := courierapi.Respond(rw, http.StatusInternalServerError, ""); err != nil {
			c.log.Println(err)
		}
		return
	}

	if err != nil {
		c.log.Println(err)
		if err := courierapi.Respond(rw, http.StatusBadRequest, err); err != nil {
			c.log.Println(err)
		}
		return
	}

	response := locationToResponse(*data)

	if err := courierapi.Respond(rw, http.StatusOK, response); err != nil {
		c.log.Println(err)
		return
	}
}

func (c *courierHandler) getLocation(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID, ok := vars[courierIDKey]
	if !ok {
		if err := courierapi.Respond(rw, http.StatusBadRequest, errNoCourierIDParam.Error()); err != nil {
			c.log.Println(err)
		}
	}

	data, err := c.courierService.GetLocation(userID)

	if err != nil && err == systemErr {
		c.log.Println(err)
		if err := courierapi.Respond(rw, http.StatusInternalServerError, ""); err != nil {
			c.log.Println(err)
		}
		return
	}

	if err != nil {
		c.log.Println(err)
		if err := courierapi.Respond(rw, http.StatusBadRequest, err); err != nil {
			c.log.Println(err)
		}
		return
	}

	response := locationToResponse(*data)

	if err := courierapi.Respond(rw, http.StatusOK, response); err != nil {
		c.log.Println(err)
		return
	}
}
