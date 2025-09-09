package main

import (
	"context"
	"log"
	"net"

	pb "grpcprotovalidate/proto/gen"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
)

// server is used to implement mainapi.GreeterServer.
type server struct {
	pb.UnimplementedGreeterServer
}

// Greet implements mainapi.GreeterServer
func (s *server) Greet(ctx context.Context, req *pb.HelloRequest) (*pb.HelloResponse, error) {
	// Validate the incoming request
	err := req.Validate()
	if err != nil {
		log.Printf("Validation failed: %v", err)
		// Return an invalid argument error if validation fails
		return nil, status.Errorf(codes.InvalidArgument, "Invalid request: %v", err)
	}

	return &pb.HelloResponse{Message: "Hello, " + req.GetName()}, nil
}

func main() {
	// Create a listener on TCP port 50051
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	// Create a gRPC server
	grpcServer := grpc.NewServer()
	pb.RegisterGreeterServer(grpcServer, &server{})

	reflection.Register(grpcServer)

	log.Println("Server is running on port :50051...")
	// Start the server
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
