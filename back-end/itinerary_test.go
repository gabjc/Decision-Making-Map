package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gorilla/mux"
)

func TestCreateItinerary(t *testing.T) {
	InitDB()

	// create test router
	router := mux.NewRouter()
	router.HandleFunc("/itinerary/home", CreateItinerary).Methods("POST")

	// define test request
	reqBody := strings.NewReader(`{"itin_id": 0, "itin_name": "Testing Itinerary", "saved_plans": "testBreakfast 8:00, testLunch 12:30, testDinner 6:50", "creator_id": 100}`)
	req, err := http.NewRequest("POST", "/itinerary/home", reqBody)
	BackendError(err, "Error Defining Request")

	// pass request to router
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusCreated)
	}

	// verify that the new itinerary is in the database
	var itinerary Itinerary
	itinName := "Testing Itinerary"
	err = db.First(&itinerary, "itin_name = ?", itinName).Error
	if err != nil {
		t.Errorf("Unable to find new itinerary in the database")
	}
}

// TODO: implement route for getting by itin ID, not all itins of a user ID, then implement this
/*
func TestGetItineraryByID(t *testing.T) {
	InitDB()
	// create test router
	router := mux.NewRouter()
	router.HandleFunc("/itinerary/get/{id}", GetItineraryByID).Methods("GET")
	// define test request
	req, err := http.NewRequest("GET", "/itinerary/get/33", nil)
	BackendError(err, "Error defining request")
	// pass request to router
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)
	if status := rr.Code; status == http.StatusNotFound {
		t.Errorf("Hander returned wrong status code: got %v want %v", status, http.StatusOK)
	} else {
		// check that result matches expected
		expected := `[{
			"itin_id": 33,
			"itin_name": "Testing Itinerary",
			"saved_plans": "testBreakfast 8:00, testLunch 12:30, testDinner 6:50",
			"creator": {
				"ID": 0,
				"CreatedAt": "0001-01-01T00:00:00Z",
				"UpdatedAt": "0001-01-01T00:00:00Z",
				"DeletedAt": null,
				"id": 0,
				"name": "",
				"email": "",
				"hash": null,
				"owned_itineraries": ""
			},
			"creator_id": 101
		]`

		if rr.Body.String() != expected {
			t.Errorf("Itineraries not found")
		}
	}
}
*/

func TestGetAllItineraries(t *testing.T) {
	InitDB()

	// create test router
	router := mux.NewRouter()
	router.HandleFunc("/itinerary/get/{id}", GetAllItineraries).Methods("GET")

	// define test request
	req, err := http.NewRequest("GET", "/itinerary/get/101", nil)
	BackendError(err, "Error defining request")

	// pass request to router
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)
	if status := rr.Code; status == http.StatusNotFound {
		t.Errorf("Hander returned wrong status code: got %v want %v", status, http.StatusOK)
	} else {
		// check that result matches expected
		expected := `[
			{
				"itin_id": 33,
				"itin_name": "Testing Itinerary",
				"saved_plans": "testBreakfast 8:00, testLunch 12:30, testDinner 6:50",
				"creator": {
					"ID": 0,
					"CreatedAt": "0001-01-01T00:00:00Z",
					"UpdatedAt": "0001-01-01T00:00:00Z",
					"DeletedAt": null,
					"id": 0,
					"name": "",
					"email": "",
					"hash": null,
					"owned_itineraries": ""
				},
				"creator_id": 101
			},
			{
				"itin_id": 34,
				"itin_name": "Testing Itinerary",
				"saved_plans": "testBreakfast 8:00, testLunch 12:30, testDinner 6:50",
				"creator": {
					"ID": 0,
					"CreatedAt": "0001-01-01T00:00:00Z",
					"UpdatedAt": "0001-01-01T00:00:00Z",
					"DeletedAt": null,
					"id": 0,
					"name": "",
					"email": "",
					"hash": null,
					"owned_itineraries": ""
				},
				"creator_id": 101
			}
		]`
		result := rr.Body.String()
		if result != expected { // TODO: figure out null terminator issue
			t.Errorf("Itineraries not found")
		}
	}
}
