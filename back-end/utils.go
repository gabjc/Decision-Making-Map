package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

// http routes
func SetContentJson(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
}
func DecodeJson(r *http.Request, myStruct interface{}) {
	err := json.NewDecoder(r.Body).Decode(&myStruct)
	BackendError(err, "Error decoding JSON")
}
func EncodeJson(w http.ResponseWriter, statusCode int, myStruct interface{}) {
	w.WriteHeader(statusCode)
	err := json.NewEncoder(w).Encode(myStruct)
	BackendError(err, "Error encoding JSON")
}

// error checking
func BackendError(err error, msg string) {
	if err != nil {
		panic(msg)
	}
}
func FrontendError(w http.ResponseWriter, err error, statusCode int) {
	if err != nil {
		EncodeJson(w, statusCode, struct {
			Error string `json:"error"`
		}{
			Error: err.Error(),
		})
		// print error info on backend
		fmt.Println("Unauthorized request")
	}
}

// bcrypt
func HashPassword(password string) ([]byte, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return hash, err
}
func HashMatches(password string, hash []byte) bool {
	err := bcrypt.CompareHashAndPassword(hash, []byte(password))
	return err == nil
}

// JWT tokens
func ConfirmAuthentication(w http.ResponseWriter, r *http.Request) {
	err := TokenValid(r)
	FrontendError(w, err, http.StatusUnauthorized)
}
