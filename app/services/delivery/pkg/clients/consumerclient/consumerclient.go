package consumerclient

import (
	"context"
	"fmt"
	"log"

	pb "github.com/nndergunov/deliveryApp/app/services/consumer/api/v1/grpc/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ConsumerClient struct {
	consumerURL string
}

func NewConsumerClient(url string) *ConsumerClient {
	return &ConsumerClient{consumerURL: url}
}

func (a ConsumerClient) GetLocation(consumerID int64) (*pb.Location, error) {
	// Set up a connection to the server.
	conn, err := grpc.Dial(a.consumerURL, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, fmt.Errorf("did not connect: %v", err)
	}
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			log.Println(err)
		}
	}(conn)

	c := pb.NewConsumerClient(conn)

	// Contact the server and print out its response.
	ctx := context.TODO()

	r, err := c.GetConsumerLocation(ctx, &pb.UserID{UserID: consumerID})
	if err != nil {
		return nil, fmt.Errorf("could not get locaiton: %v", err)
	}
	return r, nil
}
