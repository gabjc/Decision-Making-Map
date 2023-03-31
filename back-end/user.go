package main

import (
	"net/http"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name             string `json:"name"`
	Email            string `json:"email"`
	Hash             []byte `json:"hash"`
	OwnedItineraries string `json:"owned_itineraries"`
	//itinerariesMap   map[bool]int
}

// Old route function
/* func GetUser(w http.ResponseWriter, r *http.Request) {
	SetContentJson(w, r)

	// search db for user w/ username
	var user User
	username := mux.Vars(r)["username"]
	err := db.First(&user, "username = ?", username).Error
	CheckError(err, "Error retrieving user from database")

	// return user JSON
	EncodeJson(w, user)
} */

func RegisterUser(w http.ResponseWriter, r *http.Request) {
	SetContentJson(w, r) // TODO: is this needed if we take, but don't return, a JSON?

	// pull user info into struct from request
	var user User
	DecodeJson(r, &user)

	// set user vars not included in register request
	user.OwnedItineraries = ""

	// convert password to hash
	password := user.Hash
	hash, err := HashPassword(string(password))
	CheckError(err, "Error hashing password")

	// replace string password with hash in user struct
	user.Hash = hash

	// create user in database
	db.Create(&user)

	w.WriteHeader(http.StatusCreated)
}

func Login(w http.ResponseWriter, r *http.Request) {
	SetContentJson(w, r)

	// pull user info into struct from request
	var user User
	DecodeJson(r, &user)
	password := string(user.Hash)

	// find the user's hash in the database
	var user2 User
	err := db.First(&user2, "email = ?", user.Email).Error

	if err != nil { // no user, return 404 Not Found
		w.WriteHeader(http.StatusNotFound)
	} else if !HashMatches(password, user2.Hash) { // hash doesn't match, return 403 Forbidden
		w.WriteHeader(http.StatusForbidden)
	} else { // else return 201 Created and JWT token
		w.WriteHeader(http.StatusCreated)
		// TODO: send token
	}
}
