package function

import (
	"encoding/json"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
)

const (
	pwdMinLen = 8
)

// DataResponse represents a standard data frame for returning lists.
// Not returning lists as top level object, but passing lists in data: json field for
// easier parsing
type DataResponse struct {
	Data interface{} `json:"data"`
}

// FailureResponse represents an errors with an associated HTTP status code.
type FailureResponse struct {
	Error   string                 `json:"errors"`
	Unknown map[string]interface{} `json:"-"` // Rest of the fields should go here.
}

// EncryptRequest - POST Request Body
type EncryptRequest struct {
	Encrypt string
	Cost    int
}

// HashAndSalt - Hash and salt password using bcrypt
func HashAndSalt(pwd string, cost int) (string, error) {
	// Use GenerateFromPassword to hash & salt pwd.
	// MinCost is just an integer constant provided by the bcrypt
	// package along with DefaultCost & MaxCost.
	// The cost can be any value you want provided it isn't lower
	// than the MinCost (4)
	if len(pwd) < pwdMinLen {
		return "", fmt.Errorf("password provided needs to be at least %d characters long", pwdMinLen)
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(pwd), cost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
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
	var req EncryptRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		log.Println("JSON Decode error: ", err)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(FailureResponse{Error: "JSON Decode error"})
		return
	}

	// Validate defaults and encrypt
	if req.Cost == 0 {
		req.Cost = bcrypt.DefaultCost
	}
	log.Printf("%v", req)
	bc, err := HashAndSalt(req.Encrypt, req.Cost)
	if err != nil {
		log.Println("Encrypt error: ", err)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(FailureResponse{Error: "Encrypt error"})
		return
	}

	json.NewEncoder(w).Encode(DataResponse{Data: bc})
}
