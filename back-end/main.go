package main

/*
TODO:
	- go unit tests
	- encrypting/decrypting passwords
	- conditional compilation and combining frontend and backend when compiling
*/

// for now do gorm, proper user struct for storing login into, and go unit tests for the routes and stuff
// and look into install angular and karma (testing) cli globally => npm install -g @angular/clinpm, npm install -g karma-cli
// and maybe the status command and sending status codes so they can receive on the frontend

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func InitRouter() {
	r := mux.NewRouter()

	r.HandleFunc("/status", GetStatus).Methods("GET")
	//r.HandleFunc("/database/get", GetUser).Methods("GET")
	//r.HandleFunc("/database/post/{ufid}/{name}", PostUser).Methods("POST")

	log.Fatal(http.ListenAndServe(":9000", r))
}

func main() {
	InitDatabase()
	InitRouter()
}
