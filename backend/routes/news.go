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

// ConnectNewsSubrouter adds the different restFUL apis to the main router
func ConnectNewsSubrouter(r *mux.Router) *mux.Router {
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

	return f
}

// AddNews will add a new article the list of articles
func AddNews(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("News Added!\n"))
}

// GetAllNews gets all the news articles
func GetAllNews(w http.ResponseWriter, r *http.Request) {

	jsonFile, err := os.Open("mockdata/mocknews.json")
	byteValue, _ := ioutil.ReadAll(jsonFile)
	defer jsonFile.Close()
	news := []models.News{}
	json.Unmarshal([]byte(byteValue), &news)
	if err != nil {
		fmt.Println(err)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(news)
}

// GetNewsByID gets a new article by ID
func GetNewsByID(w http.ResponseWriter, r *http.Request) {
	jsonFile, err := os.Open("mockdata/mocknews.json")
	byteValue, _ := ioutil.ReadAll(jsonFile)
	jsonFile.Close()
	news := []models.News{}
	json.Unmarshal([]byte(byteValue), &news)
	if err != nil {
		fmt.Println(err)
	}
	vars := mux.Vars(r)
	id := vars["ID"]
	aNews := models.News{}
	// TODO: integrate DB for proper ID lookups
	for i := range news {
		if strconv.Itoa(news[i].ID) == id {
			aNews = news[i]
		}
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(aNews)
}

// DeleteNewsByID will delete an article by its ID
func DeleteNewsByID(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("There!\n"))
}

// UpdateNewsByID updates the news by id.
func UpdateNewsByID(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Where!\n"))
}
