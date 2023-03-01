package main

type votedLocation struct {
	currentVote int    `json:"vote"`
	voteName    string `json:"vote_name"`
}

type Itinerary struct {
	ItineraryID string        `json:"id"`
	Name        string        `json:"name"`
	Address     string        `json:"address"`
	City        string        `json:"city"`
	State       string        `json:"state"`
	Zip         string        `json:"zip"`
	Country     string        `json:"country"`
	Phone       string        `json:"phone"`
	Radius      int           `json:"radius"`
	voteList    votedLocation `json:"voted_Locations"`
}

/*

func GetItinerary(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	username := mux.Vars(r)["username"]
	var user User
	db.First(&user, "username = ?", username)

	json.NewEncoder(w).Encode(user)
}

func PostItinerary(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	user := User{Username: params["username"], Password: params["password"], Name: params["name"]}
	json.NewDecoder(r.Body).Decode(&user)
	db.Create(&user)

	json.NewEncoder(w).Encode(user)
}

*/
