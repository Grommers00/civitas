package db

import (
	"encoding/json"
	"errors"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
	"github.com/google/uuid"
)

const (
	ErrorFailedToUnmarshalRecord = "failed to unmarshal record"
	ErrorFailedToFetchRecord     = "failed to fetch record"
	ErrorInvalidDivisionData     = "invalid Division data"
	ErrorInvalidEmail            = "invalid name"
	ErrorCouldNotMarshalItem     = "could not marshal item"
	ErrorCouldNotDeleteItem      = "could not delete item"
	ErrorCouldNotDynamoPutItem   = "could not dynamo put item error"
	ErrorDivisionAlreadyExists   = "match.Division already exists"
	ErrorDivisionDoesNotExists   = "match.Division does not exist"
	TableName                    = "civitas-division"
)

type Division struct {
	ID        string `json:"id, omitempty"`
	StartDate string `json:"startdate"`
	EndDate   string `json:"enddate"`
	Game      string `json:"game"`
	Map       string `json:"map"`
	Desc      string `json:"description"`
	Season    int    `json:"season"`
}

func FetchDivision(id string, dynaClient dynamodbiface.DynamoDBAPI) (*Division, error) {
	input := &dynamodb.GetItemInput{
		Key: map[string]*dynamodb.AttributeValue{
			"id": {
				S: aws.String(id),
			},
		},
		TableName: aws.String(TableName),
	}

	result, err := dynaClient.GetItem(input)
	if err != nil {
		return nil, err

	}

	item := new(Division)
	err = dynamodbattribute.UnmarshalMap(result.Item, item)
	if err != nil {
		return nil, err
	}
	return item, nil
}

func FetchDivisiones(dynaClient dynamodbiface.DynamoDBAPI) (*[]Division, error) {
	input := &dynamodb.ScanInput{
		TableName: aws.String(TableName),
	}
	result, err := dynaClient.Scan(input)
	if err != nil {
		return nil, errors.New(ErrorFailedToFetchRecord)
	}
	item := new([]Division)
	err = dynamodbattribute.UnmarshalListOfMaps(result.Items, item)
	return item, nil
}

func CreateDivision(req events.APIGatewayProxyRequest, dynaClient dynamodbiface.DynamoDBAPI) (
	*Division,
	error,
) {
	var mt Division
	if err := json.Unmarshal([]byte(req.Body), &mt); err != nil {
		return nil, errors.New(ErrorInvalidDivisionData)
	}

	mt.ID = uuid.New().String()

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

func UpdateDivision(req events.APIGatewayProxyRequest, dynaClient dynamodbiface.DynamoDBAPI) (
	*Division,
	error,
) {
	var mt Division
	if err := json.Unmarshal([]byte(req.Body), &mt); err != nil {
		return nil, errors.New(err.Error())
	}

	// Check if user exists
	currentDivision, _ := FetchDivision(mt.ID, dynaClient)
	if currentDivision != nil && len(currentDivision.ID) == 0 {
		return nil, errors.New(ErrorDivisionDoesNotExists)
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

func DeleteDivision(req events.APIGatewayProxyRequest, dynaClient dynamodbiface.DynamoDBAPI) error {
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
