package main

import (
	"context"
	"io"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	mainpb "14-grpc-streams-client/proto/gen"
)

func main() {
	// Establish a connection to the gRPC server
	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect to server: %v", err)
	}
	defer conn.Close()

	// Create a client for the Calculator service
	client := mainpb.NewCalculatorClient(conn)

	// Create a context
	ctx := context.Background()

	// Create a request for the Fibonacci stream with N=10
	req := &mainpb.FibonacciRequest{
		N: 10,
	}

	// Call the GenerateFibonacci streaming RPC
	stream, err := client.GenerateFibonacci(ctx, req)
	if err != nil {
		log.Fatalf("Error calling GenerateFibonacci: %v", err)
	}

	// Receive and process the stream of Fibonacci numbers
	for {
		resp, err := stream.Recv()
		if err == io.EOF {
			log.Println("End of Fibonacci stream")
			break
		}
		if err != nil {
			log.Fatalf("Error receiving data from GenerateFibonacci: %v", err)
		}
		log.Printf("Fibonacci number: %d", resp.GetNumber())
	}
}