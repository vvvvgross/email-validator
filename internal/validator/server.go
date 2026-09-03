package validator

import (
	"context"
	"strings"

	pb "github.com/vvvvgross/email-validator/api/validator_v1"
)

type Server struct {
	pb.UnimplementedEmailValidationServer
}

func (s *Server) Validate(ctx context.Context, req *pb.ValidateRequest) (*pb.ValidateResponse, error) {
	email := req.GetEmail()

	if len(email) <= 3 {
		return &pb.ValidateResponse{
			IsCorrect: false,
			Answer:    "Email is too short",
		}, nil
	}

	if !strings.Contains(email, "@") {
		return &pb.ValidateResponse{
			IsCorrect: false,
			Answer:    "Email doesn't contain symbol '@'",
		}, nil
	}

	return &pb.ValidateResponse{
		IsCorrect: true,
		Answer:    "Email is correct",
	}, nil
}
