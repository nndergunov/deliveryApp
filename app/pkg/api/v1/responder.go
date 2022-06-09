package v1

import (
	"fmt"
	"net/http"
)

func Respond(response any, w http.ResponseWriter) error {
	data, err := Encode(response)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)

		return fmt.Errorf("encoding data: %w", err)
	}

	w.WriteHeader(http.StatusOK)

	_, err = w.Write(data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)

		return fmt.Errorf("sending data: %w", err)
	}

	return nil
}

func RespondWithError(errMsg string, httpStatus int, w http.ResponseWriter) error {
	data, err := Encode(ServiceError{
		HTTPStatus: httpStatus,
		ErrorText:  errMsg,
	})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)

		return fmt.Errorf("encoding service error: %w", err)
	}

	w.WriteHeader(httpStatus)

	_, err = w.Write(data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)

		return fmt.Errorf("sending error message: %w", err)
	}

	return nil
}
