package main

/*
TODO:	*general stuff*
	- encrypting/decrypting passwords
	- conditional compilation and combining frontend and backend when compiling
	- install angular and karma (testing) cli globally => npm install -g @angular/clinpm, npm install -g karma-cli
*/

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func InitRouter() {
	r := mux.NewRouter()

	// TODO: make usernames only take emails, remove non-emails in the db
	r.HandleFunc("/status", GetStatus).Methods("GET")
	r.HandleFunc("/user/get/{username}", GetUser).Methods("GET")
	r.HandleFunc("/user/post/{username}/{password}/{name}", PostUser).Methods("POST")
	r.HandleFunc("/itinerary/get/{id}", GetUser).Methods("GET")
	r.HandleFunc("/itinerary/post/{name}/{address}/{radius}", PostUser).Methods("POST")

	log.Fatal(http.ListenAndServe(":9000", r))
}

func main() {
	InitDB()
	InitRouter()
}
