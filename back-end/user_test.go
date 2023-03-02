package main

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPostUser(t *testing.T) {
	fmt.Print("is it even here??? ")

	var jsonStr = []byte(`{"username":gabjc,"password":"supersecretpassword","name":"gabriel"}`)
	fmt.Print("jsonStr ")

	req, err := http.NewRequest("POST", "/user/post/{username}/{password}/{name}", bytes.NewBuffer(jsonStr))
	fmt.Print("req, err ")
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	fmt.Print("setheader ")
	rr := httptest.NewRecorder()
	fmt.Print(("httptest "))
	handler := http.HandlerFunc(PostUser)
	fmt.Print("handler ")
	handler.ServeHTTP(rr, req)
	fmt.Print("servehttp ")
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	expected := `{"username":gabjc,"password":"supersecretpassword","name":"gabriel"}`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}
