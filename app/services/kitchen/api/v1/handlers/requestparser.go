package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func getIDFromEndpoint(code string, r *http.Request) (int, error) {
	vars := mux.Vars(r)

	idVar := vars[code]
	if idVar == "" {
		return 0, fmt.Errorf("getIDFromEndpoint: %w", errNoIDInEndpoint)
	}

	id, err := strconv.Atoi(idVar)
	if err != nil {
		return 0, fmt.Errorf("getIDFromEndpoint: %w", err)
	}

	return id, nil
}
