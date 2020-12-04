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
	ErrorInvalidMatchData        = "invalid Match data"
	ErrorInvalidEmail            = "invalid name"
	ErrorCouldNotMarshalItem     = "could not marshal item"
	ErrorCouldNotDeleteItem      = "could not delete item"
	ErrorCouldNotDynamoPutItem   = "could not dynamo put item error"
	ErrorMatchAlreadyExists      = "match.Match already exists"
	ErrorMatchDoesNotExists      = "match.Match does not exist"
	TableName                    = "civitas-matches"
)

type Match struct {
	ID        string `json:"id"`
	StartDate string `json:"startdate"`
	EndDate   string `json:"enddate"`
	Game      string `json:"game"`
	Map       string `json:"map"`
	Desc      string `json:"description"`
	Season    int    `json:"season"`
}

func FetchMatch(id string, dynaClient dynamodbiface.DynamoDBAPI) (*Match, error) {
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

	item := new(Match)
	err = dynamodbattribute.UnmarshalMap(result.Item, item)
	if err != nil {
		return nil, errors.New(ErrorFailedToUnmarshalRecord)
	}
	return item, nil
}

func FetchMatches(dynaClient dynamodbiface.DynamoDBAPI) (*[]Match, error) {
	input := &dynamodb.ScanInput{
		TableName: aws.String(TableName),
	}
	result, err := dynaClient.Scan(input)
	if err != nil {
		return nil, errors.New(ErrorFailedToFetchRecord)
	}
	item := new([]Match)
	err = dynamodbattribute.UnmarshalListOfMaps(result.Items, item)
	return item, nil
}

func CreateMatch(req events.APIGatewayProxyRequest, dynaClient dynamodbiface.DynamoDBAPI) (
	*Match,
	error,
) {
	var mt Match
	if err := json.Unmarshal([]byte(req.Body), &mt); err != nil {
		return nil, errors.New(ErrorInvalidMatchData)
	}

	// Check if user exists
	currentMatch, _ := FetchMatch(mt.ID, dynaClient)
	if currentMatch != nil && len(currentMatch.ID) != 0 {
		return nil, errors.New(ErrorMatchAlreadyExists)
	}

	// Save user
	av, err := dynamodbattribute.MarshalMap(mt)
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
	return &mt, nil
}

func UpdateMatch(req events.APIGatewayProxyRequest, dynaClient dynamodbiface.DynamoDBAPI) (
	*Match,
	error,
) {
	var mt Match
	if err := json.Unmarshal([]byte(req.Body), &mt); err != nil {
		return nil, errors.New(err.Error())
	}

	// Check if user exists
	currentMatch, _ := FetchMatch(mt.ID, dynaClient)
	if currentMatch != nil && len(currentMatch.ID) == 0 {
		return nil, errors.New(ErrorMatchDoesNotExists)
	}

	// Save user
	av, err := dynamodbattribute.MarshalMap(mt)
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
	return &mt, nil
}

func DeleteMatch(req events.APIGatewayProxyRequest, dynaClient dynamodbiface.DynamoDBAPI) error {
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
