package main

import (
	"net/http"
)

func GetStatus(w http.ResponseWriter, r *http.Request) { // FINISH IMPLEMENTING, UN-COMMENT STATUS FETCH IN app.component.ts
	/* w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	resp := make(map[string]string)
	resp["message"] = "status: UP"
	jsonResp, err := json.Marshal(resp)
	if err != nil {
		fmt.Println(err.Error())
		panic("Cannot convert status OK to JSON")
	}
	w.Write(jsonResp)
	return */
}
