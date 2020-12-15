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
	ErrorInvalidNewsData         = "invalid News data"
	ErrorInvalidEmail            = "invalid name"
	ErrorCouldNotMarshalItem     = "could not marshal item"
	ErrorCouldNotDeleteItem      = "could not delete item"
	ErrorCouldNotDynamoPutItem   = "could not dynamo put item error"
	ErrorNewsAlreadyExists       = "news.News already exists"
	ErrorNewsDoesNotExists       = "news.News does not exist"
	TableName                    = "civitas-news"
)

type News struct {
	ID     string `json:"id"`
	Date   string `json:"date"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Body   string `json:"description"`
	Image  string `json:"image"`
}

func FetchNews(id string, dynaClient dynamodbiface.DynamoDBAPI) (*News, error) {
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

	item := new(News)
	err = dynamodbattribute.UnmarshalMap(result.Item, item)
	if err != nil {
		return nil, errors.New(ErrorFailedToUnmarshalRecord)
	}
	return item, nil
}

func FetchNewsAll(dynaClient dynamodbiface.DynamoDBAPI) (*[]News, error) {
	input := &dynamodb.ScanInput{
		TableName: aws.String(TableName),
	}
	result, err := dynaClient.Scan(input)
	if err != nil {
		return nil, errors.New(ErrorFailedToFetchRecord)
	}
	item := new([]News)
	err = dynamodbattribute.UnmarshalListOfMaps(result.Items, item)
	return item, nil
}

func CreateNews(req events.APIGatewayProxyRequest, dynaClient dynamodbiface.DynamoDBAPI) (
	*News,
	error,
) {
	var nw News
	if err := json.Unmarshal([]byte(req.Body), &nw); err != nil {
		return nil, errors.New(ErrorInvalidNewsData)
	}

	// Check if user exists
	currentNews, _ := FetchNews(nw.ID, dynaClient)
	if currentNews != nil && len(currentNews.ID) != 0 {
		return nil, errors.New(ErrorNewsAlreadyExists)
	}

	// Save user
	av, err := dynamodbattribute.MarshalMap(nw)
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
	return &nw, nil
}

func UpdateNews(req events.APIGatewayProxyRequest, dynaClient dynamodbiface.DynamoDBAPI) (
	*News,
	error,
) {
	var nw News
	if err := json.Unmarshal([]byte(req.Body), &nw); err != nil {
		return nil, errors.New(err.Error())
	}

	// Check if user exists
	currentNews, _ := FetchNews(nw.ID, dynaClient)
	if currentNews != nil && len(currentNews.ID) == 0 {
		return nil, errors.New(ErrorNewsDoesNotExists)
	}

	// Save user
	av, err := dynamodbattribute.MarshalMap(nw)
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
	return &nw, nil
}

func DeleteNews(req events.APIGatewayProxyRequest, dynaClient dynamodbiface.DynamoDBAPI) error {
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
