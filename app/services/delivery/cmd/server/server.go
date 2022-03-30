package server

import (
	"log"
	"net/http"

	"github.com/nndergunov/deliveryApp/app/services/delivery/cmd/server/config"
)

type Server struct {
	HTTPServer *http.Server
}

func NewServer(handler http.Handler, serverConfig config.Config, errorLog *log.Logger) *Server {
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
			ErrorLog:          errorLog,
			BaseContext:       nil,
			ConnContext:       nil,
		},
	}
}

func (s *Server) StartListening() {
	err := http.ListenAndServe(s.HTTPServer.Addr, s.HTTPServer.Handler)
	if err != nil {
		s.Log(err)
	}
}

func (s Server) Log(data any) {
	s.HTTPServer.ErrorLog.Println(data)
}
