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
	"github.com/nndergunov/deliveryApp/app/services/restaurant/pkg/db"
	"github.com/nndergunov/deliveryApp/app/services/restaurant/pkg/service"
)

const configFile = "config.yaml"

func main() {
	mainLogger := logger.NewLogger(os.Stdout, "main")

	err := configreader.SetConfigFile(configFile)
	if err != nil {
		mainLogger.Fatalln(err)
	}

	dbURL := fmt.Sprintf("host=" + configreader.GetString("database.host") +
		" port=" + configreader.GetString("database.port") +
		" user=" + configreader.GetString("database.user") +
		" password=" + configreader.GetString("database.password") +
		" dbname=" + configreader.GetString("database.dbName") +
		" sslmode=" + configreader.GetString("database.sslmode"))

	database, err := db.NewDatabase(dbURL)
	if err != nil {
		mainLogger.Fatalln(err)
	}

	serviceInstance := service.NewService(database)
	handlerLogger := logger.NewLogger(os.Stdout, "endpoint")
	endpointHandler := handlers.NewEndpointHandler(serviceInstance, handlerLogger)

	apiLogger := logger.NewLogger(os.Stdout, "api")
	serverAPI := api.NewAPI(endpointHandler, apiLogger)

	serverLogger := logger.NewLogger(os.Stdout, "server")
	serverConfig := getServerConfig(serverAPI, nil, serverLogger)
	serviceServer := server.NewServer(serverConfig)

	serverStopChan := make(chan interface{})

	serviceServer.StartListening(serverStopChan)

	<-serverStopChan
}

func getServerConfig(handler http.Handler, errorLog *log.Logger, serverLogger *logger.Logger) *config.Config {
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
	}
}
