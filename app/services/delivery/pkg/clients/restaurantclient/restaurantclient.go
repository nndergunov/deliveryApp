package restaurantclient

import (
	"context"
	"fmt"
	"log"

	pb "github.com/nndergunov/deliveryApp/app/services/restaurant/api/v1/grpclogic/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type RestaurantClient struct {
	restaurantURL string
}

func NewRestaurantClient(url string) *RestaurantClient {
	return &RestaurantClient{restaurantURL: url}
}

func (a RestaurantClient) GetRestaurant(restaurantID int) (*pb.RestaurantResponse, error) {
	// Set up a connection to the server.
	conn, err := grpc.Dial(a.restaurantURL, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, fmt.Errorf("did not connect: %v", err)
	}
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			log.Println(err)
		}
	}(conn)
	c := pb.NewRestaurantServiceClient(conn)

	// Contact the server and print out its response.
	ctx := context.TODO()

	r, err := c.GetRestaurant(ctx, &pb.Request{ID: int32(restaurantID)})
	if err != nil {
		return nil, fmt.Errorf("could not get restaurant: %v", err)
	}
	return r, nil
}
