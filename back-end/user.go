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
