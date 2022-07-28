package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"

	pb "github.com/nndergunov/deliveryApp/app/services/consumer/api/v1/grpc/proto"
)

func main() {
	location, err := GetLocation(1)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(location)
}
func GetLocation(consumerID int64) (*pb.Location, error) {
	// Set up a connection to the server.
	conn, err := grpc.Dial(":441", grpc.WithTransportCredentials(insecure.NewCredentials()))
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
