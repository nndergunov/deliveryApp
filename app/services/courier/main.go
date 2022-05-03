package main

//go:generate sqlboiler postgres
import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"courier/conf"
	"courier/database"
	"courier/handler"
	"courier/service"
	"github.com/gorilla/mux"
	"github.com/nndergunov/deliveryApp/app/pkg/configreader"
	"github.com/nndergunov/deliveryApp/app/pkg/logger"
)

func main() {
	// Construct the application logger.
	log := logger.NewLogger(os.Stdout, "carrier-api")

	// Perform the startup and shutdown sequence.
	if err := run(log); err != nil {
		log.Fatal("startup", "ERROR", err)
	}
}

func run(log *logger.Logger) error {
	dbConf, err := conf.GetConf("DB.dev")
	if err != nil {
		return err
	}

	db, err := database.Open(dbConf)
	if err != nil {
		return err
	}
	defer db.Close()

	log.Println("starting service", "version", configreader.GetString("buildmode"))
	defer log.Println("shutdown complete")

	newCarrierService, err := service.NewCourierService(service.Params{
		DB:     db,
		Logger: log,
	})
	if err != nil {
		return err
	}

	// Make a channel to listen for an interrupt or terminate signal from the OS.
	// Use a buffered channel because the signal package requires it.
	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, syscall.SIGINT, syscall.SIGTERM)

	// Construct a server to service the requests against the mux.
	app := handler.NewCarrierHandler(handler.Params{
		Logger:         log,
		CourierService: newCarrierService,
		Srv:            mux.NewRouter(),
		Shutdown:       shutdown,
	})

	api := http.Server{
		Addr:         "localhost:7070",
		Handler:      app,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
	}

	// Make a channel to listen for errors coming from the listener. Use a
	// buffered channel so the goroutine can exit if we don't collect this error.
	serverErrors := make(chan error, 1)

	// Start the service listening for requests.
	go func() {
		log.Printf("main : API listening on %s", api.Addr)
		serverErrors <- api.ListenAndServe()
	}()

	// Blocking main and waiting for shutdown.
	select {
	case err := <-serverErrors:
		log.Fatalf("error: listening and serving: %s", err)

	case <-shutdown:
		log.Printf("main : Start shutdown")

		// Give outstanding requests a deadline for completion.
		const timeout = 5 * time.Second
		ctx, cancel := context.WithTimeout(context.Background(), timeout)
		defer cancel()

		// Asking listener to shutdown and load shed.
		err := api.Shutdown(ctx)
		if err != nil {
			log.Printf("main : Graceful shutdown did not complete in %v : %v", timeout, err)
			err = api.Close()
		}

		if err != nil {
			log.Fatalf("main : could not stop server gracefully : %v", err)
		}
	}
	return nil
}
