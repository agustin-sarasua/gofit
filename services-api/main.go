package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/agustin-sarasua/gofit-companies-api/model"
	"github.com/agustin-sarasua/gofit-companies-api/util"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/awslabs/aws-lambda-go-api-proxy/gin"
	"github.com/gin-gonic/gin"
)

var ginLambda *ginadapter.GinLambda

var errorLogger = log.New(os.Stderr, "ERROR ", log.Llongfile)
var infoLogger = log.New(os.Stdout, "INFO ", log.Llongfile)

func createCompanyService(c *gin.Context) {
	companyID := c.Param("id")
	apiGwContext, _ := ginLambda.GetAPIGatewayContext(c.Request)
	userSub := util.GetClaimsSub(apiGwContext)
	e := model.NewCompanyService(companyID, userSub)
	err := c.BindJSON(&e)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
	}
	// TODO validate UserSub exists
	err = putCompanyService(&e)
	if err != nil {
		fmt.Printf("Error saving item in db %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
	}
	c.JSON(http.StatusCreated, e)
}

// Handler is the main entry point for Lambda. Receives a proxy request and
// returns a proxy response
func Handler(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	if ginLambda == nil {
		// stdout and stderr are sent to AWS CloudWatch Logs
		log.Printf("Gin cold start")
		r := gin.Default()
		r.POST("/companies/:id/services", createCompanyService)

		ginLambda = ginadapter.New(r)
	}

	return ginLambda.Proxy(req)
}

func main() {
	lambda.Start(Handler)
}
