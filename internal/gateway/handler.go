package gateway

import (
	"encoding/json"
	"log"
	"net/http"

	pb "github.com/vvvvgross/email-validator/api/validator_v1"
)

type Handler struct {
	ValidatorClient pb.EmailValidationClient
}

func (h *Handler) ValidateEmail(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	type Request struct {
		Email string `json:"email"`
	}

	req := &Request{}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		return
	}

	resp, err := h.ValidatorClient.Validate(r.Context(), &pb.ValidateRequest{Email: req.Email})
	if err != nil {
		http.Error(w, "Failed to validate email via gRPC", http.StatusInternalServerError)
		log.Printf("gRPC call failed: %v", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err = json.NewEncoder(w).Encode(resp); err != nil {
		log.Printf("Failed to encode response: %v", err)
	}
}
