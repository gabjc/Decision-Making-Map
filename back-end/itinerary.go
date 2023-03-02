package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// TODO: potentially combine voted locations and saved locations into one struct to migrate to db
type votedLocation struct {
	voteCount   int    `json:"vote"`
	currentVote bool   `json:"currentVote"`
	voteName    string `json:"vote_name"`
}

type Itinerary struct {
	//ItineraryID    string        `json:"id"`
	Name           string         `json:"name"`
	Address        string         `json:"address"`
	Radius         string         `json:"radius"`
	savedLocations map[string]int `json:"saved_locations"`
}

// FIXME: finish implementing, add test file and functions
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
	itinerary := Itinerary{ /*ItineraryID: params["id"],*/ Name: params["name"], Address: params["address"], Radius: params["radius"]} //, savedLocations: params["saved_locations"]
	json.NewDecoder(r.Body).Decode(&itinerary)
	db.Create(&itinerary)

	json.NewEncoder(w).Encode(itinerary)
}
