package routes

import (
	"encoding/json"
	"net/http"

	// "github.com/grommers00/civitas/backend/models"

	"github.com/gorilla/mux"
	"github.com/grommers00/civitas/backend/internal"
	"github.com/grommers00/civitas/backend/models"
)

func ConnectProfileSubrouter(r *mux.Router) {
	// "/profile/"
	f := r.PathPrefix("/profile").Subrouter()

	// "/profile/"
	f.HandleFunc("/", AddNews).Methods("POST")

	// "/profile/"
	f.HandleFunc("/", GetAllProfiles).Methods("GET")

	// "/profile/{ID}"
	f.HandleFunc("/{ID}", GetProfileByID).Methods("GET")

	// "/profile/{ID}"
	f.HandleFunc("/{ID}", DeleteProfileByID).Methods("DELETE")

	// "/profile/{ID}"
	f.HandleFunc("/{ID}", UpdateProfileByID).Methods("PATCH")
}

// GetAllProfiles retrives all public profiles
func GetAllProfiles(w http.ResponseWriter, r *http.Request) {
	profile := []models.Profile{}
	err := internal.UnwrapJSONData("mockdata/mockprofiles.json", &profile)

	if err != nil {
		panic(err)
	}

	// Sends the json object of a singular device
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(profile)
}

// GetProfileByID retrieves specified profiled based on /profile/{ID} route
func GetProfileByID(w http.ResponseWriter, r *http.Request) {
	internal.NotImplementedHandler("GetProfileByID", w)
}

// SaveProfile creates and commits new profle object via form submission
func SaveProfile(w http.ResponseWriter, r *http.Request) {
	internal.NotImplementedHandler("SaveProfile", w)

}

// DeleteProfileByID will delete a profile specified via /profile/{ID} route
func DeleteProfileByID(w http.ResponseWriter, r *http.Request) {
	internal.NotImplementedHandler("DeleteProfleByID", w)
}

// UpdateProfileByID will update a profile specified via /profile/{ID} route
func UpdateProfileByID(w http.ResponseWriter, r *http.Request) {
	internal.NotImplementedHandler("UpdateProfileByID", w)
}
