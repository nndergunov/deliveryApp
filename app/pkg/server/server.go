package server

import (
	"context"
	"errors"
	"net/http"

	"github.com/nndergunov/deliveryApp/app/pkg/logger"
	"github.com/nndergunov/deliveryApp/app/pkg/server/config"
)

// Server defines an http server.
type Server struct {
	httpServer *http.Server
	logger     *logger.Logger
}

// NewServer returns new instance of server.Server with specified configuration.
func NewServer(serverConfig *config.Config) *Server {
	return &Server{
		httpServer: &http.Server{
			Addr:              serverConfig.Address,
			Handler:           serverConfig.Handler,
			TLSConfig:         nil,
			ReadTimeout:       serverConfig.ReadTimeout,
			ReadHeaderTimeout: serverConfig.ReadHeaderTimeout,
			WriteTimeout:      serverConfig.WriteTimeout,
			IdleTimeout:       serverConfig.IdleTimeout,
			MaxHeaderBytes:    0,
			TLSNextProto:      nil,
			ConnState:         nil,
			ErrorLog:          serverConfig.ErrorLog,
			BaseContext:       nil,
			ConnContext:       nil,
		},
		logger: serverConfig.ServerLogger,
	}
}

// StartListening runs server.Server.
func (s *Server) StartListening(stopChan chan interface{}) {
	go func() {
		if err := s.httpServer.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			s.logger.Panicln(err)
		}

		close(stopChan)
	}()

	s.logger.Printf("\nlistening on address %s", s.httpServer.Addr)
}

// Shutdown gracefully stops server.Server.
func (s *Server) Shutdown() {
	if err := s.httpServer.Shutdown(context.TODO()); err != nil {
		s.logger.Println(err)
	}
}
