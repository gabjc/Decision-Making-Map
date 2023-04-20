package main

import (
	"net/http"
)

func GetStatus(w http.ResponseWriter, r *http.Request) {
	SetContentJson(w, r)
	w.WriteHeader(http.StatusOK)
}
