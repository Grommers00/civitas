package main

import (
	"os"

	"github.com/gorilla/mux"
	"github.com/grommers00/civitas/backend/routes"

	"log"
	"net/http"

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

//HandleRequests listens for any requests that come in.
func HandleRequests(applicationConfig ApplicationConfiguration) {
	r := mux.NewRouter()
	r = routes.ConnectNewsSubrouter(r)
	log.Fatal(http.ListenAndServe(applicationConfig.Port, r))
}

func main() {
	applicationConfig := GoDotEnvVariable()
	HandleRequests(applicationConfig)
}
