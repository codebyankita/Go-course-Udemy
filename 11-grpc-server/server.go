package main

import (
	"context"
	"log"
	"net"

	pb "simplegrpcserver/proto/gen"
	farewellpb "simplegrpcserver/proto/gen/farewell"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

type server struct {
	pb.UnimplementedCalculateServer
	pb.UnimplementedGreeterServer
	farewellpb.UnimplementedAufWiedersehenServer
}

func (s *server) Add(ctx context.Context, req *pb.AddRequest) (*pb.AddResponse, error) {
	sum := req.A + req.B
	log.Println("Sum:", sum)
	return &pb.AddResponse{
		Sum: sum,
	}, nil
}
func main() {
	cert := "cert.pem"
	key := "key.pem"
	port := ":50051"
	lis, err := net.Listen("tcp", port)

	if err != nil {
		log.Fatal("Failed to listen:", err)
	}
	creds, err := credentials.NewServerTLSFromFile(cert, key)

	if err != nil {
		log.Fatal("Failed to load creadientials", err)

	}
	grpcServer := grpc.NewServer(grpc.Creds(creds))
	pb.RegisterCalculateServer(grpcServer, &server{})
	pb.RegisterGreeterServer(grpcServer, &server{})
	 farewellpb.RegisterAufWiedersehenServer(grpcServer, &server{})
	// pb.RegisterBidFarewellServer(grpcServer, &server{})

	log.Println("Server is running on port", port)
	err = grpcServer.Serve(lis)
	if err != nil {
		log.Fatal("Failed to listen:", err)
	}

}
