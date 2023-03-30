package main

import (
	"encoding/json"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

// cors
func EnableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

// http routes
func SetContentJson(w http.ResponseWriter, r *http.Request) {
	EnableCors(&w)
	w.Header().Set("Content-Type", "application/json")
}
func EncodeJson(w http.ResponseWriter, myStruct interface{}) {
	err := json.NewEncoder(w).Encode(myStruct)
	CheckError(err, "Error encoding JSON")
}
func DecodeJson(r *http.Request, myStruct interface{}) {
	err := json.NewDecoder(r.Body).Decode(&myStruct)
	CheckError(err, "Error decoding JSON")
}

// errors
func CheckError(err error, msg string) {
	if err != nil {
		panic(msg)
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
