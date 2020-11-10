package routes

import (
	"encoding/json"
	"log"
	"net/http"
	"sort"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/grommers00/civitas/backend/internal"
	"github.com/grommers00/civitas/backend/models"
)

// ConnectStandingSubrouter creates the /standing subroutes
func ConnectStandingSubrouter(r *mux.Router) {
	// "/standing/"
	f := r.PathPrefix("/standing").Subrouter()

	// "/standing/"
	f.HandleFunc("/", GetAllStandings).Methods("GET")
	f.HandleFunc("/", CreateStanding).Methods("POST")
	f.HandleFunc("/{ID}", GetStandingByID).Methods("GET")
	f.HandleFunc("/{ID}", UpdateStandingByID).Methods("PUT")
	f.HandleFunc("/{ID}", DeleteStandingByID).Methods("DELETE")

}

// GetAllStandings retreives all standings without filters
func GetAllStandings(w http.ResponseWriter, r *http.Request) {
	standings := []models.Standing{}

	if err := internal.UnwrapJSONData("mockdata/mockstanding.json", &standings); err != nil {
		log.Fatalf("Error loading data from file")
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(standings)
}

// GetStandingByID retreives a standing based on supplied ID
func GetStandingByID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["ID"])

	if err != nil {
		log.Fatalf("Error loading retrieving ID from URL")
	}

	standings := []models.Standing{}

	if err := internal.UnwrapJSONData("mockdata/mockstanding.json", &standings); err != nil {
		log.Fatalf("Error loading data from file")
	}

	standingsFiltered := []models.Standing{}

	// TODO: Integrate as a DB Query instead
	for _, standing := range standings {
		if standing.SeasonID == id {
			standingsFiltered = append(standingsFiltered, standing)
		}
	}

	// TODO: Integrate into dedicated data service
	// Sort based on rank logic ((A-B wins)+(A-B losses))/2)

	sort.Slice(standingsFiltered, func(i, j int) bool {

		wins := standingsFiltered[i].Wins - standingsFiltered[j].Wins
		loses := standingsFiltered[i].Loses - standingsFiltered[j].Loses

		return ((wins + loses) / 2) < 2
	})

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(standingsFiltered)
}

// DeleteStandingByID will delete an standing by its ID
func DeleteStandingByID(w http.ResponseWriter, r *http.Request) {
	internal.NotImplementedHandler("DeleteStandingByID", w)
}

// CreateStanding will create a new standing object
func CreateStanding(w http.ResponseWriter, r *http.Request) {
	internal.NotImplementedHandler("CreateStanding", w)
}

// UpdateStandingByID will update an standing by its ID
func UpdateStandingByID(w http.ResponseWriter, r *http.Request) {
	internal.NotImplementedHandler("UpdateStandingByID", w)
}
