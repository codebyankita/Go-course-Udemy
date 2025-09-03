package main

import (
	"context"
	"log"
	"time"

	mainapipb "simplegrpcclient/proto/gen"
	farewellpb "simplegrpcclient/proto/gen/farewell"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func main() {
	// cert := "cert.pem"
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

	// client2 := mainapipb.NewGreeterClient(conn)
	fwClient := farewellpb.NewAufWiedersehenClient(conn)

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
	// reqGreet := &mainapipb.HelloRequest{
	// 	Name: "John",
	// }
	// res1, err := client2.Greet(ctx, reqGreet)
	// if err != nil {
	// 	log.Fatalln("Could not greet", err)
	// }
	// resAddFromGreeter, err := client2.Add(ctx, reqGreet)
	// if err != nil {
	// 	log.Println("Could not add-------", err)
	// }
	reqGoodBye := &farewellpb.GoodByeRequest{
		Name: "Jane",
	}
	resFw, err := fwClient.BidGoodBye(ctx, reqGoodBye)
	if err != nil {
		log.Fatalln("Could not bid Goodbye", err)
	}
	log.Println("Sum:", res.Sum)
	// log.Println("Greeting message:", res1.Message)
	log.Println("Goodbye message:", resFw.Message)
	state := conn.GetState()
	log.Println("Connection State:", state)
}
