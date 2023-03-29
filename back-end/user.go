package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `json:"username"`
	Password string `json:"password"`
	Name     string `json:"name"`
	//ItineraryList set    `json:"itinerary_list"`
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

func GetUser(w http.ResponseWriter, r *http.Request) {
	SetContentJson(w, r)

	var user User
	username := mux.Vars(r)["username"]
	err := db.First(&user, "username = ?", username).Error
	if err != nil {
		panic("Error retrieving user from database")
	}

	EncodeJson(w, user)
}

func PostUser(w http.ResponseWriter, r *http.Request) {
	SetContentJson(w, r)

	params := mux.Vars(r)
	user := User{Username: params["username"], Password: params["password"], Name: params["name"]}
	DecodeJson(r, user)
	db.Create(&user)

	EncodeJson(w, user)
}
