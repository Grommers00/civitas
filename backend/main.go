package main

import (
	"os"

	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/grommers00/civitas/backend/models"
	"github.com/grommers00/civitas/backend/routes"

	"github.com/joho/godotenv"
)

// GoDotEnvVariable gets a list of all the variables into the handle requests
func GoDotEnvVariable() models.ApplicationConfiguration {
	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return models.ApplicationConfiguration{
		Port: os.Getenv("PORT"),
	}
}

//HandleRequests listens for any requests that come in.
func ConstructRoutes() *mux.Router {
	r := mux.NewRouter()
	routes.ConnectNewsSubrouter(r)
	routes.ConnectProfileSubrouter(r)

	return r
}

func IntializeApplication() models.Application {
	return models.Application{
		Config: GoDotEnvVariable(),
		Router: ConstructRoutes(),
	}
}

func main() {
	application := IntializeApplication()
	log.Fatal(http.ListenAndServe(application.Config.Port, application.Router))
}
