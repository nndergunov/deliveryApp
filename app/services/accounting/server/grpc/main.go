package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"github.com/nndergunov/deliveryApp/app/pkg/configreader"
	"github.com/nndergunov/deliveryApp/app/pkg/logger"
	"google.golang.org/grpc"

	"github.com/nndergunov/deliveryApp/app/services/accounting/api/v1/grpc/handler"
	pb "github.com/nndergunov/deliveryApp/app/services/accounting/api/v1/grpc/proto"
	"github.com/nndergunov/deliveryApp/app/services/accounting/pkg/db"
	"github.com/nndergunov/deliveryApp/app/services/accounting/pkg/service/accountingservice"
	"github.com/nndergunov/deliveryApp/app/services/accounting/pkg/storage/accountingstorage"
)

const configFile = "/config.yaml"

func main() {
	// Construct the application logger.
	l := logger.NewLogger(os.Stdout, "grpc main: ")

	// Perform the startup and shutdown sequence.
	if err := run(l); err != nil {
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

	lis, err := net.Listen("tcp", configreader.GetString("server.grpc.address"))
	if err != nil {
		return fmt.Errorf("failed to listen: %v", err)
	}

	log.Println("starting grpc server", "version", configreader.GetString("buildmode"))
	defer log.Println("shutdown complete")

	dbURL := fmt.Sprintf("host=" + configreader.GetString("database.host") +
		" port=" + configreader.GetString("database.port") +
		" user=" + configreader.GetString("database.user") +
		" password=" + configreader.GetString("database.password") +
		" dbname=" + configreader.GetString("database.dbName") +
		" sslmode=" + configreader.GetString("database.sslmode"))

	database, err := db.OpenDB("postgres", dbURL)
	if err != nil {
		return err
	}

	storage := accountingstorage.NewStorage(accountingstorage.Params{DB: database})

	service := accountingservice.NewService(accountingservice.Params{
		Storage: storage,
		Logger:  logger.NewLogger(os.Stdout, "service: "),
	})

	h := handler.NewHandler(handler.Params{
		Logger:         logger.NewLogger(os.Stdout, "endpoint: "),
		AccountService: service,
	})

	s := grpc.NewServer()
	pb.RegisterAccountingServer(s, h)

	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		return fmt.Errorf("failed to serve: %v", err)
	}
	return nil
}
