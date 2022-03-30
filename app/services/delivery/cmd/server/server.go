package server

import (
	"net/http"

	"github.com/nndergunov/deliveryApp/app/libs/logger"
	"github.com/nndergunov/deliveryApp/app/services/delivery/cmd/server/config"
)

type Server struct {
	HTTPServer *http.Server
	Logger     logger.Logger
}

func NewServer(handler http.Handler, serverConfig config.Config, logger logger.Logger) *Server {
	return &Server{
		HTTPServer: &http.Server{
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
		Logger: logger,
	}
}

func (s *Server) StartListening() {
	err := http.ListenAndServe(s.HTTPServer.Addr, s.HTTPServer.Handler)
	if err != nil {
		s.Log(err)
	}
}

func (s Server) Log(data any) {
	s.Logger.Println(data)
}
