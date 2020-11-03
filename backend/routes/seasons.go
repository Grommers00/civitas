package routes

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/grommers00/civitas/backend/internal"
	"github.com/grommers00/civitas/backend/models"
)

// ConnectSeasonsRouter creates the /seasons subroutes
func ConnectSeasonsRouter(r *mux.Router) {
	// "/profile/"
	f := r.PathPrefix("/seasons").Subrouter()

	// "/profile/"
	f.HandleFunc("/", GetAllSeasons).Methods("GET")
}

// GetAllSeasons retrieves all seasons without a filter
func GetAllSeasons(w http.ResponseWriter, r *http.Request) {
	standings := []models.Season{}

	if err := internal.UnwrapJSONData("mockdata/mockseason.json", &standings); err != nil {
		log.Fatalf("Error loading data from file")
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(standings)
}
