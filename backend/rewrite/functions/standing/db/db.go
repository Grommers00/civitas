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
	ErrorInvalidStandingData     = "invalid standing data"
	ErrorInvalidEmail            = "invalid name"
	ErrorCouldNotMarshalItem     = "could not marshal item"
	ErrorCouldNotDeleteItem      = "could not delete item"
	ErrorCouldNotDynamoPutItem   = "could not dynamo put item error"
	ErrorStandingAlreadyExists   = "standing.Standing already exists"
	ErrorStandingDoesNotExists   = "standing.Standing does not exist"
	TableName                    = "civitas-standing"
)

type Standing struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Desc string `json:"desc"`
}

func FetchStanding(id string, dynaClient dynamodbiface.DynamoDBAPI) (*Standing, error) {
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

	item := new(Standing)
	err = dynamodbattribute.UnmarshalMap(result.Item, item)
	if err != nil {
		return nil, errors.New(ErrorFailedToUnmarshalRecord)
	}
	return item, nil
}

func FetchStandings(dynaClient dynamodbiface.DynamoDBAPI) (*[]Standing, error) {
	input := &dynamodb.ScanInput{
		TableName: aws.String(TableName),
	}
	result, err := dynaClient.Scan(input)
	if err != nil {
		return nil, errors.New(ErrorFailedToFetchRecord)
	}
	item := new([]Standing)
	err = dynamodbattribute.UnmarshalListOfMaps(result.Items, item)
	return item, nil
}

func CreateStanding(req events.APIGatewayProxyRequest, dynaClient dynamodbiface.DynamoDBAPI) (
	*Standing,
	error,
) {
	var pf Standing
	if err := json.Unmarshal([]byte(req.Body), &pf); err != nil {
		return nil, errors.New(ErrorInvalidStandingData)
	}

	// Check if user exists
	currentStanding, _ := FetchStanding(pf.ID, dynaClient)
	if currentStanding != nil && len(currentStanding.ID) != 0 {
		return nil, errors.New(ErrorStandingAlreadyExists)
	}

	// Save user
	av, err := dynamodbattribute.MarshalMap(pf)
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
	return &pf, nil
}

func UpdateStanding(req events.APIGatewayProxyRequest, dynaClient dynamodbiface.DynamoDBAPI) (
	*Standing,
	error,
) {
	var pf Standing
	if err := json.Unmarshal([]byte(req.Body), &pf); err != nil {
		return nil, errors.New(err.Error())
	}

	// Check if user exists
	currentStanding, _ := FetchStanding(pf.ID, dynaClient)
	if currentStanding != nil && len(currentStanding.ID) == 0 {
		return nil, errors.New(ErrorStandingDoesNotExists)
	}

	// Save user
	av, err := dynamodbattribute.MarshalMap(pf)
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
	return &pf, nil
}

func DeleteStanding(req events.APIGatewayProxyRequest, dynaClient dynamodbiface.DynamoDBAPI) error {
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
