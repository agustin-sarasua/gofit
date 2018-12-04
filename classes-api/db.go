package main

import (
	"fmt"

	"github.com/agustin-sarasua/gofit/model"
	"github.com/agustin-sarasua/gofit/util"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
)

const tableName = "Classes"

var classesUserSubGSI = "classesUserSubGSI"

// Declare a new DynamoDB instance. Note that this is safe for concurrent
// use.
var db dynamodbiface.DynamoDBAPI = dynamodb.New(session.New(), aws.NewConfig().WithRegion("us-east-1"))

// [Class-{uuid}] | [Start-Time] | ...
func putClass(c *model.Class) error {
	avEntity, err := dynamodbattribute.MarshalMap(c)
	if err != nil {
		return err
	}
	util.AddType(avEntity, model.DocTypeClass)
	partitionKey := fmt.Sprintf("%s-%s", model.DocTypeClass, c.ID)
	util.AddDyanmoDBKeys(avEntity, partitionKey, c.StartTime)

	input := &dynamodb.PutItemInput{
		TableName: aws.String(tableName),
		Item:      avEntity,
	}

	_, err = db.PutItem(input)
	return err
}
