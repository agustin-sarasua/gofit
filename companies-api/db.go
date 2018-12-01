package main

import (
	"fmt"
	"log"

	"github.com/agustin-sarasua/gofit-companies-api/model"
	"github.com/agustin-sarasua/gofit-companies-api/util"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
	"github.com/aws/aws-sdk-go/service/dynamodb/expression"
)

const tableName = "Companies"

var companiesUserSubGSI = "companiesUserSubGSI"

// Declare a new DynamoDB instance. Note that this is safe for concurrent
// use.
var db dynamodbiface.DynamoDBAPI = dynamodb.New(session.New(), aws.NewConfig().WithRegion("us-east-1"))

func putCompany(e *model.Company) error {
	av, err := dynamodbattribute.MarshalMap(e)
	if err != nil {
		panic(fmt.Sprintf("failed to DynamoDB marshal Record, %v", err))
	}
	util.AddType(av, model.DocTypeCompany)
	partitionKey := fmt.Sprintf("%s-%s", model.DocTypeCompany, e.ID)
	util.AddDyanmoDBKeys(av, partitionKey, partitionKey)

	input := &dynamodb.PutItemInput{
		TableName: aws.String(tableName),
		Item:      av,
	}

	_, err = db.PutItem(input)
	return err
}

func getUserCompanies(userSub string, limit int64) ([]*model.Company, error) {
	log.Printf("Loading user companies for %s \n", userSub)

	// Construct the Key condition builder
	keyCond := expression.Key("UserSub").
		Equal(expression.Value(userSub)).
		And(expression.KeyBeginsWith(expression.Key("SortKey"), fmt.Sprintf("%s-", model.DocTypeCompany)))

	// Construct the filter builder with a name and value.
	//filt := expression.Name("DocType").Equal(expression.Value("Company"))

	// Using the filter and projections create a DynamoDB expression from the two.
	expr, err := expression.NewBuilder().
		WithKeyCondition(keyCond).
		// WithFilter(filt).
		Build()
	if err != nil {
		fmt.Println(err)
	}

	// Prepare the input for the query.
	input := &dynamodb.QueryInput{
		TableName: aws.String(tableName),
		IndexName: aws.String(companiesUserSubGSI),
		Limit:     &limit,
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		KeyConditionExpression:    expr.KeyCondition(),
	}

	resp, err := db.Query(input)
	if err == nil {
		fmt.Println(resp)
		ps := []*model.Company{}
		err = dynamodbattribute.UnmarshalListOfMaps(resp.Items, &ps)
		log.Printf("Response count %d", *resp.Count)
		return ps, nil
	}
	log.Printf("Error %s", err.Error())
	return nil, err
}

func loadCompanyData(companyID string, limit int64) (*model.Company, error) {
	log.Printf("Loading company data for %s \n", companyID)

	// Construct the Key condition builder
	keyCond := expression.Key("PartitionKey").Equal(expression.Value(fmt.Sprintf("%s-%s", model.DocTypeCompany, companyID)))

	// Construct the filter builder with a name and value.
	//filt := expression.Name("DocType").Equal(expression.Value("Company"))

	// Using the filter and projections create a DynamoDB expression from the two.
	expr, err := expression.NewBuilder().
		WithKeyCondition(keyCond).
		//WithFilter(filt).
		Build()
	if err != nil {
		fmt.Println(err)
	}

	// Prepare the input for the query.
	input := &dynamodb.QueryInput{
		TableName: aws.String(tableName),
		Limit:     &limit,
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		KeyConditionExpression:    expr.KeyCondition(),
		FilterExpression:          expr.Filter(),
	}
	resp, err := db.Query(input)
	if err == nil {
		var company model.Company
		compStaff := make([]*model.Staff, 0)
		compServices := make([]*model.CompanyService, 0)
		// err = dynamodbattribute.UnmarshalListOfMaps(resp.Items, &ps)
		for _, m := range resp.Items {
			if t, ok := m["DocType"]; ok {
				if *t.S == model.DocTypeCompany {
					err = dynamodbattribute.UnmarshalMap(m, &company)
				} else if *t.S == model.DocTypeStaff {
					c := model.Staff{}
					err = dynamodbattribute.UnmarshalMap(m, &c)
					c.CompanyID = ""
					if err == nil {
						compStaff = append(compStaff, &c)
					}
				} else if *t.S == model.DocTypeService {
					c := model.CompanyService{}
					err = dynamodbattribute.UnmarshalMap(m, &c)
					c.CompanyID = ""
					if err == nil {
						compServices = append(compServices, &c)
					}
				} else {
					log.Print(m)
				}
			}
		}
		company.Staff = compStaff
		company.Services = compServices
		return &company, nil
	}
	return nil, err

}
