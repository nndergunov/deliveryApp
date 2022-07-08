package handlers

import (
	"io"
	"net/http"

	"github.com/gorilla/mux"
	v1 "github.com/nndergunov/deliveryApp/app/pkg/api/v1"
	"github.com/nndergunov/deliveryApp/app/pkg/logger"
	"github.com/nndergunov/deliveryApp/app/services/kitchen/pkg/service"
)

const kitchenIDKey = "kitchenID"

type endpointHandler struct {
	service  service.AppService
	serveMux *mux.Router
	log      *logger.Logger
}

// NewEndpointHandler returns new http multiplexer with configured endpoints.
func NewEndpointHandler(serviceInstance service.AppService, log *logger.Logger) *mux.Router {
	serveMux := mux.NewRouter()

	handler := endpointHandler{
		service:  serviceInstance,
		serveMux: serveMux,
		log:      log,
	}

	handler.handlerInit()

	return handler.serveMux
}

func (e *endpointHandler) handlerInit() {
	e.serveMux.HandleFunc("/status", e.statusHandler)

	e.serveMux.HandleFunc("/v1/tasks/{"+kitchenIDKey+"}", e.returnTasks).Methods(http.MethodGet)
}

func (e endpointHandler) statusHandler(responseWriter http.ResponseWriter, _ *http.Request) {
	data := v1.Status{
		ServiceName: "kitchen",
		IsUp:        "up",
	}

	status, err := v1.EncodeIndent(data, "", " ")
	if err != nil {
		e.log.Println(err)

		responseWriter.WriteHeader(http.StatusInternalServerError)

		return
	}

	_, err = io.WriteString(responseWriter, string(status))
	if err != nil {
		e.log.Println(err)

		responseWriter.WriteHeader(http.StatusInternalServerError)

		return
	}

	responseWriter.WriteHeader(http.StatusOK)

	e.log.Printf("gave status %s", data.IsUp)
}

func (e endpointHandler) returnTasks(responseWriter http.ResponseWriter, request *http.Request) {
	// swagger:operation GET /tasks/{id} returnTasks
	//
	// Returns tasks for the specified restaurant
	//
	// ---
	// produces:
	// - application/json
	// responses:
	//   '200':
	//     description: requested data
	kitchenID, err := getIDFromEndpoint(kitchenIDKey, request)
	if err != nil {
		e.log.Println(err)

		responseWriter.WriteHeader(http.StatusBadRequest)

		return
	}

	tasks, err := e.service.GetTasks(kitchenID)
	if err != nil {
		e.log.Println(err)

		responseWriter.WriteHeader(http.StatusInternalServerError)

		return
	}

	response := tasksToResponse(tasks)

	err = v1.Respond(response, responseWriter)
	if err != nil {
		e.log.Println(err)
	}
}
