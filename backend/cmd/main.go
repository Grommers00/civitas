package main

import (
	"context"
	"log"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
	"github.com/awslabs/aws-lambda-go-api-proxy/gorillamux"
	"github.com/gorilla/mux"
)

var (
	dynaClient dynamodbiface.DynamoDBAPI
	muxLambda  *gorillamux.GorillaMuxAdapter
)

func init() {
	// stdout and stderr are sent to AWS CloudWatch Logs
	log.Printf("Mux: Cold Start")

	r := mux.NewRouter()

	// routes.ConnectNewsSubrouter(r)
	// routes.ConnectProfileSubrouter(r)
	// routes.ConnectLeaguesSubrouter(r)
	// routes.ConnectStandingSubrouter(r)
	// routes.ConnectSeasonsRouter(r)

	muxLambda = gorillamux.New(r)
}

func main() {
	region := os.Getenv("AWS_REGION")
	awsSession, err := session.NewSession(&aws.Config{
		Region: aws.String(region)},
	)
	if err != nil {
		return
	}

	dynaClient = dynamodb.New(awsSession)
	lambda.Start(Handler)
}

const tableName = "civitas-dev"

func Handler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// If no name is provided in the HTTP request body, throw an error
	return muxLambda.ProxyWithContext(ctx, req)
}

// func handler(req events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
// 	switch req.Path {
// 	case "profile":
// 		return ProfileHandlers(req)
// 	default:
// 		return handlers.UnhandledMethod()
// 	}
// }

// func ProfileHandlers(req events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
// 	switch req.HTTPMethod {
// 	case "GET":
// 		return handlers.GetProfile(req, tableName, dynaClient)
// 	case "POST":
// 		return handlers.CreateProfile(req, tableName, dynaClient)
// 	case "PUT":
// 		return handlers.UpdateProfile(req, tableName, dynaClient)
// 	case "DELETE":
// 		return handlers.DeleteProfile(req, tableName, dynaClient)
// 	default:
// 		return handlers.UnhandledMethod()
// 	}
// }
