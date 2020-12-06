package db

import (
	"encoding/json"
	"errors"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
)

const (
	ErrorFailedToUnmarshalRecord = "failed to unmarshal record"
	ErrorFailedToFetchRecord     = "failed to fetch record"
	ErrorInvalidLeagueData       = "invalid League data"
	ErrorInvalidEmail            = "invalid name"
	ErrorCouldNotMarshalItem     = "could not marshal item"
	ErrorCouldNotDeleteItem      = "could not delete item"
	ErrorCouldNotDynamoPutItem   = "could not dynamo put item error"
	ErrorLeagueAlreadyExists     = "league.League already exists"
	ErrorLeagueDoesNotExists     = "league.League does not exist"
	TableName                    = "civitas-league"
)

type League struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Season int    `json:"season"`
	Game   string `json:"game"`
}

func FetchLeague(id string, dynaClient dynamodbiface.DynamoDBAPI) (*League, error) {
	input := &dynamodb.GetItemInput{
		Key: map[string]*dynamodb.AttributeValue{
			"id": {
				N: aws.String(id),
			},
		},
		TableName: aws.String(TableName),
	}

	result, err := dynaClient.GetItem(input)
	if err != nil {
		return nil, errors.New(ErrorFailedToFetchRecord)

	}

	item := new(League)
	err = dynamodbattribute.UnmarshalMap(result.Item, item)
	if err != nil {
		return nil, errors.New(ErrorFailedToUnmarshalRecord)
	}
	return item, nil
}

func FetchLeagues(dynaClient dynamodbiface.DynamoDBAPI) (*[]League, error) {
	input := &dynamodb.ScanInput{
		TableName: aws.String(TableName),
	}
	result, err := dynaClient.Scan(input)
	if err != nil {
		return nil, errors.New(ErrorFailedToFetchRecord)
	}
	item := new([]League)
	err = dynamodbattribute.UnmarshalListOfMaps(result.Items, item)
	return item, nil
}

func CreateLeague(req events.APIGatewayProxyRequest, dynaClient dynamodbiface.DynamoDBAPI) (
	*League,
	error,
) {
	var lg League
	if err := json.Unmarshal([]byte(req.Body), &lg); err != nil {
		return nil, errors.New(ErrorInvalidLeagueData)
	}

	// Check if user exists
	currentLeague, _ := FetchLeague(lg.ID, dynaClient)
	if currentLeague != nil && len(currentLeague.ID) != 0 {
		return nil, errors.New(ErrorLeagueAlreadyExists)
	}

	// Save user
	av, err := dynamodbattribute.MarshalMap(lg)
	if err != nil {
		return nil, errors.New(ErrorCouldNotMarshalItem)
	}

	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String(TableName),
	}

	_, err = dynaClient.PutItem(input)
	if err != nil {
		return nil, errors.New(err.Error())
	}
	return &lg, nil
}

func UpdateLeague(req events.APIGatewayProxyRequest, dynaClient dynamodbiface.DynamoDBAPI) (
	*League,
	error,
) {
	var lg League
	if err := json.Unmarshal([]byte(req.Body), &lg); err != nil {
		return nil, errors.New(err.Error())
	}

	// Check if user exists
	currentLeague, _ := FetchLeague(lg.ID, dynaClient)
	if currentLeague != nil && len(currentLeague.ID) == 0 {
		return nil, errors.New(ErrorLeagueDoesNotExists)
	}

	// Save user
	av, err := dynamodbattribute.MarshalMap(lg)
	if err != nil {
		return nil, errors.New(err.Error())
	}

	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String(TableName),
	}

	_, err = dynaClient.PutItem(input)
	if err != nil {
		return nil, errors.New(err.Error())
	}
	return &lg, nil
}

func DeleteLeague(req events.APIGatewayProxyRequest, dynaClient dynamodbiface.DynamoDBAPI) error {
	id := req.QueryStringParameters["id"]
	input := &dynamodb.DeleteItemInput{
		Key: map[string]*dynamodb.AttributeValue{
			"id": {
				S: aws.String(id),
			},
		},
		TableName: aws.String(TableName),
	}
	_, err := dynaClient.DeleteItem(input)
	if err != nil {
		return errors.New(ErrorCouldNotDeleteItem)
	}

	return nil
}
