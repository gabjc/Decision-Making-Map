package main

import (
	"encoding/json"
	"net/http"

	//"github.com/glebarez/sqlite"
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
/*
type SignUpInput struct {
	Username        string `json:"name" binding:"required"`
	Password        string `json:"password" binding:"required,min=8"`
	PasswordConfirm string `json:"passwordConfirm" binding:"required"`
}

type SignInInput struct {
	Username string `json:"username"  binding:"required"`
	Password string `json:"password"  binding:"required"`
}
*/

// TODO: maybe use array of itineraries instead of set
// funcs for editing a set of itineraries within a user struct
/*
var exists = struct{}{}

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
}
*/

// TODO: add more routes and more specific routes
func GetUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	username := mux.Vars(r)["username"]
	var user User
	db.First(&user, "username = ?", username)

	json.NewEncoder(w).Encode(user)
}

func PostUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	user := User{Username: params["username"], Password: params["password"], Name: params["name"]}
	json.NewDecoder(r.Body).Decode(&user)
	db.Create(&user)

	json.NewEncoder(w).Encode(user)
}
