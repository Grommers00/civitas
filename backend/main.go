package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

// ApplicationConfiguration contains all the ENV variables that will be used within the backend
type ApplicationConfiguration struct {
	Port string
}

// GoDotEnvVariable gets a list of all the variables into the handle requests
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

//HomePage lets you get to the home page and notifies you with a println terminal.
func HomePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

// FlutterInitRequest TODO: This is a test API for Flutter Integration, Temp.sends "Hello, World". Meta.
func FlutterInitRequest(w http.ResponseWriter, r *http.Request) {

}

//HandleRequests listens for any requests that come in.
func HandleRequests(applicationConfig ApplicationConfiguration) {

	http.HandleFunc("/", HomePage)
	log.Fatal(http.ListenAndServe(applicationConfig.Port, nil))
}

func main() {
	applicationConfig := goDotEnvVariable()
	HandleRequests(applicationConfig)
}
