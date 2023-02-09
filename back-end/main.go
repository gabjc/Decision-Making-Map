package main

// GENERAL:
// FIGURE OUT AUTO-INSTALLING FROM go.mod AND UPDATE README.TXT
// HAVE SERVE.SH NOT DELETE AND REBUILD EXE?

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func InitRouter() { // FIGURE OUT CONDITIONAL COMPILATION AND CORS SETTINGS FROM LECTURE EXAMPLE
	r := mux.NewRouter()

	r.HandleFunc("/status", GetStatus).Methods("GET")
	r.HandleFunc("/database/get", GetUser).Methods("GET")
	r.HandleFunc("/database/post/{ufid}/{name}", PostUser).Methods("POST")

	log.Fatal(http.ListenAndServe(":9000", r))
}

func main() {
	InitDatabase() // MOVE DB FUNC & FILE? OPEN & CLOSE DB ON A CALL-BY-CALL BASIS?
	InitRouter()
}
