package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	_ "modernc.org/sqlite"
	//"gorm.io/gorm"
	//"github.com/glebarez/sqlite"
	// NOT USING GIN FOR HTTP?
)

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

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "http://localhost:4200") // IS THIS EVEN NEEDED?
}

func GetStatus(w http.ResponseWriter, r *http.Request) { // FINISH IMPLEMENTING THIS & UN-COMMENT STATUS FETCH IN app.component.ts
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
	enableCors(&w)
	w.Header().Set("Content-Type", "application/json")
	data, err := db.Query("SELECT * FROM mytable")
	if err != nil {
		fmt.Println(err.Error())
		panic("Cannot select data from DB")
	}
	json.NewEncoder(w).Encode(data) // RETURN MESSES UP, CURRENT DATA TYPE IS *sql.Rows
}

func DatabasePost(w http.ResponseWriter, r *http.Request) {
	// IMPLEMENT CREATING A USER WITH A NAME AND UFID
}

/*
func GetUsers(w http.ResponseWriter, r *http.Request) {

}

func GetUser(w http.ResponseWriter, r *http.Request) {

}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user User
	json.NewDecoder(r.Body).Decode(&user)
	db.Create(&user)
	json.NewEncoder(w).Encode(user)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {

}

func DeleteUser(w http.ResponseWriter, r *http.Request) {

}
*/
