package main

/*
TODO:	*general stuff*
	- authentication (I am who I say I am) vs authorization (knowing who I am, what can I do?)
		- cookies
		- private-public key cryptography
		- password encryption (see bcrypt library)
			- creating -> take password, generate extra parts like seeds, compute hash with bcrypt, put hash and seed in db with gorm
			- checking -> retrieve hash and seed, feed into bcrypt, test if password matches, give or don't give the cookie

	- continuous integration (versus c-delivery and c-deployment)
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

	r.HandleFunc("/status", GetStatus).Methods("GET")
	r.HandleFunc("/user/register", RegisterUser).Methods("POST")
	r.HandleFunc("/user/login", Login).Methods("POST")

	log.Fatal(http.ListenAndServe(":9000", r))
}

func main() {
	InitDB()
	InitRouter()
}
