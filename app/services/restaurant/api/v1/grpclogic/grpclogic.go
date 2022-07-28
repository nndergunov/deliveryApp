// Package grpclogic implements a server using gRPC connection.
package grpclogic

import (
	"context"
	"fmt"

	"google.golang.org/grpc"

	"github.com/nndergunov/deliveryApp/app/services/restaurant/api/v1/grpclogic/pb"
	"github.com/nndergunov/deliveryApp/app/services/restaurant/pkg/service"
)

type restaurantGRPC struct {
	pb.UnimplementedRestaurantServiceServer
	service service.AppService
}

// NewRestaurantRawGRPCServer returns new instance of RestaurantGRPCServer.
func NewRestaurantRawGRPCServer(service service.AppService) *grpc.Server {
	grpcLogic := new(restaurantGRPC)

	grpcLogic.service = service

	srv := grpc.NewServer()

	pb.RegisterRestaurantServiceServer(srv, grpcLogic)

	return srv
}

// GetRestaurant handles request of getting the restaurant from the service.
func (r *restaurantGRPC) GetRestaurant(ctx context.Context, req *pb.Request) (*pb.RestaurantResponse, error) {
	rest, err := r.service.ReturnRestaurant(int(req.GetID()))
	if err != nil {
		return nil, fmt.Errorf("getting restaurant: %w", err)
	}

	return &pb.RestaurantResponse{
		ID:              int32(rest.ID),
		Name:            rest.Name,
		AcceptingOrders: rest.AcceptingOrders,
		City:            rest.City,
		Address:         rest.Address,
		Longitude:       float32(rest.Longitude),
		Latitude:        float32(rest.Latitude),
	}, nil
}

// GetMenu handles request of getting the restaurant menu from the service.
func (r *restaurantGRPC) GetMenu(ctx context.Context, req *pb.Request) (*pb.MenuResponse, error) {
	menu, err := r.service.ReturnMenu(int(req.GetID()))
	if err != nil {
		return nil, fmt.Errorf("getting menu: %w", err)
	}

	returnMenuItems := make([]*pb.ReturnMenuItem, len(menu.Items))

	for ind, item := range menu.Items {
		returnMenuItems[ind] = &pb.ReturnMenuItem{
			ID:     int32(item.ID),
			Name:   item.Name,
			Price:  float32(item.Price),
			Course: item.Course,
		}
	}

	return &pb.MenuResponse{
		RestaurantID: int32(menu.RestaurantID),
		MenuItems:    returnMenuItems,
	}, nil
}
