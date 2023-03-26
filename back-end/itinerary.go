package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

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
	Name           string `json:"name"`
	Address        string `json:"address"`
	Radius         string `json:"radius"`
	SavedLocations string `json:"saved_locations"`
	locationMap    map[string]int
}

func AddSavedLocation(locale string, itin Itinerary) {
	//TODO: parse data from the frontend, then add it to the map, discuss with frontend how they will send the data to the backend
	var location string
	stringLocale := strings.Fields(locale)
	for i, word := range stringLocale {
		if i == 0 {
			location = word
		}
		/*
			else if (i == 1) {

			}
		*/
	}
	_, found := itin.locationMap[location]
	if !found {
		itin.locationMap[location] = 0
	}

}

// Used to increase or decrease the vote for a voted location, or to delete a location from the map if it is 0
func UpdateSavedLocation(locale string, itin Itinerary, update bool) {
	if update {
		itin.locationMap[locale] += 1
	}
	if !update {
		itin.locationMap[locale] -= 1
	}
}

// Used to create a string so you can send it through JSON
func CreateJsonFromMap(data map[string]int) string {
	jsonString, err := json.Marshal(data)
	if err != nil {
		fmt.Printf("Error: %s", err.Error())
	} else {
		return string(jsonString)
	}
	return ""
}

// Potentially used when a person presses a "New Itinerary" button
func createNewItinerary() Itinerary {
	//TODO: save this to a person
	newItin := Itinerary{"Untitled Itinerary", "No Address", "No Radius", "", make(map[string]int)}
	return newItin
}

// FIXME: finish implementing route functions, create unit tests
func GetItinerary(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	w.Header().Set("Content-Type", "application/json")

	itineraryID := mux.Vars(r)["itineraryID"]
	var itinerary Itinerary
	db.First(&itinerary, "itinerary ID = ?", itineraryID)

	json.NewEncoder(w).Encode(itinerary)
}

func PostItinerary(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	itinerary := Itinerary{ /*ItineraryID: params["id"],*/ Name: params["name"], Address: params["address"], Radius: params["radius"], SavedLocations: params["saved_locations"]}
	json.NewDecoder(r.Body).Decode(&itinerary)
	db.Create(&itinerary)

	json.NewEncoder(w).Encode(itinerary)
}
