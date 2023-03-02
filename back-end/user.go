package main

import (
	"encoding/json"
	"net/http"

	"github.com/glebarez/sqlite"
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

type SignUpInput struct {
	Username        string `json:"name" binding:"required"`
	Password        string `json:"password" binding:"required,min=8"`
	PasswordConfirm string `json:"passwordConfirm" binding:"required"`
}

type SignInInput struct {
	Username string `json:"username"  binding:"required"`
	Password string `json:"password"  binding:"required"`
}

var db *gorm.DB
var err error

func InitDB() {
	db, err = gorm.Open(sqlite.Open("database.db"), &gorm.Config{})
	if err != nil {
		panic("Cannot connect to DB")
	}
	db.AutoMigrate(&User{})
}

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

/*
// potential code that we can use to see if a user exists at all
func (u User) Exists(id int) bool {
	exists := false
	for _, user := range u {
		if user.ID == id {
			return true
		}
	}
	return exists
}

// FindByName returns the user with the given name, or returns an error
func (u User) FindByName(name string) (User, error) {
	for _, user := range u {
		if user.Name == name {
			return user, nil
		}
	}
	return User{}, errors.New("USER_NOT_FOUND")
}
*/

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
