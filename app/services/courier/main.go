package main

import (
	"context"
	"courier/internal/database"
	"courier/internal/database/storage"
	"courier/internal/handlers/courierhandler"
	"os"
	"os/signal"
	"syscall"
	"time"

	"courier/app"
	"courier/conf"
	"courier/internal/services"
	"github.com/nndergunov/deliveryApp/app/pkg/configreader"
	"github.com/nndergunov/deliveryApp/app/pkg/logger"
)

func main() {
	// Construct the application logger.
	log := logger.NewLogger(os.Stdout, "courier-api")

	// Perform the startup and shutdown sequence.
	if err := run(log); err != nil {
		log.Fatal("startup", "ERROR", err)
	}
}

func run(log *logger.Logger) error {
	if err := conf.SetConfPath(); err != nil {
		return err
	}

	db, err := database.Open(configreader.GetString("DB.dev"))
	if err != nil {
		return err
	}
	defer db.Close()

	log.Println("starting services", "version", configreader.GetString("buildmode"))
	defer log.Println("shutdown complete")

	newCourierStorage, err := storage.NewCourierStorage(storage.Params{
		DB: db,
	})
	if err != nil {
		return err
	}

	courierService, err := services.NewCourierService(services.Params{
		CourierStorage: newCourierStorage,
		Logger:         logger.NewLogger(os.Stdout, "courier-service"),
	})
	if err != nil {
		return err
	}

	// Make a channel to listen for an interrupt or terminate signal from the OS.
	// Use a buffered channel because the signal package requires it.
	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, syscall.SIGINT, syscall.SIGTERM)

	router, server, err := app.NewHandlerServer(app.Params{
		Logger:   log,
		Shutdown: shutdown,
	})

	// Construct a server to services the requests against the mux.
	courierhandler.NewCourierHandler(courierhandler.Params{
		Logger:         logger.NewLogger(os.Stdout, "courier-handler"),
		CourierService: courierService,
		Route:          router,
		Shutdown:       shutdown,
	})

	// Make a channel to listen for errors coming from the listener. Use a
	// buffered channel so the goroutine can exit if we don't collect this error.
	serverErrors := make(chan error, 1)

	// Start the services listening for requests.
	go func() {
		log.Printf("main : API listening on %s", server.Addr)
		serverErrors <- server.ListenAndServe()
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
		err := server.Shutdown(ctx)
		if err != nil {
			log.Printf("main : Graceful shutdown did not complete in %v : %v", timeout, err)
			err = server.Close()
		}

		if err != nil {
			log.Fatalf("main : could not stop server gracefully : %v", err)
		}
	}
	return nil
}
