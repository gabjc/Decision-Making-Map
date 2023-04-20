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
	reqBody := strings.NewReader(`{"itin_id": "0", "itin_name": "Testing Itinerary", "saved_plans": "testBreakfast 8:00, testLunch 12:30, testDinner 6:50", "creator_id": "100"}`)
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
	err = db.First(&itinerary, "itin_name", itinName).Error
	if err != nil {
		t.Errorf("Unable to find new itinerary in the database")
	}
}

func TestGetItineraryByID(t *testing.T) {
	InitDB()

	// create test router
	router := mux.NewRouter()
	router.HandleFunc("/itinerary/get", CreateItinerary).Methods("POST")

	// define test request
	reqBody := strings.NewReader(`{"itin_id": "1", "itin_name": "Summer Vacation", "saved_plans": "testBreakfast 8:00, testLunch 12:30, testDinner 6:50", "creator_id": "100"}`)
	req, err := http.NewRequest("POST", "/itinerary/get", reqBody)
	BackendError(err, "Error defining request")

	// pass request to router
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)
	if status := rr.Code; status == http.StatusNotFound || status == http.StatusForbidden {
		t.Errorf("Hander returned wrong status code: got %v want %v", status, http.StatusOK)
	} else {
		// verify that the new itinerary is in the database
		var itinerary Itinerary
		itinId := "0"
		err = db.First(&itinerary, "itin_id", itinId).Error
		if err != nil {
			t.Errorf("Unable to find new itinerary in the database")
		}
	}
}

func TestGetAllItineraries(t *testing.T) {
	InitDB()

	// create test router
	router := mux.NewRouter()
	router.HandleFunc("/itinerary/get", CreateItinerary).Methods("POST")

	// define test request
	reqBody := strings.NewReader(`{"itin_id": "2", "itin_name": "Summer Vacation", "saved_plans": "testBreakfast 8:00, testLunch 12:30, testDinner 6:50", "creator_id": "100"}`)
	req, err := http.NewRequest("POST", "/itinerary/get", reqBody)
	BackendError(err, "Error defining request")
	//Second request
	reqBody = strings.NewReader(`{"itin_id": "3", "itin_name": "Spring Vacation", "saved_plans": "test 12:35, secondTest 12:40", "creator_id": "100"`)
	req, err = http.NewRequest("POST", "/itinerary/get", reqBody)
	BackendError(err, "Error defining request")
	//Third request
	reqBody = strings.NewReader(`{	"itin_id": "4", "itin_name": "Fall Vacation", "saved_plans": "test 10:00, secondTest 1:50", "creator_id": "100"	`)
	req, err = http.NewRequest("POST", "/itinerary/get", reqBody)
	BackendError(err, "Error defining request")

	// pass request to router
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)
	if status := rr.Code; status == http.StatusNotFound || status == http.StatusForbidden {
		t.Errorf("Hander returned wrong status code: got %v want %v", status, http.StatusOK)
	} else {
		// verify that the new itinerary is in the database
		var itineraries []Itinerary
		creatorId := "100"
		err = db.First(&itineraries, "creator_id", creatorId).Error
		if err != nil {
			t.Errorf("Unable to find new itinerary in the database")
		}
	}
}
