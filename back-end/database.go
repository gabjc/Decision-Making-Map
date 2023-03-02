package main

import (
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB
var err error

func InitDB() {
	db, err = gorm.Open(sqlite.Open("database.db"), &gorm.Config{})
	if err != nil {
		panic("Cannot connect to DB")
	}
	db.AutoMigrate(&User{}) // TODO: add struct schema for itineraries
}
