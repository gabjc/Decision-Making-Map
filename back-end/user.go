package main

import (
	"net/http"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID               uint32 `json:"id"`
	Name             string `json:"name"`
	Email            string `json:"email"`
	Hash             []byte `json:"hash"`
	OwnedItineraries string `json:"owned_itineraries"`
	//itinerariesMap   map[bool]int
}

// TODO: make sure user does not already exist, add frontend handling as well
func RegisterUser(w http.ResponseWriter, r *http.Request) {
	SetContentJson(w, r)

	// pull user info into struct from request
	var user User
	DecodeJson(r, &user)

	// convert password to hash
	password := user.Hash
	hash, err := HashPassword(string(password))
	BackendError(err, "Error hashing password")

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
	} else { // else create JWT token, return it and 201 Created
		tok, err := CreateToken(user2.ID)
		BackendError(err, "Error creating JWT")
		EncodeJson(w, http.StatusCreated, tok)
	}
}
