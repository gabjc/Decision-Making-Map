package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gorilla/mux"
)

// TODO: abstract testing funcs by adding more helper funcs
func TestGetUser(t *testing.T) {
	InitDB()

	// create test router
	router := mux.NewRouter()
	router.HandleFunc("/user/get/{username}", GetUser).Methods("GET")

	// define test request
	req, err := http.NewRequest("GET", "/user/get/admin", nil)
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
	expected := `{"ID":0,"CreatedAt":"2023-02-28T19:02:34.8143566-05:00","UpdatedAt":"2023-02-28T19:02:34.8143566-05:00","DeletedAt":null,"username":"admin","password":"password","name":"Admin Adminton"}`
	if strings.TrimRight(rr.Body.String(), "\n") != expected {
		t.Errorf("Handler returned unexpected body:\ngot \n%v want \n%v", rr.Body.String(), expected)
	}
}

func TestPostUser(t *testing.T) {
	InitDB()

	// create test router
	router := mux.NewRouter()
	router.HandleFunc("/user/post/{username}/{password}/{name}", PostUser).Methods("POST")

	// define test request
	// TODO: why do we need the request body? Route vs query parameters?
	reqBody := strings.NewReader(`{"Username": "test2", "Password": "test2", "Name": "test2"}`)
	req, err := http.NewRequest("POST", "/user/post/test2/test2/test2", reqBody)
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
	var user User
	username := "test2"
	err = db.First(&user, "username = ?", username).Error
	if err != nil {
		t.Errorf("Error finding user in database")
	}
}
