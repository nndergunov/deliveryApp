package server

import (
	"context"
	"net/http"

	"github.com/nndergunov/deliveryApp/app/libs/logger"
	"github.com/nndergunov/deliveryApp/app/services/kitchen/cmd/server/config"
)

type Server struct {
	httpServer *http.Server
	logger     *logger.Logger
}

func NewServer(handler http.Handler, serverConfig *config.Config, logger *logger.Logger) *Server {
	return &Server{
		httpServer: &http.Server{
			Addr:              serverConfig.Address,
			Handler:           handler,
			TLSConfig:         nil,
			ReadTimeout:       serverConfig.ReadTimeout,
			ReadHeaderTimeout: serverConfig.ReadHeaderTimeout,
			WriteTimeout:      serverConfig.WriteTimeout,
			IdleTimeout:       serverConfig.IdleTimeout,
			MaxHeaderBytes:    0,
			TLSNextProto:      nil,
			ConnState:         nil,
			ErrorLog:          nil,
			BaseContext:       nil,
			ConnContext:       nil,
		},
		logger: logger,
	}
}

func (s *Server) StartListening(stopChan chan interface{}) {
	go func() {
		if err := s.httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			s.logger.Panicln(err)
		}

		close(stopChan)
	}()

	s.logger.Printf("listening on address %s", s.httpServer.Addr)
}

func (s *Server) Shutdown() {
	if err := s.httpServer.Shutdown(context.TODO()); err != nil {
		s.log(err)
	}
}

func (s Server) log(data any) {
	s.logger.Println(data)
}
