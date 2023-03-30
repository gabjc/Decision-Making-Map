package main

/*
import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gorilla/mux"
)

func TestGetItinerary(t *testing.T) {
	InitDB()

	// create test router
	router := mux.NewRouter()
	router.HandleFunc("/itinerary/get/{id}", GetItinerary).Methods("GET")

	// define test request
	req, err := http.NewRequest("GET", "/itinerary/get/admin", nil)
	if err != nil {
		t.Errorf("Error defining request")
	}

	// pass request to router
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// verify that the response body is what we expect
	expected := `{"itin_id": "0","name":"Summer Vacation","address":"123 Summer Rd","radius":"200"}`
	if strings.TrimRight(rr.Body.String(), "\n") != expected {
		t.Errorf("Handler returned unexpected body:\ngot \n%v want \n%v", rr.Body.String(), expected)
	}
}

func TestPostItinerary(t *testing.T) {
	InitDB()

	// create test router
	router := mux.NewRouter()
	router.HandleFunc("/itinerary/post/{name}/{address}/{radius}", PostItinerary).Methods("POST")

	// define test request
	reqBody := strings.NewReader(`{"itin_id": "0", "name": "Summer Vacation", "address": "123 Summer Rd", "radius": "200"}`)
	req, err := http.NewRequest("POST", "/itinerary/post/Summer Vacation/123 Summer Rd/200", reqBody)
	if err != nil {
		t.Errorf("Error defining request")
	}

	// pass request to router
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Hander returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// verify that the new user is in the database
	var itinerary Itinerary
	name := "Summer Vacation"
	err = db.First(&itinerary, "name = ?", name).Error
	if err != nil {
		t.Errorf("Error finding user in database")
	}
}
*/
