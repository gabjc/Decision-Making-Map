package main

/*
TODO:	*general stuff*
	- learn more about Cypress testing
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

	// TODO: make post route only take emails for usernames, remove non-emails in the db
	// TODO: figure out password encryption
	r.HandleFunc("/status", GetStatus).Methods("GET")
	r.HandleFunc("/user/get/{username}", GetUser).Methods("GET")
	r.HandleFunc("/user/post/{username}/{password}/{name}", PostUser).Methods("POST")
	r.HandleFunc("/itinerary/get/{id}", GetUser).Methods("GET")
	r.HandleFunc("/itinerary/post/{name}/{address}/{radius}", PostUser).Methods("POST")
	r.HandleFunc("/refresh", Refresh)
	r.HandleFunc("/logout", Logout)

	log.Fatal(http.ListenAndServe(":9000", r))
}

func main() {
	InitDB()
	InitRouter()
}
