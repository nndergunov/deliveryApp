package main

import (
	"courier/api/v1/handler/courierhandler"
	"courier/db"
	"courier/db/storage"
	"courier/service"
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
	log := logger.NewLogger(os.Stdout, "courier-api")

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

	database, err := db.Open("postgres", configreader.GetString("DB.dev"))
	if err != nil {
		return err
	}
	defer database.Close()

	log.Println("starting service", "version", configreader.GetString("buildmode"))
	defer log.Println("shutdown complete")

	newCourierStorage, err := storage.NewCourierStorage(storage.Params{
		DB: database,
	})
	if err != nil {
		return err
	}

	courierService, err := service.NewCourierService(service.Params{
		CourierStorage: newCourierStorage,
		Logger:         logger.NewLogger(os.Stdout, "courier-service"),
	})
	if err != nil {
		return err
	}

	handlerLogger := logger.NewLogger(os.Stdout, "endpoint")
	courierHandler := courierhandler.NewCourierHandler(courierhandler.Params{
		Logger:         handlerLogger,
		CourierService: courierService,
	})

	apiLogger := logger.NewLogger(os.Stdout, "api")
	serverAPI := api.NewAPI(courierHandler, apiLogger)

	serverLogger := logger.NewLogger(os.Stdout, "server")
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
