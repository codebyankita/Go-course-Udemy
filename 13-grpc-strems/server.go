package main

import (
	"log"
	"net"
	"time"

	mainpb "13-grpc-streams/proto/gen"

	"google.golang.org/grpc"
)

// server implements the Calculator service
type server struct {
	mainpb.UnimplementedCalculatorServer
}

// GenerateFibonacci streams Fibonacci numbers up to n
func (s *server) GenerateFibonacci(req *mainpb.FibonacciRequest, stream mainpb.Calculator_GenerateFibonacciServer) error {
	n := req.N
	a, b := int32(0), int32(1)

	for i := int32(0); i < n; i++ {
		// Send the current Fibonacci number
		err := stream.Send(&mainpb.FibonacciResponse{
			Number: a,
		})
		if err != nil {
			return err
		}
		log.Println("Sent number:", a)
		// Calculate the next Fibonacci number
		a, b = b, a+b

		// Simulate processing delay
		time.Sleep(time.Second)
	}

	return nil
}

func main() {
	// Listen on TCP port 50051
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	// Create a new gRPC server
	grpcServer := grpc.NewServer()

	// Register the Calculator service
	mainpb.RegisterCalculatorServer(grpcServer, &server{})

	// Start serving
	log.Println("Starting gRPC server on :50051")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
