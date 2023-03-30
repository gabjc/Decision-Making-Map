package main

import (
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB
var err error

func InitDB() {
	db, err = gorm.Open(sqlite.Open("database.db"), &gorm.Config{})
	CheckError(err, "Cannot connect to DB")
	db.AutoMigrate(&User{}) // TODO: migrate struct for itineraries once complete
}
