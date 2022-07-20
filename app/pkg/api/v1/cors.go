package v1

import (
	"net/http"

	"github.com/gorilla/handlers"
)

func EnableCORS(api http.Handler) http.Handler {
	headersOK := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type"})
	originsOK := handlers.AllowedOrigins([]string{"*"})
	methodsOK := handlers.AllowedMethods([]string{"GET", "POST", "OPTIONS", "DELETE", "PUT"})

	return handlers.CORS(headersOK, originsOK, methodsOK)(api)
}
