package main

import (
	"context"
	"os"

	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/awslabs/aws-lambda-go-api-proxy/gorillamux"
	"github.com/gorilla/mux"
	"github.com/grommers00/civitas/backend/models"
	"github.com/grommers00/civitas/backend/routes"

	"github.com/joho/godotenv"
)

var muxLambda *gorillamux.GorillaMuxAdapter

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

//ConstructRoutes generates the applications primary router, and subroutes
func ConstructRoutes() *mux.Router {
	r := mux.NewRouter()

	routes.ConnectNewsSubrouter(r)
	routes.ConnectProfileSubrouter(r)
	routes.ConnectLeaguesSubrouter(r)
	routes.ConnectStandingSubrouter(r)
	routes.ConnectSeasonsRouter(r)

	return r
}

// InitializeApplication creates the application instance
func InitializeApplication() models.Application {
	return models.Application{
		// Config: GoDotEnvVariable(),
		Router: ConstructRoutes(),
	}
}

func main() {
	application := InitializeApplication()
	muxLambda = gorillamux.New(application.Router)
	// log.Fatal(http.ListenAndServe(application.Config.Port, application.Router))

	lambda.Start(handler)
}

func handler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// If no name is provided in the HTTP request body, throw an error
	return muxLambda.ProxyWithContext(ctx, req)
}
