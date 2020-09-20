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

func ConnectNewsSubrouter(r *mux.Router) *mux.Router {
	// "/news/"
	f := r.PathPrefix("/news").Subrouter()

	// "/news/"
	f.HandleFunc("/", AddNews).Methods("POST")

	// "/news/"
	f.HandleFunc("/", GetAllNews).Methods("GET")

	// "/news/{ID}"
	f.HandleFunc("/{ID}", GetNewsById).Methods("GET")

	// "/news/{ID}"
	f.HandleFunc("/{ID}", DeleteNewsById).Methods("DELETE")

	// "/news/{ID}"
	f.HandleFunc("/{ID}", UpdateNewsById).Methods("PATCH")

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

// Gets a new article by ID
func GetNewsById(w http.ResponseWriter, r *http.Request) {
	jsonFile, err := os.Open("mockdata/mocknews.json")
	byteValue, _ := ioutil.ReadAll(jsonFile)
	defer jsonFile.Close()
	news := []models.News{}
	json.Unmarshal([]byte(byteValue), &news)
	if err != nil {
		fmt.Println(err)
	}
	vars := mux.Vars(r)

	id := vars["ID"]
	aNews := models.News{}
	for i := range news {
		if strconv.Itoa(news[i].ID) == id {
			fmt.Println("Do yiou get in this loop?")
			aNews = news[i]
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(aNews)
}

// DeleteNewsById will delete an article by its ID
func DeleteNewsById(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("There!\n"))
}

// UpdateNewsById updates the news by id.
func UpdateNewsById(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Where!\n"))
}
