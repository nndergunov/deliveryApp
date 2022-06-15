package deliveryapi

import (
	"net/http"

	"github.com/nndergunov/deliveryApp/app/pkg/api/v1"
)

// Respond converts a Go value to JSON and sends it to the client.
func Respond(w http.ResponseWriter, status int, data any) error {
	// Convert the response value to JSON.
	jsonData, err := v1.Encode(data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return err
	}

	// Set the content type and headers once we know marshaling has succeeded.
	w.Header().Set("Content-Type", "application/json")

	// Write the status code to the response.
	w.WriteHeader(status)

	// Send the result back to the client.
	if _, err := w.Write(jsonData); err != nil {
		return err
	}

	return nil
}
