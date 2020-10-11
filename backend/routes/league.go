package routes

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/grommers00/civitas/backend/internal"
	"github.com/grommers00/civitas/backend/models"

	"github.com/gorilla/mux"
)

// ConnectLeaguesSubrouter adds the different restFUL apis to the main router
func ConnectLeaguesSubrouter(r *mux.Router) {
	// "/leagues/"
	f := r.PathPrefix("/leagues").Subrouter()

	// "/leagues/"
	f.HandleFunc("/", AddLeagues).Methods("POST")

	// "/leagues/"
	f.HandleFunc("/", GetAllLeagues).Methods("GET")

	// "/leagues/{ID}"
	f.HandleFunc("/{ID}", GetLeaguesByID).Methods("GET")

	// "/leagues/{ID}"
	f.HandleFunc("/{ID}", DeleteLeaguesByID).Methods("DELETE")

	// "/leagues/{ID}"
	f.HandleFunc("/{ID}", UpdateLeaguesByID).Methods("PATCH")

}

// AddLeagues will add a new article the list of articles
func AddLeagues(w http.ResponseWriter, r *http.Request) {
	internal.NotImplementedHandler("leagues Added!\n", w)
}

// GetAllLeagues gets all the leagues articles
func GetAllLeagues(w http.ResponseWriter, r *http.Request) {
	// Reading from json file
	league := []models.League{}
	err := internal.UnwrapJSONData("mockdata/mockleague.json", &league)

	// Creating leagues array with all the objects in the file
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	// Sends the json object of a singular device
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(league)
}

// GetLeaguesByID gets a new article by ID
func GetLeaguesByID(w http.ResponseWriter, r *http.Request) {
	// Reading from json file
	league := []models.League{}
	err := internal.UnwrapJSONData("mockdata/mockleague.json", &league)

	// Creating leagues array with all the objects in the file
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	//Gets the parameter and looks for the object in the array
	vars := mux.Vars(r)
	id := vars["ID"]
	aleagues := models.League{}

	// TODO: integrate DB for proper ID lookups
	for i := range league {
		if strconv.Itoa(league[i].ID) == id {
			aleagues = league[i]
		}
	}

	// Sends a json object of one league
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(aleagues)
}

// DeleteLeaguesByID will delete an article by its ID
func DeleteLeaguesByID(w http.ResponseWriter, r *http.Request) {
	internal.NotImplementedHandler("Deleted League\n", w)
}

// UpdateLeaguesByID updates the leagues by id.
func UpdateLeaguesByID(w http.ResponseWriter, r *http.Request) {
	internal.NotImplementedHandler("Updated a league!\n", w)
}
