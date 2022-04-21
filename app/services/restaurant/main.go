package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/nndergunov/deliveryApp/app/pkg/api"
	"github.com/nndergunov/deliveryApp/app/pkg/configreader"
	"github.com/nndergunov/deliveryApp/app/pkg/logger"
	"github.com/nndergunov/deliveryApp/app/pkg/server"
	"github.com/nndergunov/deliveryApp/app/pkg/server/config"
	"github.com/nndergunov/deliveryApp/app/services/restaurant/api/v1/handlers"
	"github.com/nndergunov/deliveryApp/app/services/restaurant/pkg/app"
)

const configFile = "config.yaml"

func main() {
	mainLogger := logger.NewLogger(os.Stdout, "main")

	appInstance := app.NewApp()
	handlerLogger := logger.NewLogger(os.Stdout, "endpoint")
	endpointHandler := handlers.NewEndpointHandler(appInstance, handlerLogger)

	apiLogger := logger.NewLogger(os.Stdout, "api")
	serverAPI := api.NewAPI(endpointHandler, apiLogger)

	serverLogger := logger.NewLogger(os.Stdout, "server")

	serverConfig, err := getServerConfig(serverAPI, nil, serverLogger)
	if err != nil {
		mainLogger.Println(err)
	}

	serviceServer := server.NewServer(serverConfig)
	serverStopChan := make(chan interface{})

	serviceServer.StartListening(serverStopChan)

	<-serverStopChan
}

func getServerConfig(handler http.Handler, errorLog *log.Logger, serverLogger *logger.Logger) (*config.Config, error) {
	if err := configreader.SetConfigFile(configFile); err != nil {
		return nil, fmt.Errorf("config read: %w", err)
	}

	var (
		address          = configreader.GetString("server.address")
		readTime         = configreader.GetDuration("server.readTime")
		writeTime        = configreader.GetDuration("server.writeTime")
		idleTime         = configreader.GetDuration("server.idleTime")
		readerHeaderTime = configreader.GetDuration("server.readerHeaderTime")
	)

	return &config.Config{
		Address:           address,
		ReadTimeout:       readTime,
		WriteTimeout:      writeTime,
		IdleTimeout:       idleTime,
		ReadHeaderTimeout: readerHeaderTime,
		ErrorLog:          errorLog,
		ServerLogger:      serverLogger,
		Handler:           handler,
	}, nil
}
