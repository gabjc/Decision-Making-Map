package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func InitRouter() {
	r := mux.NewRouter()

	r.HandleFunc("/api/status", GetStatus).Methods("GET")
	r.HandleFunc("/database/get", DatabaseGet).Methods("GET")
	r.HandleFunc("/database/post/{ufid}", DatabasePost).Methods("POST")

	log.Fatal(http.ListenAndServe(":9000", r))
}

func main() {
	InitDatabase()
	InitRouter()
}
