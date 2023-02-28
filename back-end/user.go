package main

import (
	//"encoding/json"
	//"fmt"
	//"net/http"

	//_ "modernc.org/sqlite"
	//"gorm.io/driver/sqlite"
	"github.com/glebarez/sqlite"
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

func InitDatabase() {
	// create or open database
	db, err := gorm.Open(sqlite.Open("database.db"), &gorm.Config{})
	if err != nil {
		panic("Cannot connect to DB")
	}

	// migrate struct schema
	db.AutoMigrate(&User{})

	// create database entry
	db.Create(&User{Username: "admin", Password: "password", Name: "Admin Adminton"})

	// read database
	var aUser User
	db.First(&aUser, "name = ?", "Admin Adminton") // find a user with name Admin Adminton

}

/*
func GetUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var users []User

	data, err := db.Query("SELECT * FROM mytable")
	if err != nil {
		fmt.Println(err.Error())
		panic("Cannot select data from DB")
	}

	defer data.Close()

	for data.Next() {
		var user User
		err := data.Scan(&user.UFID, &user.Name)
		if err != nil {
			panic(err.Error())
		}
		users = append(users, user)
	}

	json.NewEncoder(w).Encode(users)
	fmt.Println(users)
}
*/
