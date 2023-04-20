package main

import (
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB
var err error

func InitDB() {
	db, err = gorm.Open(sqlite.Open("database.db"), &gorm.Config{})
	BackendError(err, "Cannot connect to DB")
	db.AutoMigrate(&User{})
	db.AutoMigrate(&Itinerary{})
}
