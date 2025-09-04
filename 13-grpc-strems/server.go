package main

import (
	"log"
	"net"
	"time"
	"io"
	mainpb "13-grpc-streams/proto/gen"
	"bufio"
	"fmt"
	"os"
	"strings"
	
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
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
func (s *server) SendNumbers(stream mainpb.Calculator_SendNumbersServer) error {
	var sum int32
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&mainpb.NumberResponse{Sum: sum})
		}
		if err != nil {
			return err
		}
		log.Println(req.GetNumber())
		sum += req.GetNumber()
	}
}
func (s *server) Chat(stream mainpb.Calculator_ChatServer) error {
	reader := bufio.NewReader(os.Stdin)

	for {
		// Receiving values/messages from stream
		req, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalln(err)
		}
		log.Println("Received Message:", req.GetMessage())

		// Read input from the terminal
		fmt.Print("Enter response: ")
		input, err := reader.ReadString('\n')
		if err != nil {
			return err
		}

		input = strings.TrimSpace(input)

		response := &mainpb.ChatMessage{
			Message: input,
		}
		// Sending data/messages/values through the stream
		err = stream.Send(response)
		if err != nil {
			return err
		}
	}
	fmt.Println("Returning control")
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
// Enable reflection
	reflection.Register(grpcServer)
	err = grpcServer.Serve(lis)
	if err != nil {
		log.Fatalln(err)
	}
	// Start serving
	log.Println("Starting gRPC server on :50051")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
