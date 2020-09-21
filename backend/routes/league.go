package routes

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"

	"github.com/grommers00/civitas/backend/models"

	"github.com/gorilla/mux"
)

// ConnectLeaguesSubrouter adds the different restFUL apis to the main router
func ConnectLeaguesSubrouter(r *mux.Router) *mux.Router {
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

	return f
}

// AddLeagues will add a new article the list of articles
func AddLeagues(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("leagues Added!\n"))
}

// GetAllLeagues gets all the leagues articles
func GetAllLeagues(w http.ResponseWriter, r *http.Request) {
	// Reading from json file
	jsonFile, err := os.Open("mockdata/mockleagues.json")
	byteValue, _ := ioutil.ReadAll(jsonFile)
	jsonFile.Close()

	// Creating leagues array with all the objects in the file
	leagues := []models.League{}
	json.Unmarshal([]byte(byteValue), &leagues)
	if err != nil {
		fmt.Println(err)
	}

	// Sends the json object of a singular device
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(leagues)
}

// GetLeaguesByID gets a new article by ID
func GetLeaguesByID(w http.ResponseWriter, r *http.Request) {
	// Reading from json file
	jsonFile, err := os.Open("mockdata/mockleagues.json")
	byteValue, _ := ioutil.ReadAll(jsonFile)
	jsonFile.Close()

	//Creating leagues array with all the objects in the file
	leagues := []models.League{}
	json.Unmarshal([]byte(byteValue), &leagues)
	if err != nil {
		fmt.Println(err)
	}

	//Gets the parameter and looks for the object in the array
	vars := mux.Vars(r)
	id := vars["ID"]
	aleagues := models.League{}

	// TODO: integrate DB for proper ID lookups
	for i := range leagues {
		if strconv.Itoa(leagues[i].ID) == id {
			aleagues = leagues[i]
		}
	}

	// Sends a json object of one league
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(aleagues)
}

// DeleteLeaguesByID will delete an article by its ID
func DeleteLeaguesByID(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Deleted a league!\n"))
}

// UpdateLeaguesByID updates the leagues by id.
func UpdateLeaguesByID(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Updated a league!\n"))
}
