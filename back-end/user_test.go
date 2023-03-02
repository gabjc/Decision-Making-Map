package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

// test getting the "admin" user specfically...will fail if admin is missing from db
func TestGetUser(t *testing.T) {
	req, err := http.NewRequest("GET", "/user/get/admin", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(GetUser)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// verify that the response body is what we expect
	expected := `[{"ID":0,"CreatedAt":"2023-02-28T19:02:34.8143566-05:00","UpdatedAt":"2023-02-28T19:02:34.8143566-05:00","DeletedAt":null,"username":"admin","password":"password","name":"Admin Adminton"}]`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}
