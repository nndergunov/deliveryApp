package main

import (
	"fmt"
	"os"
	"time"

	"github.com/nndergunov/deliveryApp/app/libs/logger"
	"github.com/nndergunov/deliveryApp/app/services/delivery/api"
	"github.com/nndergunov/deliveryApp/app/services/delivery/cmd/server"
	"github.com/nndergunov/deliveryApp/app/services/delivery/cmd/server/config"
	"github.com/spf13/viper"
)

func main() {
	mainLogger := logger.NewLogger(os.Stdout, "main")

	apiLogger := logger.NewLogger(os.Stdout, "delivery api")
	serverAPI := api.NewAPI(apiLogger)

	serverConfig, err := getServerConfig()
	if err != nil {
		mainLogger.Println(err)
	}

	serverLogger := logger.NewLogger(os.Stdout, "service server")
	serviceServer := server.NewServer(serverAPI, serverConfig, serverLogger)
	serverStopChan := make(chan interface{})

	serviceServer.StartListening(serverStopChan)

	err = write("press Enter to stop server\n")
	if err != nil {
		mainLogger.Println(err)
	}

	_, err = fmt.Scanln()
	if err != nil {
		mainLogger.Println(err)
	}

	serviceServer.Shutdown()

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

func write(s string) error {
	_, err := os.Stdout.WriteString(s)
	if err != nil {
		return fmt.Errorf("writing to stdout: %w", err)
	}

	return nil
}
