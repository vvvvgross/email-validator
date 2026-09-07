package main

import (
	"log"
	"net/http"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/vvvvgross/email-validator/api/validator_v1"
	"github.com/vvvvgross/email-validator/internal/gateway"
)

func main() {
	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect client with producer: %v", err)
	}

	defer conn.Close()

	grpcClient := pb.NewEmailValidationClient(conn)

	myHandler := &gateway.Handler{ValidatorClient: grpcClient}

	mux := http.NewServeMux()
	mux.HandleFunc("POST /api/v1/validate", myHandler.ValidateEmail)

	log.Println("Starting API Gateway on port :8080...")

	if err = http.ListenAndServe(":8080", mux); err != nil {
		log.Fatalf("Failed to listen and serve: %v", err)
	}
}
