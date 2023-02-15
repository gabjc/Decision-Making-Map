package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	//"github.com/gin-gonic/gin"
	_ "modernc.org/sqlite"
)

type User struct {
	UFID int    `json:"ufid"`
	Name string `json:"name"`
}

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

func PostUser(w http.ResponseWriter, r *http.Request) {
	// IMPLEMENT CREATING A USER WITH A NAME AND UFID
}
