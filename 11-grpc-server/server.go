package main

import (
	"context"
	"log"
	"net"

	pb "simplegrpcserver/proto/gen"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedCalculateServer
}

func (s *server) Add(ctx context.Context, req *pb.AddRequest) (*pb.AddResponse, error) {
	sum := req.A + req.B
	log.Println("Sum:", sum)
	return &pb.AddResponse{
		Sum: sum,
	}, nil
}
func main() {
	port := ":50051"
	lis, err := net.Listen("tcp", port)

	if err != nil {
		log.Fatal("Failed to listen:", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterCalculateServer(grpcServer, &server{})

	log.Println("Server is running on port", port)
	err = grpcServer.Serve(lis)
	if err != nil {
		log.Fatal("Failed to listen:", err)
	}

}
