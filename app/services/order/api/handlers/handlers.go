package handlers

import (
	"io"
	"net/http"

	v1 "github.com/nndergunov/deliveryApp/app/pkg/api/v1"
	"github.com/nndergunov/deliveryApp/app/pkg/logger"
)

type endpointHandler struct {
	mux *http.ServeMux
	log *logger.Logger
}

func NewEndpointHandler(log *logger.Logger) *http.ServeMux {
	mux := http.NewServeMux()

	handler := endpointHandler{
		mux: mux,
		log: log,
	}

	handler.handlerInit()

	return handler.mux
}

func (e *endpointHandler) handlerInit() {
	e.mux.HandleFunc("/status", e.statusHandler)
	e.mux.HandleFunc("/orderRestaurants", e.orderRestaurantsHandler)
	e.mux.HandleFunc("/orderMenu", e.orderMenuHandler)
	e.mux.HandleFunc("/finishedOrder", e.finishedOrderHandler)
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

func (e endpointHandler) orderRestaurantsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		// TODO return error "unsupported method".
	}

	// TODO logic.
}

func (e endpointHandler) orderMenuHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		// TODO return error "unsupported method".
	}

	// TODO logic.
}

func (e endpointHandler) finishedOrderHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		// TODO return error "unsupported method".
	}

	// TODO logic.
}
