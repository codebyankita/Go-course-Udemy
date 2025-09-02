package main

import (
	"context"
	"log"
	"time"

	mainapipb "simplegrpcclient/proto/gen"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func main() {
	// Load TLS certificate
	creds, err := credentials.NewClientTLSFromFile("cert.pem", "")
	if err != nil {
		log.Fatalln("Failed to load TLS cert:", err)
	}

	// Dial gRPC server with TLS
	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(creds))
	if err != nil {
		log.Fatalln("Failed to connect:", err)
	}
	defer conn.Close()

	// Create client
	client := mainapipb.NewCalculateClient(conn)

	// Set context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// Build request
	req := &mainapipb.AddRequest{
		A: 10,
		B: 20,
	}

	// Call RPC
	res, err := client.Add(ctx, req)
	if err != nil {
		log.Fatalln("Could not add:", err)
	}

	log.Println("Sum:", res.Sum)
}
