package api

import (
	"io"
	"net/http"

	v1 "github.com/nndergunov/deliveryApp/app/services/order/api/v1"
)

func (a *API) handlerInit() {
	a.mux.HandleFunc("/status", a.statusHandler)
	a.mux.HandleFunc("/orderRestaurants", a.orderRestaurantsHandler)
	a.mux.HandleFunc("/orderMenu", a.orderMenuHandler)
	a.mux.HandleFunc("/finishedOrder", a.finishedOrderHandler)
}

func (a API) statusHandler(responseWriter http.ResponseWriter, _ *http.Request) {
	data := v1.Status{
		Service: "order",
		IsUp:    "up",
	}

	status, err := v1.EncodeIndent(data, "", " ")
	if err != nil {
		a.log.Println(err)
	}

	_, err = io.WriteString(responseWriter, string(status))
	if err != nil {
		a.log.Printf("status write: %v", err)

		return
	}

	a.log.Printf("gave status %s", data.IsUp)
}

func (a API) orderRestaurantsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		// TODO return error "unsupported method".
	}

	// TODO logic.
}

func (a API) orderMenuHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		// TODO return error "unsupported method".
	}

	// TODO logic.
}

func (a API) finishedOrderHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		// TODO return error "unsupported method".
	}

	// TODO logic.
}
