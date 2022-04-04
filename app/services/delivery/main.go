package main

import (
	"fmt"
	"os"
	"time"

	"github.com/nndergunov/deliveryApp/app/pkg/apilib"
	"github.com/nndergunov/deliveryApp/app/pkg/logger"
	"github.com/nndergunov/deliveryApp/app/services/delivery/api"
	"github.com/nndergunov/deliveryApp/app/services/delivery/cmd/server"
	"github.com/nndergunov/deliveryApp/app/services/delivery/cmd/server/config"
	"github.com/spf13/viper"
)

func main() {
	mainLogger := logger.NewLogger(os.Stdout, "main")

	handlerLogger := logger.NewLogger(os.Stdout, "endpoint")
	endpointHandler := api.NewEndpointHandler(handlerLogger)

	apiLogger := logger.NewLogger(os.Stdout, "apilib")
	serverAPI := apilib.NewAPI(endpointHandler, apiLogger)

	serverConfig, err := getServerConfig()
	if err != nil {
		mainLogger.Println(err)
	}

	serverLogger := logger.NewLogger(os.Stdout, "server")
	serviceServer := server.NewServer(serverAPI, serverConfig, serverLogger)
	serverStopChan := make(chan interface{})

	serviceServer.StartListening(serverStopChan)

	<-serverStopChan
}

func getServerConfig() (*config.Config, error) {
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
	}, nil
}
