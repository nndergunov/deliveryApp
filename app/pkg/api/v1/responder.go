package v1

import (
	"fmt"
	"net/http"
)

func Respond(response any, w http.ResponseWriter) error {
	data, err := Encode(response)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)

		return fmt.Errorf("api.Respond: %w", err)
	}

	w.WriteHeader(http.StatusOK)

	_, err = w.Write(data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)

		return fmt.Errorf("api.Respond: %w", err)
	}

	return nil
}
