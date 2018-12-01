package main

import (
	"testing"

	"github.com/agustin-sarasua/gofit/util"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetCompanyData(t *testing.T) {
	s := func(s string) *string { return &s }
	i := func(i int64) *int64 { return &i }
	auth := map[string]interface{}{
		"principalId": "example-sub",
	}

	tests := []struct {
		request events.APIGatewayProxyRequest
		expect  events.APIGatewayProxyResponse
		setUp   func()
		err     error
	}{
		{
			setUp: func() {
				dbMock := &util.MockDynamoDBClient{}
				db = dbMock
				items := make([]map[string]*dynamodb.AttributeValue, 0)
				items = append(items, map[string]*dynamodb.AttributeValue{
					"Status":    &dynamodb.AttributeValue{S: s("ACTIVE")},
					"CompanyID": &dynamodb.AttributeValue{S: s("429d1518-3350-4f09-9af7-64f764ac6628")},
					"DocType":   &dynamodb.AttributeValue{S: s("Company")},
					"UserSub":   &dynamodb.AttributeValue{S: s("example-sub")},
				}, map[string]*dynamodb.AttributeValue{
					"Status":    &dynamodb.AttributeValue{S: s("PENDING_CONFIRMATION")},
					"CompanyID": &dynamodb.AttributeValue{S: s("429d1518-3350-4f09-9af7-64f764ac6628")},
					"UserSub":   &dynamodb.AttributeValue{S: s("some-other-sub")},
					"DocType":   &dynamodb.AttributeValue{S: s("Staff")},
				})
				out := &dynamodb.QueryOutput{Items: items, Count: i(2)}
				dbMock.On("Query", mock.Anything).Return(out, nil)
			},
			// Test that the handler responds with the correct response
			// when a valid name is provided in the HTTP body
			request: events.APIGatewayProxyRequest{
				HTTPMethod:     "GET",
				Path:           "/companies/429d1518-3350-4f09-9af7-64f764ac6628",
				RequestContext: events.APIGatewayProxyRequestContext{Authorizer: auth},
			},
			expect: events.APIGatewayProxyResponse{
				StatusCode: 200,
				Body:       `{"UserSub":"example-sub","CompanyID":"429d1518-3350-4f09-9af7-64f764ac6628","Staff":[{"UserSub":"some-other-sub","Status":"PENDING_CONFIRMATION"}],"Status":"ACTIVE"}`,
			},
			err: nil,
		},
	}

	for _, test := range tests {
		test.setUp()
		response, err := Handler(test.request)
		assert.IsType(t, test.err, err)
		assert.Equal(t, test.expect.StatusCode, response.StatusCode)
		assert.Equal(t, test.expect.Body, response.Body)
	}

}

func TestListCompanies(t *testing.T) {
	// Setup Test
	s := func(s string) *string { return &s }
	i := func(i int64) *int64 { return &i }
	auth := map[string]interface{}{
		"principalId": "example-sub",
	}

	tests := []struct {
		setUp   func()
		request events.APIGatewayProxyRequest
		expect  events.APIGatewayProxyResponse
		err     error
	}{
		{
			setUp: func() {
				dbMock := &util.MockDynamoDBClient{}
				db = dbMock
				items := make([]map[string]*dynamodb.AttributeValue, 0)
				items = append(items, map[string]*dynamodb.AttributeValue{
					"Status":  &dynamodb.AttributeValue{S: s("ACTIVE")},
					"DocType": &dynamodb.AttributeValue{S: s("Company")},
				})
				out := &dynamodb.QueryOutput{Items: items, Count: i(1)}
				dbMock.On("Query", mock.Anything).Return(out, nil)
			},
			// Test that the handler responds with the correct response
			// when a valid name is provided in the HTTP body
			request: events.APIGatewayProxyRequest{
				HTTPMethod:     "GET",
				Path:           "/companies",
				RequestContext: events.APIGatewayProxyRequestContext{Authorizer: auth},
			},
			expect: events.APIGatewayProxyResponse{
				StatusCode: 200,
				Body:       `{"Companies":[{"UserSub":"","CompanyID":"","Status":"ACTIVE"}]}`,
			},
			err: nil,
		},
	}

	for _, test := range tests {
		test.setUp()
		response, err := Handler(test.request)
		assert.IsType(t, test.err, err)
		assert.Equal(t, test.expect.StatusCode, response.StatusCode)
		assert.Equal(t, test.expect.Body, response.Body)
	}
}
