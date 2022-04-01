package api

import (
	"io"
	"net/http"

	v1 "github.com/nndergunov/deliveryApp/app/services/courier/api/v1"
)

func (a *API) handlerInit() {
	a.mux.HandleFunc("/status", a.statusHandler)
	a.mux.HandleFunc("/paid", a.paidHandler)
	a.mux.HandleFunc("/activeCouriers", a.activeCouriersHandler)
	a.mux.HandleFunc("/deliveryInfo", a.deliveryInfoHandler)
}

func (a API) statusHandler(responseWriter http.ResponseWriter, _ *http.Request) {
	data := v1.Status{
		Service: "courier",
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

func (a API) paidHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		// TODO return error "unsupported method".
	}

	// TODO logic.
}

func (a API) activeCouriersHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		// TODO return error "unsupported method".
	}

	// TODO logic.
}

func (a API) deliveryInfoHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		// TODO return error "unsupported method".
	}

	// TODO logic.
}
