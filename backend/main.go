package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

type ApplicationConfiguration struct {
	Port string
}

func GoDotEnvVariable() ApplicationConfiguration {

	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return ApplicationConfiguration{
		Port: os.Getenv("PORT"),
	}
}

func HomePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

// TODO: This is a test API for Flutter Integration, Temp.
func FlutterInitRequest(w http.ResponseWriter, r *http.Request) {

}

func HandleRequests(applicationConfig ApplicationConfiguration) {

	http.HandleFunc("/", HomePage)
	log.Fatal(http.ListenAndServe(applicationConfig.Port, nil))
}

func main() {
	applicationConfig := GoDotEnvVariable()
	HandleRequests(applicationConfig)
}
