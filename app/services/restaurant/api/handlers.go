package api

import (
	"io"
	"net/http"

	v1 "github.com/nndergunov/deliveryApp/app/services/restaurant/api/v1"
)

func (a *API) handlerInit() {
	a.mux.HandleFunc("/status", a.statusHandler)
	a.mux.HandleFunc("/openRestaurants", a.openRestaurantsHandler)
	a.mux.HandleFunc("/menu", a.menuHandler)
}

func (a API) statusHandler(responseWriter http.ResponseWriter, _ *http.Request) {
	data := v1.Status{
		Service: "restaurant",
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

func (a API) openRestaurantsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		// TODO return error "unsupported method".
	}

	// TODO logic.
}

func (a API) menuHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		// TODO return error "unsupported method".
	}

	// TODO logic.
}
