package db

import (
	"encoding/json"
	"errors"

	"github.com/google/uuid"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
)

const (
	ErrorFailedToUnmarshalRecord = "failed to unmarshal record"
	ErrorFailedToFetchRecord     = "failed to fetch record"
	ErrorInvalidWinnerListData   = "invalid winnerlist data"
	ErrorInvalidEmail            = "invalid name"
	ErrorCouldNotMarshalItem     = "could not marshal item"
	ErrorCouldNotDeleteItem      = "could not delete item"
	ErrorCouldNotDynamoPutItem   = "could not dynamo put item error"
	ErrorWinnerListAlreadyExists = "winnerlist.WinnerList already exists"
	ErrorWinnerListDoesNotExists = "winnerlist.WinnerList does not exist"
	TableName                    = "civitas-winnerlist"
)

type WinnerList struct {
	ID   string `json:"id, omitempty"`
	Name string `json:"name"`
	Desc string `json:"desc"`
}

func FetchWinnerList(id string, dynaClient dynamodbiface.DynamoDBAPI) (*WinnerList, error) {
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

	item := new(WinnerList)
	err = dynamodbattribute.UnmarshalMap(result.Item, item)
	if err != nil {
		return nil, err
	}
	return item, nil
}

func FetchWinnerLists(dynaClient dynamodbiface.DynamoDBAPI) (*[]WinnerList, error) {
	input := &dynamodb.ScanInput{
		TableName: aws.String(TableName),
	}
	result, err := dynaClient.Scan(input)
	if err != nil {
		return nil, errors.New(ErrorFailedToFetchRecord)
	}
	item := new([]WinnerList)
	err = dynamodbattribute.UnmarshalListOfMaps(result.Items, item)
	return item, nil
}

func CreateWinnerList(req events.APIGatewayProxyRequest, dynaClient dynamodbiface.DynamoDBAPI) (
	*WinnerList,
	error,
) {
	var wl WinnerList
	if err := json.Unmarshal([]byte(req.Body), &wl); err != nil {
		return nil, errors.New(ErrorInvalidWinnerListData)
	}

	wl.ID = uuid.New().String()

	// Save user
	av, err := dynamodbattribute.MarshalMap(wl)
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
	return &wl, nil
}

func UpdateWinnerList(req events.APIGatewayProxyRequest, dynaClient dynamodbiface.DynamoDBAPI) (
	*WinnerList,
	error,
) {
	var wl WinnerList
	if err := json.Unmarshal([]byte(req.Body), &wl); err != nil {
		return nil, errors.New(err.Error())
	}

	// Check if user exists
	currentWinnerList, _ := FetchWinnerList(wl.ID, dynaClient)
	if currentWinnerList != nil && len(currentWinnerList.ID) == 0 {
		return nil, errors.New(ErrorWinnerListDoesNotExists)
	}

	// Save user
	av, err := dynamodbattribute.MarshalMap(wl)
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
	return &wl, nil
}

func DeleteWinnerList(req events.APIGatewayProxyRequest, dynaClient dynamodbiface.DynamoDBAPI) error {
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
