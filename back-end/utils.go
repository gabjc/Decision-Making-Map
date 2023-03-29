package main

import (
	"encoding/json"
	"net/http"
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
	json.NewEncoder(w).Encode(myStruct)
}
func DecodeJson(r *http.Request, myStruct interface{}) {
	json.NewDecoder(r.Body).Decode(&myStruct)
}
