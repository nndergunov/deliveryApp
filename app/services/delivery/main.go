package main

import (
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	"net/http"
	"os"

	"github.com/nndergunov/deliveryApp/app/services/delivery/api/v1/grpc/handler"

	"github.com/gorilla/handlers"

	pb "github.com/nndergunov/deliveryApp/app/services/delivery/api/v1/grpc/proto"
	"github.com/nndergunov/deliveryApp/app/services/delivery/api/v1/rest/handler/deliveryhandler"
	"github.com/nndergunov/deliveryApp/app/services/delivery/pkg/clients/restaurantclient"

	"github.com/nndergunov/deliveryApp/app/services/delivery/pkg/clients/consumerclient"

	"github.com/nndergunov/deliveryApp/app/pkg/api"
	"github.com/nndergunov/deliveryApp/app/pkg/configreader"
	"github.com/nndergunov/deliveryApp/app/pkg/logger"
	"github.com/nndergunov/deliveryApp/app/pkg/server/config"

	"github.com/nndergunov/deliveryApp/app/services/delivery/pkg/clients/courierclient"
	"github.com/nndergunov/deliveryApp/app/services/delivery/pkg/db"
	"github.com/nndergunov/deliveryApp/app/services/delivery/pkg/service/deliveryservice"
	"github.com/nndergunov/deliveryApp/app/services/delivery/pkg/storage/deliverystorage"
)

const configFile = "/config.yaml"

func main() {
	// Construct the application logger.
	l := logger.NewLogger(os.Stdout, "main: ")

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

	defer log.Println("rest shutdown complete")
	defer log.Println("grpc shutdown complete")

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
	DeliveryStorage := deliverystorage.NewDeliveryStorage(deliverystorage.Params{DB: database})

	courierClient := courierclient.NewCourierClient(configreader.GetString("courierServiceURl"))
	restaurantClient := restaurantclient.NewRestaurantClient(configreader.GetString("restaurantServiceURl"))
	consumerClient := consumerclient.NewConsumerClient(configreader.GetString("consumerServiceURl"))

	deliveryService := deliveryservice.NewDeliveryService(deliveryservice.Params{
		DeliveryStorage:  DeliveryStorage,
		Logger:           logger.NewLogger(os.Stdout, "service: "),
		CourierClient:    courierClient,
		RestaurantClient: restaurantClient,
		ConsumerClient:   consumerClient,
	})

	//gRPC server
	lis, err := net.Listen("tcp", configreader.GetString("server.grpc.address"))
	if err != nil {
		return fmt.Errorf("failed to listen: %v", err)
	}

	h := handler.NewHandler(handler.Params{
		Logger:          logger.NewLogger(os.Stdout, "endpoint: "),
		DeliveryService: deliveryService,
	})

	s := grpc.NewServer()
	pb.RegisterDeliveryServer(s, h)

	go func() {
		if err := s.Serve(lis); err != nil {
			log.Panicln("grpc failed to serve: %v", err)
		}
	}()

	log.Printf("started grpc service on address:%v version:%v", lis.Addr(), configreader.GetString("buildmode"))

	//REST server
	restHandler := deliveryhandler.NewDeliveryHandler(deliveryhandler.Params{
		Logger:          logger.NewLogger(os.Stdout, "endpoint: "),
		DeliveryService: deliveryService,
	})

	apiLogger := logger.NewLogger(os.Stdout, "api: ")
	serverAPI := api.NewAPI(restHandler, apiLogger)

	serverLogger := logger.NewLogger(os.Stdout, "server: ")
	serverConfig := getServerConfig(serverAPI, nil, serverLogger)

	serverErrors := make(chan interface{})

	// Where ORIGIN_ALLOWED is like `scheme://dns[:port]`, or `*` (insecure)
	headersOK := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type"})
	originsOK := handlers.AllowedOrigins([]string{"*"})
	methodsOK := handlers.AllowedMethods([]string{"GET", "POST", "OPTIONS", "DELETE", "PUT"})

	go func() {
		if err := http.ListenAndServe(serverConfig.Address, handlers.CORS(headersOK, originsOK, methodsOK)(restHandler)); err != nil {
			log.Panicln(err)
		}
		close(serverErrors)
	}()
	log.Printf("started rest service on port:%v version:%v", serverConfig.Address, configreader.GetString("buildmode"))

	<-serverErrors

	return nil
}

func getServerConfig(handler http.Handler, errorLog *log.Logger, serverLogger *logger.Logger) *config.Config {
	var (
		address          = configreader.GetString("server.rest.address")
		readTime         = configreader.GetDuration("server.rest.readTime")
		writeTime        = configreader.GetDuration("server.rest.writeTime")
		idleTime         = configreader.GetDuration("server.rest.idleTime")
		readerHeaderTime = configreader.GetDuration("server.rest.readerHeaderTime")
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
