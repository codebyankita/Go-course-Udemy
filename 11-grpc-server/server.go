package main

import (
	"log"
	"net"

	"google.golang.org/grpc"
)
func main(){
	port:= ":50051"
	lis, err := net.Listen("tcp",port)

	if err != nil {
		log.Fatal("Failed to listen:", err)}
grpcserver := grpc.NewServer()

log.Println("Server is running on port", port)
err = grpcserver.Serve(lis)
if err != nil {
		log.Fatal("Failed to listen:", err)}
		
}