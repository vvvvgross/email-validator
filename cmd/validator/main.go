package main

import (
	"log"
	"net"

	"google.golang.org/grpc"

	pb "github.com/vvvvgross/email-validator/api/validator_v1"
	"github.com/vvvvgross/email-validator/internal/validator"
)

func main() {
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to start listener: %v", err)
	}

	grpcServer := grpc.NewServer()
	myServer := &validator.Server{}
	pb.RegisterEmailValidationServer(grpcServer, myServer)

	log.Println("Starting gRPC server on port :50051...")

	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
