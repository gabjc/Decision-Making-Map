package main

import (
	//"github.com/gin-gonic/gin"

	"encoding/json"
	"math/rand"
	"net/http"
	"time"

	"github.com/glebarez/sqlite"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

type votedLocation struct {
	voteCount   int    `json:"vote"`
	currentVote bool   `json:"currentVote"`
	voteName    string `json:"vote_name"`
}

type Itinerary struct {
	ItineraryID    string         `json:"id"`
	Name           string         `json:"name"`
	Address        string         `json:"address"`
	Radius         string         `json:"radius"`
	savedLocations map[string]int `json:"saved_locations"`
}

var dbItin *gorm.DB
var errItin error

func InitLocationDB() {
	dbItin, errItin = gorm.Open(sqlite.Open("database.db"), &gorm.Config{})
	if err != nil {
		panic("Cannot connect to DB")
	}
	db.AutoMigrate(&User{})
}

/*
func updateVotedLocation() {
	var savedLocations []votedLocation
	for _, savedLocations := range savedLocations {
		if savedLocations.voteName == {
			if savedLocations.currentVote == true {
				savedLocations.voteCount++;
			}
			if savedLocations.currentVote == false {
				savedLocations.voteCount--;
			}
		}

	itinerary Itinerary
	itinerary.savedLocations
}
*/

// for randomizing IDs
func init() {
	rand.Seed(time.Now().UnixNano())
}

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")

func randSeq(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func GetItinerary(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	itineraryID := mux.Vars(r)["itineraryID"]
	var itinerary Itinerary
	db.First(&itinerary, "itinerary ID = ?", itineraryID)

	json.NewEncoder(w).Encode(itinerary)
}

func PostItinerary(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	itinerary := Itinerary{ItineraryID: params["id"], Name: params["name"], Address: params["address"], Radius: params["radius"]} //, savedLocations: params["saved_locations"]
	json.NewDecoder(r.Body).Decode(&itinerary)
	db.Create(&itinerary)

	json.NewEncoder(w).Encode(itinerary)
}
