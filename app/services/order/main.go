package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/nndergunov/deliveryApp/app/pkg/api"
	"github.com/nndergunov/deliveryApp/app/pkg/logger"
	"github.com/nndergunov/deliveryApp/app/pkg/server"
	"github.com/nndergunov/deliveryApp/app/pkg/server/config"
	"github.com/nndergunov/deliveryApp/app/services/order/api/handlers"
	"github.com/spf13/viper"
)

func main() {
	mainLogger := logger.NewLogger(os.Stdout, "main")

	handlerLogger := logger.NewLogger(os.Stdout, "endpoint")
	endpointHandler := handlers.NewEndpointHandler(handlerLogger)

	apiLogger := logger.NewLogger(os.Stdout, "apilib")
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
	viper.SetConfigFile("config.yaml")

	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("config read: %w", err)
	}

	var (
		address          = viper.GetString("server.address")
		readTime         = time.Duration(viper.GetInt("server.readTime")) * time.Second
		writeTime        = time.Duration(viper.GetInt("server.writeTime")) * time.Second
		idleTime         = time.Duration(viper.GetInt("server.idleTime")) * time.Second
		readerHeaderTime = time.Duration(viper.GetInt("server.readerHeaderTime")) * time.Second
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
