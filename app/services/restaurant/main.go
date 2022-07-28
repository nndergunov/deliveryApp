package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"

	"github.com/nndergunov/deliveryApp/app/pkg/api"
	"github.com/nndergunov/deliveryApp/app/pkg/configreader"
	"github.com/nndergunov/deliveryApp/app/pkg/grpcserver"
	"github.com/nndergunov/deliveryApp/app/pkg/logger"
	"github.com/nndergunov/deliveryApp/app/pkg/server"
	"github.com/nndergunov/deliveryApp/app/pkg/server/config"

	"github.com/nndergunov/deliveryApp/app/services/restaurant/api/v1/grpclogic"
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

	grpcDatabase, err := db.NewDatabase(dbURL)
	if err != nil {
		mainLogger.Fatalln(err)
	}

	grpcServiceInstance := service.NewService(grpcDatabase)

	grpcLogger := logger.NewLogger(os.Stdout, "grpc")
	grpcRawServer := grpclogic.NewRestaurantRawGRPCServer(grpcServiceInstance)
	grpcServer := grpcserver.NewGRPCServer(grpcRawServer, grpcLogger)

	grpcServerStopChan := make(chan interface{})

	grpcServer.StartListening(configreader.GetString("grpcserver.address"), grpcServerStopChan)

	restDatabase, err := db.NewDatabase(dbURL)
	if err != nil {
		mainLogger.Fatalln(err)
	}

	restServiceInstance := service.NewService(restDatabase)

	handlerLogger := logger.NewLogger(os.Stdout, "endpoint")
	endpointHandler := handlers.NewEndpointHandler(restServiceInstance, handlerLogger)

	apiLogger := logger.NewLogger(os.Stdout, "api")
	serverAPI := api.NewAPI(endpointHandler, apiLogger)

	serverLogger := logger.NewLogger(os.Stdout, "server")
	serverConfig := getServerConfig(serverAPI, nil, serverLogger)

	restServer := server.NewServer(serverConfig)

	serverStopChan := make(chan interface{})

	restServer.StartListening(serverStopChan)

	serverWG := new(sync.WaitGroup)
	numberOfServersRunning := 2

	serverWG.Add(numberOfServersRunning)

	go func(wg *sync.WaitGroup) {
		<-grpcServerStopChan

		wg.Done()
	}(serverWG)

	go func(wg *sync.WaitGroup) {
		<-serverStopChan

		wg.Done()
	}(serverWG)

	serverWG.Wait()
}

func getServerConfig(handler http.Handler, errorLog *log.Logger, serverLogger *logger.Logger) *config.Config {
	var (
		address          = configreader.GetString("httpServer.address")
		readTime         = configreader.GetDuration("httpServer.readTime")
		writeTime        = configreader.GetDuration("httpServer.writeTime")
		idleTime         = configreader.GetDuration("httpServer.idleTime")
		readerHeaderTime = configreader.GetDuration("httpServer.readerHeaderTime")
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
