package main

import (
	"encoding/json"
	"html"
	"net/http"
	"strings"

	"github.com/joshual55/Decision-Making-Map/vendor/github.com/gorilla/mux"
)

type Itinerary struct {
	ItinID     uint64 `gorm:"primary_key;auto_increment" json:"itin_id"`
	ItinName   string `gorm:"size:255;not null;unique" json:"itin_name"`
	SavedPlans string `gorm:"not null" json:"saved_plans"`
	Creator    User   `json:"creator"`
	CreatorID  uint32 `sql:"type:int REFERENCES users(id)" json:"creator_id"`
}

func (i *Itinerary) Prepare() {
	i.ItinID = 0
	i.ItinName = html.EscapeString(strings.TrimSpace(i.ItinName))
	i.SavedPlans = html.EscapeString(strings.TrimSpace(i.SavedPlans))
	i.Creator = User{}
}

func CreateItinerary(w http.ResponseWriter, r *http.Request) {
	SetContentJson(w, r)

	// pull itin info into struct from request and basic information
	var itin Itinerary
	itin.Prepare()

	// create Itinerary in database
	db.Create(&itin)

	w.WriteHeader(http.StatusCreated)
}

// TODO: finish implementing route functions, create unit tests
func GetItineraryByID(w http.ResponseWriter, r *http.Request) {
	SetContentJson(w, r)

	var itinerary Itinerary
	err := db.First(&itinerary, "itinID = ?", itinerary.ItinID).Error

	if err != nil { //no itinerary found with that id, return 404
		w.WriteHeader(http.StatusNotFound)
	} else {
		json.NewEncoder(w).Encode(itinerary)
	}
}

func GetAllItineraries(w http.ResponseWriter, r *http.Request) {
	SetContentJson(w, r)

	creatorID := mux.Vars(r)["id"]
	var itineraries []Itinerary
	err := db.Where("CreatorID = ?", creatorID).Find(&itineraries).Error

	if err != nil || len(itineraries) <= 0 {
		w.WriteHeader(http.StatusNotFound)
	} else {
		json.NewEncoder(w).Encode(itineraries)
	}
}

/*
ignore for now
func (i *Itinerary) _PostItinerary(db *gorm.DB) (*Itinerary, error) {
	var err error
	err = db.Debug().Model(&Itinerary{}).Create(&i).Error
	if err != nil {
		return &Itinerary{}, err
	}
	if i.ItinID != 0 {
		err = db.Debug().Model(&User{}).Where("id = ?", i.CreatorID).Take(&i.Creator).Error
		if err != nil {
			return &Itinerary{}, err
		}
	}
	return i, nil
}

// Adds location to map
func AddSavedLocation(locale string, itin Itinerary) {
	var location string
	stringLocale := strings.Fields(locale)
	for i, word := range stringLocale {
		if i == 0 {
			location = word
		}
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
*/

/*
func PostItinerary(w http.ResponseWriter, r *http.Request) {
	SetContentJson(w, r)

	params := mux.Vars(r)
	itinerary := Itinerary{ItineraryID: params["itin_id"], Name: params["name"], Address: params["address"], Radius: params["radius"], SavedLocations: params["saved_locations"]}
	json.NewDecoder(r.Body).Decode(&itinerary)
	db.Create(&itinerary)

	json.NewEncoder(w).Encode(itinerary)
}
*/
