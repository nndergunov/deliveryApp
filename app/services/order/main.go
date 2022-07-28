package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"

	"github.com/nndergunov/deliveryApp/app/pkg/api"
	v1 "github.com/nndergunov/deliveryApp/app/pkg/api/v1"
	"github.com/nndergunov/deliveryApp/app/pkg/configreader"
	"github.com/nndergunov/deliveryApp/app/pkg/grpcserver"
	"github.com/nndergunov/deliveryApp/app/pkg/logger"
	"github.com/nndergunov/deliveryApp/app/pkg/messagebroker/publisher"
	"github.com/nndergunov/deliveryApp/app/pkg/server"
	"github.com/nndergunov/deliveryApp/app/pkg/server/config"

	"github.com/nndergunov/deliveryApp/app/services/order/api/v1/grpclogic"
	"github.com/nndergunov/deliveryApp/app/services/order/api/v1/handlers"
	"github.com/nndergunov/deliveryApp/app/services/order/pkg/clients/accountingclient"
	"github.com/nndergunov/deliveryApp/app/services/order/pkg/clients/restaurantclient"
	"github.com/nndergunov/deliveryApp/app/services/order/pkg/db"
	"github.com/nndergunov/deliveryApp/app/services/order/pkg/service"
)

const configFile = "config.yaml"

func main() {
	mainLogger := logger.NewLogger(os.Stdout, "main")

	err := configreader.SetConfigFile(configFile)
	if err != nil {
		mainLogger.Println(err)
	}

	dbURL := fmt.Sprintf("host=" + configreader.GetString("database.host") +
		" port=" + configreader.GetString("database.port") +
		" user=" + configreader.GetString("database.user") +
		" password=" + configreader.GetString("database.password") +
		" dbname=" + configreader.GetString("database.dbName") +
		" sslmode=" + configreader.GetString("database.sslmode"))

	publisherURL := configreader.GetString("publisher.url")

	services := configreader.GetMap("services")

	grpcDatabase, err := db.NewDatabase(dbURL)
	if err != nil {
		mainLogger.Fatalln(err)
	}

	grpcNotificationer, err := publisher.NewEventPublisher(publisherURL)
	if err != nil {
		mainLogger.Fatalln(err)
	}

	grpcAccountingClient := accountingclient.NewAccountingClient(services["accounting"])
	grpcRestaurantClient := restaurantclient.NewRestaurantClient(services["restaurant"])

	grpcServiceInstance := service.NewService(grpcDatabase, grpcNotificationer, grpcAccountingClient, grpcRestaurantClient)

	grpcLogger := logger.NewLogger(os.Stdout, "grpc")
	grpcRawServer := grpclogic.NewOrderRawGRPCServer(grpcServiceInstance)
	grpcServer := grpcserver.NewGRPCServer(grpcRawServer, grpcLogger)

	grpcServerStopChan := make(chan interface{})

	grpcServer.StartListening(configreader.GetString("grpcserver.address"), grpcServerStopChan)

	restDatabase, err := db.NewDatabase(dbURL)
	if err != nil {
		mainLogger.Fatalln(err)
	}

	restNotificationer, err := publisher.NewEventPublisher(publisherURL)
	if err != nil {
		mainLogger.Fatalln(err)
	}

	restAccountingClient := accountingclient.NewAccountingClient(services["accounting"])
	restRestaurantClient := restaurantclient.NewRestaurantClient(services["restaurant"])

	restServiceInstance := service.NewService(restDatabase, restNotificationer, restAccountingClient, restRestaurantClient)
	handlerLogger := logger.NewLogger(os.Stdout, "endpoint")
	endpointHandler := handlers.NewEndpointHandler(restServiceInstance, handlerLogger)

	apiLogger := logger.NewLogger(os.Stdout, "api")
	serverAPI := api.NewAPI(endpointHandler, apiLogger)

	serverLogger := logger.NewLogger(os.Stdout, "server")

	serverConfig, err := getServerConfig(v1.EnableCORS(serverAPI), nil, serverLogger)
	if err != nil {
		mainLogger.Println(err)
	}

	serviceServer := server.NewServer(serverConfig)
	serverStopChan := make(chan interface{})

	serviceServer.StartListening(serverStopChan)

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

func getServerConfig(handler http.Handler, errorLog *log.Logger, serverLogger *logger.Logger) (*config.Config, error) {
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
