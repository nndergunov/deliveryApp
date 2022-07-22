package handler

import (
	"context"
	"github.com/nndergunov/deliveryApp/app/pkg/logger"

	"github.com/nndergunov/deliveryApp/app/services/consumer/pkg/domain"

	pb "github.com/nndergunov/deliveryApp/app/services/consumer/api/v1/grpc/proto"

	"github.com/nndergunov/deliveryApp/app/services/consumer/pkg/service/consumerservice"
)

type Params struct {
	Logger          *logger.Logger
	ConsumerService consumerservice.ConsumerService
}

// handler is the entrypoint into our application
type handler struct {
	pb.UnsafeConsumerServer
	service consumerservice.ConsumerService
	log     *logger.Logger
}

// NewHandler returns new http multiplexer with configured endpoints.
func NewHandler(p Params) *handler {
	return &handler{
		log:     p.Logger,
		service: p.ConsumerService,
	}
}

func (h *handler) InsertNewConsumer(ctx context.Context, in *pb.NewConsumerRequest) (*pb.ConsumerResponse, error) {
	consumer := domain.Consumer{
		Firstname: in.Firstname,
		Lastname:  in.Lastname,
		Email:     in.Email,
		Phone:     in.Phone,
	}

	resp, err := h.service.InsertConsumer(consumer)
	if err != nil {
		return nil, err
	}

	return &pb.ConsumerResponse{
		ID:        int64(resp.ID),
		Firstname: resp.Firstname,
		Lastname:  resp.Lastname,
		Email:     resp.Email,
		Phone:     resp.Phone,
	}, nil
}

func (h *handler) GetAllConsumer(ctx context.Context, in *pb.SearchParam) (*pb.ConsumerListResponse, error) {
	respList, err := h.service.GetAllConsumer()
	if err != nil {
		return nil, err
	}

	if respList == nil {
		return nil, nil
	}

	var outList []*pb.ConsumerResponse

	for _, resp := range respList {
		out := &pb.ConsumerResponse{
			ID:        int64(resp.ID),
			Firstname: resp.Firstname,
			Lastname:  resp.Lastname,
			Email:     resp.Email,
			Phone:     resp.Phone,
		}
		outList = append(outList, out)
	}
	return &pb.ConsumerListResponse{ConsumerListResponse: outList}, nil
}

func (h *handler) DeleteConsumer(ctx context.Context, id *pb.ConsumerID) (*pb.ConsumerDeleteResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (h *handler) UpdateConsumer(ctx context.Context, request *pb.UpdateConsumerRequest) (*pb.ConsumerResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (h *handler) GetConsumer(ctx context.Context, id *pb.ConsumerID) (*pb.ConsumerResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (h *handler) InsertNewConsumerLocation(ctx context.Context, location *pb.Location) (*pb.Location, error) {
	//TODO implement me
	panic("implement me")
}

func (h *handler) UpdateConsumerLocation(ctx context.Context, location *pb.Location) (*pb.Location, error) {
	//TODO implement me
	panic("implement me")
}

func (h *handler) GetConsumerLocation(ctx context.Context, id *pb.UserID) (*pb.Location, error) {
	//TODO implement me
	panic("implement me")
}
