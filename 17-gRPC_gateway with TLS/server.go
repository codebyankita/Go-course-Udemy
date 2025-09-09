package main

import (
	"context"
	"crypto/tls"
	"log"
	"net"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"

	pb "grpcgatewayproject/proto"
)

type server struct {
	pb.UnimplementedGreeterServer
}

func (s *server) Greet(ctx context.Context, req *pb.HelloRequest) (*pb.HelloResponse, error) {
	if err := req.Validate(); err != nil {
		log.Printf("Validation failed: %v", err)
		return nil, status.Errorf(codes.InvalidArgument, "Invalid request: %v", err)
	}
	return &pb.HelloResponse{Message: "Hello, " + req.GetName()}, nil
}

func runGRPCServer(certFile, keyFile string) {
	// Load TLS certificates
	creds, err := credentials.NewServerTLSFromFile(certFile, keyFile)
	if err != nil {
		log.Fatalf("Failed to load TLS certificates: %v", err)
	}

	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer(grpc.Creds(creds))
	pb.RegisterGreeterServer(grpcServer, &server{})

	reflection.Register(grpcServer)
	log.Println("gRPC server is running on port :50051...")

	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}

func runGatewayServer(certFile, keyFile string) {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(credentials.NewTLS(&tls.Config{
		InsecureSkipVerify: true,
	}))}

	err := pb.RegisterGreeterHandlerFromEndpoint(ctx, mux, "localhost:50051", opts)
	if err != nil {
		log.Fatalf("Failed to register gRPC-Gateway handler: %v", err)
	}

	tlsConfig := &tls.Config{
		Certificates: []tls.Certificate{loadTLSCredentials(certFile, keyFile)},
	}

	server := &http.Server{
		Addr:      ":8080",
		Handler:   mux,
		TLSConfig: tlsConfig,
	}

	log.Println("HTTP server is running on port :8080...")
	if err := server.ListenAndServeTLS(certFile, keyFile); err != nil {
		log.Fatalf("Failed to serve HTTP: %v", err)
	}

}

func loadTLSCredentials(certFile, keyFile string) tls.Certificate {
	// Load the certificates from disk
	cert, err := tls.LoadX509KeyPair(certFile, keyFile)
	if err != nil {
		log.Fatalf("Failed to load key pair: %s", err)
	}
	return cert
}

func main() {
	certFile := "cert.pem"
	keyFile := "key.pem"
	go runGRPCServer(certFile, keyFile)
	runGatewayServer(certFile, keyFile)
}
