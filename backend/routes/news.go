package routes

import (
	"encoding/json"
	"net/http"

	"github.com/grommers00/civitas/backend/internal"
	"github.com/grommers00/civitas/backend/models"

	"github.com/gorilla/mux"
)

// ConnectNewsSubrouter adds the different restFUL apis to the main router
func ConnectNewsSubrouter(r *mux.Router) {
	// "/news/"
	f := r.PathPrefix("/news").Subrouter()

	// "/news/"
	f.HandleFunc("/", AddNews).Methods("POST")

	// "/news/"
	f.HandleFunc("/", GetAllNews).Methods("GET")

	// "/news/{ID}"
	f.HandleFunc("/{ID}", GetNewsByID).Methods("GET")

	// "/news/{ID}"
	f.HandleFunc("/{ID}", DeleteNewsByID).Methods("DELETE")

	// "/news/{ID}"
	f.HandleFunc("/{ID}", UpdateNewsByID).Methods("PATCH")
}

// AddNews will add a new article the list of articles
func AddNews(w http.ResponseWriter, r *http.Request) {
	internal.NotImplementedHandler("AddNews", w)
}

// GetAllNews gets all the news articles
func GetAllNews(w http.ResponseWriter, r *http.Request) {
	news := []models.News{}
	err := internal.UnwrapJSONData("mockdata/mocknews.json", &news)

	if err != nil {
		panic(err)
	}

	// Sends the json object of a singular device
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(news)
}

// GetNewsByID gets a new article by ID
func GetNewsByID(w http.ResponseWriter, r *http.Request) {
	internal.NotImplementedHandler("GetNewsByID", w)
}

// DeleteNewsByID will delete an article by its ID
func DeleteNewsByID(w http.ResponseWriter, r *http.Request) {
	internal.NotImplementedHandler("DeleteNewsByID", w)
}

// UpdateNewsByID updates the news by id.
func UpdateNewsByID(w http.ResponseWriter, r *http.Request) {
	internal.NotImplementedHandler("UpdateNewsByID", w)
}
