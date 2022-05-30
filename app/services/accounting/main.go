package main

import (
	"accounting/api/v1/handlers"
	"accounting/pkg/db"
	"accounting/pkg/service/consumerservice"
	"accounting/pkg/service/courierservice"
	"accounting/pkg/service/restaurantrservice"
	"accounting/pkg/storage/consumerstorage"
	"accounting/pkg/storage/courierrstorage"
	"accounting/pkg/storage/restaurantstorage"

	"github.com/nndergunov/deliveryApp/app/pkg/api"
	"github.com/nndergunov/deliveryApp/app/pkg/configreader"
	"github.com/nndergunov/deliveryApp/app/pkg/logger"
	"github.com/nndergunov/deliveryApp/app/pkg/server"
	"github.com/nndergunov/deliveryApp/app/pkg/server/config"

	"log"
	"net/http"
	"os"
)

const configFile = "/config.yaml"

func main() {

	// Construct the application logger.
	log := logger.NewLogger(os.Stdout, "main: ")

	// Perform the startup and shutdown sequence.
	if err := run(log); err != nil {
		log.Fatal("startup", "ERROR", err)
	}
}

func run(log *logger.Logger) error {
	confPath, err := os.Getwd()
	if err != nil {
		return err
	}

	err = configreader.SetConfigFile(confPath + configFile)
	if err != nil {
		return err
	}

	log.Println("starting service", "version", configreader.GetString("buildmode"))
	defer log.Println("shutdown complete")

	database, err := db.OpenDB("postgres", configreader.GetString("DB.dev"))
	if err != nil {
		return err
	}

	consumerAccountingStorage := consumerstorage.NewStorage(consumerstorage.Params{DB: database})

	serviceLogger := logger.NewLogger(os.Stdout, "service: ")
	consumerAccountingService := consumerservice.NewService(consumerservice.Params{
		Storage: consumerAccountingStorage,
		Logger:  serviceLogger})

	courierAccountingStorage := courierstorage.NewStorage(courierstorage.Params{DB: database})

	courierAccountingService := courierservice.NewService(courierservice.Params{
		Storage: courierAccountingStorage,
		Logger:  serviceLogger})

	restaurantAccountingStorage := restaurantstorage.NewStorage(restaurantstorage.Params{DB: database})

	restaurantAccountingService := restaurantservice.NewService(restaurantservice.Params{
		Storage: restaurantAccountingStorage,
		Logger:  serviceLogger})

	ConsumerHandler := handlers.NewAPIMux(handlers.Params{
		Logger:            logger.NewLogger(os.Stdout, "endpoint: "),
		ConsumerService:   consumerAccountingService,
		CourierService:    courierAccountingService,
		RestaurantService: restaurantAccountingService,
	})

	apiLogger := logger.NewLogger(os.Stdout, "api: ")
	serverAPI := api.NewAPI(ConsumerHandler, apiLogger)

	serverLogger := logger.NewLogger(os.Stdout, "server: ")
	serverConfig := getServerConfig(serverAPI, nil, serverLogger)

	serviceServer := server.NewServer(serverConfig)

	serverErrors := make(chan interface{})

	serviceServer.StartListening(serverErrors)

	<-serverErrors

	return nil
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
