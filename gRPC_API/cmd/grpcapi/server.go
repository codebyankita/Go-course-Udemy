// package main

// import (
// 	"embed"
// 	"fmt"
// 	"grpcapi/internals/api/handlers"
// 	"grpcapi/internals/api/interceptors"
// 	"grpcapi/pkg/utils"
// 	pb "grpcapi/proto/gen"
// 	"log"
// 	"net"
// 	"os"

// 	"github.com/joho/godotenv"
// 	"google.golang.org/grpc"
// 	"google.golang.org/grpc/reflection"
// )

// //go:embed .env
// var envFile embed.FS

// func loadEnvFromEmbeddedFile() {
// 	content, err := envFile.ReadFile((".env"))
// 	if err != nil {
// 		log.Fatalf("Error reading .env file: %v", err)
// 	}

// 	tempFile, err := os.CreateTemp("", ".env")
// 	if err != nil {
// 		log.Fatalf("Error creating .env file: %v", err)
// 	}
// 	defer os.Remove(tempFile.Name())

// 	_, err = tempFile.Write(content)
// 	if err != nil {
// 		log.Fatalf("Error writing to temp file: %v", err)
// 	}

// 	err = tempFile.Close()
// 	if err != nil {
// 		log.Fatalf("Error writing to temp file: %v", err)
// 	}

// 	err = godotenv.Load(tempFile.Name())
// 	if err != nil {
// 		log.Fatalf("Error loading .env file: %v", err)
// 	}
// }

// func main() {
// 	// err := godotenv.Load()
// 	// if err != nil {
// 	// 	log.Fatal("Error loading .env file:", err)
// 	// }
// 	loadEnvFromEmbeddedFile()

// 	// cert := os.Getenv("CERT_FILE")
// 	// key := os.Getenv("KEY_FILE")

// 	// creds, err := credentials.NewServerTLSFromFile(cert, key)
// 	// if err != nil {
// 	// 	log.Fatalf("Failed to load TLS certificates")
// 	// }

// 	// Not using while benchmarking
// 	// r := interceptors.NewRateLimiter(50, time.Minute)
// 	// s := grpc.NewServer(grpc.ChainUnaryInterceptor(r.RateLimitInterceptor, interceptors.ResponseTimeInterceptor, interceptors.AuthenticationInterceptor), grpc.Creds(creds))

// 	s := grpc.NewServer(grpc.ChainUnaryInterceptor(interceptors.ResponseTimeInterceptor, interceptors.AuthenticationInterceptor))

// 	pb.RegisterExecsServiceServer(s, &handlers.Server{})
// 	pb.RegisterStudentsServiceServer(s, &handlers.Server{})
// 	pb.RegisterTeachersServiceServer(s, &handlers.Server{})

// 	reflection.Register(s)

// 	port := os.Getenv("SERVER_PORT")

// 	go utils.JwtStore.CleanUpExpiredTokens()

// 	fmt.Println("gRPC Server is running on port:", port)
// 	lis, err := net.Listen("tcp", port)
// 	if err != nil {
// 		log.Fatal("Error listening on specified port:", err)
// 	}

//		err = s.Serve(lis)
//		if err != nil {
//			log.Fatal("Failed to serve:", err)
//		}
//	}
package main

import (
	"fmt"
	"grpcapi/internals/api/handlers"
	// "grpcapi/internals/repositories/mongodb"
	pb "grpcapi/proto/gen"
	"log"
	"net"
	"os"

	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	// mongodb.CreateMongoClient()
	// Load environment variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file: ", err)
	}

	// Create new gRPC server
	s := grpc.NewServer()

	// Register services
	pb.RegisterExecsServiceServer(s, &handlers.Server{})
	pb.RegisterStudentsServiceServer(s, &handlers.Server{})
	pb.RegisterTeachersServiceServer(s, &handlers.Server{})

	// Enable reflection (useful for debugging with grpcurl)
	reflection.Register(s)

	// Get server port from env
	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = ":50051" // fallback default
	}

	fmt.Println("🚀 gRPC Server is running on port", port)

	// Start listening
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal("❌ Error listening on specified port:", err)
	}

	// Serve gRPC
	if err := s.Serve(lis); err != nil {
		log.Fatal("❌ Failed to serve:", err)
	}
}
