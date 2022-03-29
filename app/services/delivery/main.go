package main

import (
	"log"
	"os"

	"github.com/nndergunov/deliveryApp/app/services/delivery/api"
	"github.com/nndergunov/deliveryApp/app/services/delivery/cmd/server"
)

func main() {
	// apiLogger := logger.NewLogger(os.Stdout, "delivery api")

	deliveryAPI := api.NewAPI(apiLogger)

	serverLogger := log.New(os.Stdout, "delivery server ", log.LstdFlags)

	deliveryServer := server.NewServer(deliveryAPI, serverLogger)

	deliveryServer.StartListening()
}
