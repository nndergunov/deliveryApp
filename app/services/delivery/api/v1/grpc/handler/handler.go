package handler

import (
	"context"
	"strconv"

	"google.golang.org/grpc"

	"github.com/nndergunov/deliveryApp/app/pkg/logger"

	"github.com/nndergunov/deliveryApp/app/services/delivery/pkg/domain"

	pb "github.com/nndergunov/deliveryApp/app/services/delivery/api/v1/grpc/proto"
	"github.com/nndergunov/deliveryApp/app/services/delivery/pkg/service/deliveryservice"
)

type Params struct {
	Logger  *logger.Logger
	Service deliveryservice.DeliveryService
}

// handler is the entrypoint into our application
type handler struct {
	pb.UnsafeDeliveryServer
	service deliveryservice.DeliveryService
	log     *logger.Logger
}

// NewHandler returns new http multiplexer with configured endpoints.
func NewHandler(p Params) *grpc.Server {
	h := &handler{
		log:     p.Logger,
		service: p.Service,
	}

	srv := grpc.NewServer()
	pb.RegisterDeliveryServer(srv, h)

	return srv
}

func (h handler) GetEstimateDeliveryValues(ctx context.Context, in *pb.EstimateDeliveryRequest) (*pb.EstimateDeliveryResponse, error) {
	out, err := h.service.GetEstimateDelivery(strconv.FormatInt(in.ConsumerID, 10), strconv.FormatInt(in.RestaurantID, 10))
	if err != nil {
		return nil, err
	}
	return &pb.EstimateDeliveryResponse{
		Time: out.Time,
		Cost: out.Cost,
	}, nil
}

func (h handler) AssignOrder(ctx context.Context, in *pb.AssignOrderRequest) (*pb.AssignOrderResponse, error) {
	order := &domain.Order{
		OrderID:          int(in.OrderID),
		FromUserID:       int(in.FromUserID),
		FromRestaurantID: int(in.RestaurantID),
	}

	out, err := h.service.AssignOrder(strconv.Itoa(order.OrderID), order)
	if err != nil {
		return nil, err
	}
	return &pb.AssignOrderResponse{
		OrderID:   int64(out.OrderID),
		CourierID: int64(out.CourierID),
	}, nil
}
