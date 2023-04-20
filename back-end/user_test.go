package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gorilla/mux"
)

// TODO: abstract testing funcs by adding more helper funcs like for errors
// TODO: can also do Postman tests
func TestRegisterUser(t *testing.T) {
	InitDB()

	// create test router
	router := mux.NewRouter()
	router.HandleFunc("/user/register", RegisterUser).Methods("POST")

	// define test request
	reqBody := strings.NewReader(`{"name": "Testy Man", "email": "test@test.com", "hash": "password", "owned_itineraries": ""}`)
	req, err := http.NewRequest("POST", "/user/register", reqBody)
	BackendError(err, "Error defining request")

	// pass request to router
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("Hander returned wrong status code: got %v want %v", status, http.StatusCreated)
	}

	// verify that the new user is in the database
	var user User
	email := "test@test.com"
	err = db.First(&user, "email = ?", email).Error
	if err != nil {
		t.Errorf("Unable to find new user in database")
	}
}

func TestLoginUser(t *testing.T) {
	InitDB()

	// create test router
	router := mux.NewRouter()
	router.HandleFunc("/user/login", RegisterUser).Methods("POST")

	// define test request
	reqBody := strings.NewReader(`{"email": "test", "hash": "testpassword"}`)
	req, err := http.NewRequest("POST", "/user/login", reqBody)
	BackendError(err, "Error defining request")

	// pass request to router
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)
	if status := rr.Code; status == http.StatusNotFound || status == http.StatusForbidden {
		t.Errorf("Hander returned wrong status code: got %v want %v", status, http.StatusOK)
	} else {
		// TODO: somehow check 201 Created and the JWT token?
	}
}
