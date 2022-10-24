package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "github.com/kecci/go-grpc/common/model"
	"google.golang.org/grpc"
)

// userService is a struct implementing UserServiceServer
type userService struct {
	pb.UnimplementedUserServiceServer
}

// GreetUser return greeting message given the name and salutation
// in gRPC protocol
func (*userService) GreetUser(ctx context.Context, req *pb.GreetingRequest) (*pb.GreetingResponse, error) {
	log.Printf("Received: %v %v", req.Salutation, req.Name)

	// business logic
	salutationMessage := fmt.Sprintf("Howdy, %s %s, nice to see you in the future!",
		req.Salutation, req.Name)

	return &pb.GreetingResponse{GreetingMessage: salutationMessage}, nil
}

func main() {
	// Create TCP Server on localhost:5001
	lis, err := net.Listen("tcp", ":5001")
	if err != nil {
		log.Fatalf("Failed to listen on port: %v", err)
	}
	// Create new gRPC server handler
	server := grpc.NewServer()

	// register gRPC UserService to gRPC server handler
	pb.RegisterUserServiceServer(server, &userService{})

	// Run server
	log.Println("Run server", ":5001")
	if err := server.Serve(lis); err != nil {
		log.Fatal(err.Error())
	}
}
