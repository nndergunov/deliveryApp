package main

import (
	"consumer/api/v1/handler/consumerhandler"
	"consumer/pkg/db"
	"consumer/pkg/service/consumerservice"
	"consumer/pkg/storage/consumerstorage"
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
	consumerStorage := consumerstorage.NewConsumerStorage(consumerstorage.Params{DB: database})

	consumerService := consumerservice.NewConsumerService(consumerservice.Params{
		ConsumerStorage: consumerStorage,
		Logger:          logger.NewLogger(os.Stdout, "service: ")})

	consumerHandler := consumerhandler.NewConsumerHandler(consumerhandler.Params{
		Logger:          logger.NewLogger(os.Stdout, "endpoint: "),
		ConsumerService: consumerService,
	})

	apiLogger := logger.NewLogger(os.Stdout, "api: ")
	serverAPI := api.NewAPI(consumerHandler, apiLogger)

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
