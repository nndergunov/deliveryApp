package handler

import (
	"context"
	"strconv"

	"google.golang.org/grpc"

	"github.com/nndergunov/deliveryApp/app/pkg/logger"

	"github.com/nndergunov/deliveryApp/app/services/consumer/pkg/domain"

	pb "github.com/nndergunov/deliveryApp/app/services/consumer/api/v1/grpc/proto"

	"github.com/nndergunov/deliveryApp/app/services/consumer/pkg/service/consumerservice"
)

type Params struct {
	Logger  *logger.Logger
	Service consumerservice.ConsumerService
}

// handler is the entrypoint into our application
type handler struct {
	pb.UnsafeConsumerServer
	service consumerservice.ConsumerService
	log     *logger.Logger
}

// NewHandler returns new http multiplexer with configured endpoints.
func NewHandler(p Params) *grpc.Server {
	h := &handler{
		log:     p.Logger,
		service: p.Service,
	}

	srv := grpc.NewServer()
	pb.RegisterConsumerServer(srv, h)

	return srv
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

func (h *handler) DeleteConsumer(ctx context.Context, in *pb.ConsumerID) (*pb.ConsumerDeleteResponse, error) {
	out, err := h.service.DeleteConsumer(strconv.Itoa(int(in.ConsumerID)))
	if err != nil {
		return nil, err
	}

	return &pb.ConsumerDeleteResponse{ConsumerDeleteResponse: out}, nil
}

func (h *handler) UpdateConsumer(ctx context.Context, in *pb.UpdateConsumerRequest) (*pb.ConsumerResponse, error) {
	consumer := domain.Consumer{
		Firstname: in.Firstname,
		Lastname:  in.Lastname,
		Email:     in.Email,
		Phone:     in.Phone,
	}

	resp, err := h.service.UpdateConsumer(consumer, strconv.FormatInt(in.ID, 10))
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

func (h *handler) GetConsumer(ctx context.Context, in *pb.ConsumerID) (*pb.ConsumerResponse, error) {
	resp, err := h.service.GetConsumer(strconv.FormatInt(in.ConsumerID, 10))
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

func (h *handler) InsertNewConsumerLocation(ctx context.Context, in *pb.Location) (*pb.Location, error) {
	location := domain.Location{
		UserID:     int(in.UserID),
		Latitude:   *in.Latitude,
		Longitude:  *in.Longitude,
		Country:    *in.Country,
		City:       *in.City,
		Region:     *in.Region,
		Street:     *in.Street,
		HomeNumber: *in.HomeNumber,
		Floor:      *in.Floor,
		Door:       *in.Door,
	}

	resp, err := h.service.InsertLocation(location, strconv.Itoa(location.UserID))
	if err != nil {
		return nil, err
	}

	return &pb.Location{
		UserID:     int64(resp.UserID),
		Latitude:   &resp.Latitude,
		Longitude:  &resp.Longitude,
		Country:    &resp.Country,
		City:       &resp.City,
		Region:     &resp.Region,
		Street:     &resp.Street,
		HomeNumber: &resp.HomeNumber,
		Floor:      &resp.Floor,
		Door:       &resp.Door,
	}, nil
}

func (h *handler) UpdateConsumerLocation(ctx context.Context, in *pb.Location) (*pb.Location, error) {
	location := domain.Location{
		UserID:     int(in.UserID),
		Latitude:   *in.Latitude,
		Longitude:  *in.Longitude,
		Country:    *in.Country,
		City:       *in.City,
		Region:     *in.Region,
		Street:     *in.Street,
		HomeNumber: *in.HomeNumber,
		Floor:      *in.Floor,
		Door:       *in.Door,
	}

	resp, err := h.service.UpdateLocation(location, strconv.Itoa(location.UserID))
	if err != nil {
		return nil, err
	}

	return &pb.Location{
		UserID:     int64(resp.UserID),
		Latitude:   &resp.Latitude,
		Longitude:  &resp.Longitude,
		Country:    &resp.Country,
		City:       &resp.City,
		Region:     &resp.Region,
		Street:     &resp.Street,
		HomeNumber: &resp.HomeNumber,
		Floor:      &resp.Floor,
		Door:       &resp.Door,
	}, nil
}

func (h *handler) GetConsumerLocation(ctx context.Context, in *pb.UserID) (*pb.Location, error) {
	resp, err := h.service.GetLocation(strconv.FormatInt(in.UserID, 10))
	if err != nil {
		return nil, err
	}

	return &pb.Location{
		UserID:     int64(resp.UserID),
		Latitude:   &resp.Latitude,
		Longitude:  &resp.Longitude,
		Country:    &resp.Country,
		City:       &resp.City,
		Region:     &resp.Region,
		Street:     &resp.Street,
		HomeNumber: &resp.HomeNumber,
		Floor:      &resp.Floor,
		Door:       &resp.Door,
	}, nil
}
