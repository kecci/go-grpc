package main

import (
	"context"
	"fmt"
	"log"
	"time"

	pb "github.com/kecci/go-grpc/common/model"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const userServiceAddress = "localhost:5001"

func main() {
	// Create connection to gRPC server
	conn, err := grpc.Dial(userServiceAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Could not connect to service: %v", err)
		return
	}
	defer conn.Close()

	// Create new userService client
	userServiceClient := pb.NewUserServiceClient(conn)

	// create connection timeout
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	response, err := userServiceClient.GreetUser(ctx, &pb.GreetingRequest{
		Salutation: "Mr.",
		Name:       "Angelo",
	})
	if err != nil {
		log.Fatalf("Could not create request: %v", err)
	}

	// show response
	fmt.Println(response.GreetingMessage)
}
