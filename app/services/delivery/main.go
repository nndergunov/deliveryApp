package main

import (
	"os"
	"time"

	"github.com/nndergunov/deliveryApp/app/libs/logger"
	"github.com/nndergunov/deliveryApp/app/services/delivery/api"
	"github.com/nndergunov/deliveryApp/app/services/delivery/cmd/server"
	"github.com/nndergunov/deliveryApp/app/services/delivery/cmd/server/config"
)

func main() {
	apiLogger := logger.NewLogger(os.Stdout, "delivery api")

	deliveryAPI := api.NewAPI(apiLogger)

	var (
		address          = ":8080"
		readTime         = 5 * time.Second
		writeTime        = 5 * time.Second
		idleTime         = 300 * time.Second
		readerHeaderTime = 5 * time.Second
	) // TODO change this to viper and yaml.

	serverConfig := config.Config{
		Address:           address,
		ReadTimeout:       readTime,
		WriteTimeout:      writeTime,
		IdleTimeout:       idleTime,
		ReadHeaderTimeout: readerHeaderTime,
	}

	serverLogger := logger.NewLogger(os.Stdout, "delivery server")

	deliveryServer := server.NewServer(deliveryAPI, serverConfig, serverLogger)

	deliveryServer.StartListening()
}
