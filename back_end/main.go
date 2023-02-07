package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func InitRouter() {
	r := mux.NewRouter()

	// r.HandleFunc("/users", GetUsers).Methods("GET")
	// r.HandleFunc("/users/{id}", GetUser).Methods("GET")
	// r.HandleFunc("/users", CreateUser).Methods("POST")
	// r.HandleFunc("/users/{id}", UpdateUser).Methods("PUT")
	// r.HandleFunc("/users/{id}", DeleteUser).Methods("DELETE")

	r.HandleFunc("/api/status", GetStatus).Methods("GET")
	r.HandleFunc("/database/get", DatabaseGet).Methods("GET")
	r.HandleFunc("/database/post/{ufid}", DatabasePost).Methods("POST")

	log.Fatal(http.ListenAndServe(":9000", r))
}

func main() {
	InitDatabase()
	InitRouter()
}
