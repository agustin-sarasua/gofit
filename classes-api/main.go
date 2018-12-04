package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/agustin-sarasua/gofit/model"
	"github.com/agustin-sarasua/gofit/util"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/awslabs/aws-lambda-go-api-proxy/gin"
	"github.com/gin-gonic/gin"
)

var ginLambda *ginadapter.GinLambda

var errorLogger = log.New(os.Stderr, "ERROR ", log.Llongfile)
var infoLogger = log.New(os.Stdout, "INFO ", log.Llongfile)

func createClass(c *gin.Context) {
	apiGwContext, _ := ginLambda.GetAPIGatewayContext(c.Request)
	userSub := util.GetClaimsSub(apiGwContext)
	e := model.NewClass(userSub)
	err := c.BindJSON(&e)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
	}
	err = validateCreateClass(&e)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
	}
	// TODO validate UserSub exists
	err = putClass(&e)
	if err != nil {
		fmt.Printf("Error saving class in db %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
	}
	c.JSON(http.StatusCreated, e)
}

func loadUserClasses(c *gin.Context) {
	apiGwContext, _ := ginLambda.GetAPIGatewayContext(c.Request)
	userSub := util.GetClaimsSub(apiGwContext)

	c.JSON(http.StatusCreated, userSub)
}

// Handler is the main entry point for Lambda. Receives a proxy request and
// returns a proxy response
func Handler(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	if ginLambda == nil {
		// stdout and stderr are sent to AWS CloudWatch Logs
		log.Printf("Gin cold start")
		r := gin.Default()
		r.POST("/classes", createClass)
		r.GET("/classes", loadUserClasses)

		ginLambda = ginadapter.New(r)
	}

	return ginLambda.Proxy(req)
}

func main() {
	lambda.Start(Handler)
}
