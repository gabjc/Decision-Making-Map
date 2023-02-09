package main

import (
	"database/sql"
	"fmt"

	//"github.com/gin-gonic/gin"
	_ "modernc.org/sqlite"
)

var db *sql.DB
var err error

func InitDatabase() { // FIGURE OUT BETTER ORGANIZATION FOR CREATING DB WITHOUT COMMENTING OUT SECTIONS
	// create or open database
	db, err = sql.Open("sqlite", "database.db")
	if err != nil {
		fmt.Println(err.Error())
		panic("Cannot connect to DB")
	}
	// create table
	/*_, err = db.Exec("CREATE TABLE mytable (UFID int, Name varchar(255))")
	if err != nil {
		fmt.Println(err.Error())
		panic("Cannot create table in DB")
	}*/
	// insert example data
	/*_, err = db.Exec("INSERT INTO mytable (UFID, Name) VALUES (20099125, 'Joshua Lamb')")
	if err != nil {
		fmt.Println(err.Error())
		panic("Cannot insert into DB")
	}*/
}
