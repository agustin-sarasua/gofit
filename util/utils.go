package util

import (
	"encoding/json"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func GetClaimsSub(ctx events.APIGatewayProxyRequestContext) string {
	jc, _ := json.Marshal(ctx.Authorizer)
	fmt.Print(string(jc))
	r := make(map[string]interface{})
	err := json.Unmarshal(jc, &r)
	if err != nil {
		fmt.Printf("Something went wrong %v", err)
	}
	return r["principalId"].(string)
}

func AddType(av map[string]*dynamodb.AttributeValue, itype string) {
	av["DocType"] = &dynamodb.AttributeValue{
		S: aws.String(itype),
	}
}

func AddDyanmoDBKeys(av map[string]*dynamodb.AttributeValue, partitionKey string, rangeKey string) {
	av["PartitionKey"] = &dynamodb.AttributeValue{
		S: aws.String(partitionKey),
	}
	av["SortKey"] = &dynamodb.AttributeValue{
		S: aws.String(rangeKey),
	}
}
