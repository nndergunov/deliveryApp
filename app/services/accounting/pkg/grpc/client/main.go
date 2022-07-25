// Package main implements a client for Greeter service.
package main

import (
	"context"
	"flag"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/nndergunov/deliveryApp/app/services/accounting/api/v1/grpc/proto"
)

func main() {
	flag.Parse()
	// Set up a connection to the server.
	conn, err := grpc.Dial("localhost:7070", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewAccountingClient(conn)

	// Contact the server and print out its response.
	ctx := context.TODO()

	r, err := c.InsertNewAccount(ctx, &pb.NewAccountRequest{
		UserID:   1,
		UserType: "courier",
	})
	if err != nil {
		log.Fatalf("could not insert: %v", err)
	}

	log.Println(r)
}
