package main

// GENERAL:
// MOVE GITHUB USER STORY WHEN EVERYTHING IS COMPLETED
// FIGURE OUT AUTO-INSTALLING FROM go.mod AND UPDATE README.TXT
// HAVE SERVE.SH NOT DELETE AND REBUILD EXE?
// LOOK INTO CYPRESS

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
	InitDatabase() // MOVE DB FUNC CALL & MERGE database.go?
	// OPEN & CLOSE DB ON A CALL-BY-CALL BASIS?
	// IMPLEMENT GORM?
	InitRouter()
}
