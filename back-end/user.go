package main

import (
	"net/http"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email string `json:"email"`
	Hash  []byte `json:"hash"`
	Name  string `json:"name"`
	//ItineraryList set    `json:"itinerary_list"`
	OwnedItineraries string `json:"owned_itineraries"`
	itinerariesMap   map[bool]int
}

// TODO: potentially use these structs for user login and signup
/* type SignUpInput struct {
	Username        string `json:"name" binding:"required"`
	Password        string `json:"password" binding:"required,min=8"`
	PasswordConfirm string `json:"passwordConfirm" binding:"required"`
}

type SignInInput struct {
	Username string `json:"username"  binding:"required"`
	Password string `json:"password"  binding:"required"`
} */

// TODO: maybe use array of itineraries instead of set
// funcs for editing a set of itineraries within a user struct
/* var exists = struct{}{}

// a set to find the itineraries that a user holds
type set struct {
	m map[string]struct{}
}

func NewSet() *set {
	s := &set{}
	s.m = make(map[string]struct{})
	return s
}

func (s *set) Add(value string) {
	s.m[value] = exists
}

func (s *set) Remove(value string) {
	delete(s.m, value)
}

func (s *set) Contains(value string) bool {
	_, c := s.m[value]
	return c
} */

/*
/
func GetUser(w http.ResponseWriter, r *http.Request) {
	SetContentJson(w, r)

	// search db for user w/ username
	var user User
	username := mux.Vars(r)["username"]
	err := db.First(&user, "username = ?", username).Error
	CheckError(err, "Error retrieving user from database")

	EncodeJson(w, user)
}
*/

func RegisterUser(w http.ResponseWriter, r *http.Request) {
	SetContentJson(w, r)

	// pull user info into struct from request
	var user User
	DecodeJson(r, user)

	// convert string password to hash
	password := string(user.Hash)
	hash, err := HashPassword(password)
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
	DecodeJson(r, user)
	password := string(user.Hash)

	// find the user's hash in the database
	var user2 User
	err := db.First(&user2, "email = ?", user.Email).Error

	if err != nil { // if no user, return 404 Not Found
		w.WriteHeader(http.StatusNotFound)
	} else if !HashMatches(password, user2.Hash) { // else if hash doesn't match, return 403 Forbidden
		w.WriteHeader(http.StatusForbidden)
	} else { // else return 201 status Created and JWT token
		w.WriteHeader(http.StatusCreated)
		// TODO: send token
	}
}
