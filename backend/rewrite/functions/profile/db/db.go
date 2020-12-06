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
	ErrorInvalidProfileData      = "invalid profile data"
	ErrorInvalidEmail            = "invalid name"
	ErrorCouldNotMarshalItem     = "could not marshal item"
	ErrorCouldNotDeleteItem      = "could not delete item"
	ErrorCouldNotDynamoPutItem   = "could not dynamo put item error"
	ErrorProfileAlreadyExists    = "profile.Profile already exists"
	ErrorProfileDoesNotExists    = "profile.Profile does not exist"
	TableName                    = "civitas-profile"
)

type Profile struct {
	ID   string `json:"id, omitempty"`
	Name string `json:"name"`
	Desc string `json:"desc"`
}

func FetchProfile(id string, dynaClient dynamodbiface.DynamoDBAPI) (*Profile, error) {
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

	item := new(Profile)
	err = dynamodbattribute.UnmarshalMap(result.Item, item)
	if err != nil {
		return nil, err
	}
	return item, nil
}

func FetchProfiles(dynaClient dynamodbiface.DynamoDBAPI) (*[]Profile, error) {
	input := &dynamodb.ScanInput{
		TableName: aws.String(TableName),
	}
	result, err := dynaClient.Scan(input)
	if err != nil {
		return nil, errors.New(ErrorFailedToFetchRecord)
	}
	item := new([]Profile)
	err = dynamodbattribute.UnmarshalListOfMaps(result.Items, item)
	return item, nil
}

func CreateProfile(req events.APIGatewayProxyRequest, dynaClient dynamodbiface.DynamoDBAPI) (
	*Profile,
	error,
) {
	var pf Profile
	if err := json.Unmarshal([]byte(req.Body), &pf); err != nil {
		return nil, errors.New(ErrorInvalidProfileData)
	}

	pf.ID = uuid.New().String()

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

func UpdateProfile(req events.APIGatewayProxyRequest, dynaClient dynamodbiface.DynamoDBAPI) (
	*Profile,
	error,
) {
	var pf Profile
	if err := json.Unmarshal([]byte(req.Body), &pf); err != nil {
		return nil, errors.New(err.Error())
	}

	// Check if user exists
	currentProfile, _ := FetchProfile(pf.ID, dynaClient)
	if currentProfile != nil && len(currentProfile.ID) == 0 {
		return nil, errors.New(ErrorProfileDoesNotExists)
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

func DeleteProfile(req events.APIGatewayProxyRequest, dynaClient dynamodbiface.DynamoDBAPI) error {
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
