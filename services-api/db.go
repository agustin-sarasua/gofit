package main

import (
	"fmt"

	"github.com/agustin-sarasua/gofit-companies-api/model"
	"github.com/agustin-sarasua/gofit-companies-api/util"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
)

const tableName = "Companies"

var companiesUserSubGSI = "companiesUserSubGSI"

// Declare a new DynamoDB instance. Note that this is safe for concurrent
// use.
var db dynamodbiface.DynamoDBAPI = dynamodb.New(session.New(), aws.NewConfig().WithRegion("us-east-1"))

// This method saves two records in DynamoDB
// [Service-{uuid}] | [Service-{uuid}] | ...
// [Company-{uuid}] | [Service-{uuid}] | ...
func putCompanyService(s *model.CompanyService) error {

	// [Service-{uuid}] | [Service-{uuid}] | ...
	avEntity, err := dynamodbattribute.MarshalMap(s)
	util.AddType(avEntity, model.DocTypeService)
	partitionKey := fmt.Sprintf("%s-%s", model.DocTypeService, s.ID)
	util.AddDyanmoDBKeys(avEntity, partitionKey, partitionKey)

	// [Company-{uuid}] | [Service-{uuid}] | ...
	avCompanyService, err := dynamodbattribute.MarshalMap(s)
	util.AddType(avCompanyService, model.DocTypeService)
	sortKey := fmt.Sprintf("%s-%s", model.DocTypeCompany, s.CompanyID)
	util.AddDyanmoDBKeys(avCompanyService, sortKey, partitionKey)

	input := &dynamodb.BatchWriteItemInput{
		RequestItems: map[string][]*dynamodb.WriteRequest{
			tableName: {
				{
					PutRequest: &dynamodb.PutRequest{
						Item: avEntity,
					},
				},
				{
					PutRequest: &dynamodb.PutRequest{
						Item: avCompanyService,
					},
				},
			},
		},
	}

	result, err := db.BatchWriteItem(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case dynamodb.ErrCodeProvisionedThroughputExceededException:
				fmt.Println(dynamodb.ErrCodeProvisionedThroughputExceededException, aerr.Error())
			case dynamodb.ErrCodeResourceNotFoundException:
				fmt.Println(dynamodb.ErrCodeResourceNotFoundException, aerr.Error())
			case dynamodb.ErrCodeItemCollectionSizeLimitExceededException:
				fmt.Println(dynamodb.ErrCodeItemCollectionSizeLimitExceededException, aerr.Error())
			case dynamodb.ErrCodeInternalServerError:
				fmt.Println(dynamodb.ErrCodeInternalServerError, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
	}

	fmt.Println(result)

	return err
}
