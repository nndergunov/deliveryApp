package grpcserver

import (
	"net"

	"github.com/nndergunov/deliveryApp/app/pkg/logger"
	"google.golang.org/grpc"
)

// GRPCServer defines a gRPC server.
type GRPCServer struct {
	srv    *grpc.Server
	logger *logger.Logger
}

// NewGRPCServer returns new instance of GRPCServer with specified configuration.
func NewGRPCServer(srv *grpc.Server, log *logger.Logger) *GRPCServer {
	return &GRPCServer{
		srv:    srv,
		logger: log,
	}
}

// StartListening runs gRPC server.
func (s *GRPCServer) StartListening(addr string, stopChan chan interface{}) {
	go func() {
		listener, err := net.Listen("tcp", addr)
		if err != nil {
			s.logger.Fatalf("failed to listen on address %s: %v", addr, err)
		}

		if err := s.srv.Serve(listener); err != nil {
			s.logger.Fatalf("failed to serve: %v", err)
		}

		close(stopChan)
	}()

	s.logger.Printf("\nlistening on address %s", addr)
}

// Shutdown gracefully stops gRPC server.
func (s *GRPCServer) Shutdown() {
	s.srv.GracefulStop()
}
