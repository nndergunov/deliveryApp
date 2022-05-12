package handlers

import (
	"encoding/json"
	"net/http"
)

// Respond converts a Go value to JSON and sends it to the client.
func Respond(w http.ResponseWriter, data any, err error) error {

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return nil
	}

	if data == nil {
		w.WriteHeader(http.StatusInternalServerError)
		return nil
	}

	// Convert the response value to JSON.
	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}

	// Set the content type and headers once we know marshaling has succeeded.
	w.Header().Set("Content-Type", "application/json")

	// Write the status code to the response.
	w.WriteHeader(http.StatusOK)

	// Send the result back to the client.
	if _, err := w.Write(jsonData); err != nil {
		return err
	}

	return nil
}
