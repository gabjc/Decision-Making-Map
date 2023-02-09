package main

import (
	"database/sql"
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

var db *sql.DB
var err error

func InitDatabase() {
	// create or open database
	db, err = sql.Open("sqlite", "mydata.db")
	if err != nil {
		fmt.Println(err.Error())
		panic("Cannot connect to DB")
	}
	// create table
	/* _, err = db.Exec("CREATE TABLE mytable (UFID int, Name varchar(255))")
	if err != nil {
		fmt.Println(err.Error())
		panic("Cannot create table in DB")
	} */
	// insert example data
	/* _, err = db.Exec("INSERT INTO mytable (UFID, Name) VALUES (20099125, 'Joshua Lamb')")
	if err != nil {
		fmt.Println(err.Error())
		panic("Cannot insert into DB")
	} */
}

func GetStatus(w http.ResponseWriter, r *http.Request) { // FINISH IMPLEMENTING, UN-COMMENT STATUS FETCH IN app.component.ts
	/* w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	resp := make(map[string]string)
	resp["message"] = "status: UP"
	jsonResp, err := json.Marshal(resp)
	if err != nil {
		fmt.Println(err.Error())
		panic("Cannot convert status OK to JSON")
	}
	w.Write(jsonResp)
	return */
}

func DatabaseGet(w http.ResponseWriter, r *http.Request) {
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

	json.NewEncoder(w).Encode(users) // RETURN MESSES UP, CURRENT DATA TYPE IS *sql.Rows
	fmt.Println(users)
}

func DatabasePost(w http.ResponseWriter, r *http.Request) {
	// IMPLEMENT CREATING A USER WITH A NAME AND UFID
}
