package function

import (
	"encoding/json"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
)

// FailureResponse represents an errors with an associated HTTP status code.
type FailureResponse struct {
	Error   string                 `json:"errors"`
	Unknown map[string]interface{} `json:"-"` // Rest of the fields should go here.
}

// ValidateRequest - POST Request Body
type ValidateRequest struct {
	Text string
	Hash string
}

// ComparePasswords - Validate password and hash
// Returns error if not match
func ComparePasswords(plainPwd, hashedPwd string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPwd), []byte(plainPwd))
}

func Handle(w http.ResponseWriter, r *http.Request) {
	// Make sure it is post
	if r.Method != http.MethodPost {
		log.Println("Only POST allowed, found: ", r.Method, " on: "+r.RequestURI)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(FailureResponse{Error: "only POST allowed for this endpoint"})
		return
	}

	// Decode JSON
	var req ValidateRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		log.Println("JSON Decode error: ", err)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(FailureResponse{Error: "JSON Decode error"})
		return
	}

	// Validate
	err = ComparePasswords(req.Text, req.Hash)
	if err != nil {
		log.Println("Validation error: ", err)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(FailureResponse{Error: "Validation error. Not a match!"})
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
