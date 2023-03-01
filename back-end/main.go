package main

/*
TODO:
	- encrypting/decrypting passwords
	- renaming this module with github/gabjc instead of github/joshual55
	- conditional compilation and combining frontend and backend when compiling
*/

// go unit tests
// install angular and karma (testing) cli globally => npm install -g @angular/clinpm, npm install -g karma-cli
// status function

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func InitRouter() {
	r := mux.NewRouter()

	r.HandleFunc("/status", GetStatus).Methods("GET")
	r.HandleFunc("/user/get/{username}", GetUser).Methods("GET")
	r.HandleFunc("/user/post/{username}/{password}/{name}", PostUser).Methods("POST")

	log.Fatal(http.ListenAndServe(":9000", r))
}

func main() {
	InitDB()
	InitRouter()
}
