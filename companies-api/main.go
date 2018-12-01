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

func createCompany(c *gin.Context) {
	apiGwContext, _ := ginLambda.GetAPIGatewayContext(c.Request)
	userSub := util.GetClaimsSub(apiGwContext)
	e := model.NewCompany(userSub)

	err := c.BindJSON(&e)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
	}
	if e.Name == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Name not found",
		})
		return
	}
	err = putCompany(&e)
	if err != nil {
		fmt.Printf("Error saving item in db %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
	}
	c.JSON(http.StatusCreated, e)
}

func listCompanies(c *gin.Context) {
	apiGwContext, _ := ginLambda.GetAPIGatewayContext(c.Request)
	userSub := util.GetClaimsSub(apiGwContext)
	cs, err := getUserCompanies(userSub, 10)
	if err != nil {
		fmt.Printf("Error saving item in db %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
	}
	log.Printf("Length of cs %d", len(cs))

	c.JSON(http.StatusOK, &struct {
		Companies []*model.Company `json:"Companies"`
	}{Companies: cs})
}

func getCompanyData(c *gin.Context) {
	companyID := c.Param("id")
	co, err := loadCompanyData(companyID, 100)
	if err != nil {
		fmt.Printf("Error saving item in db %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
	}
	c.JSON(http.StatusOK, co)
}

// Handler is the main entry point for Lambda. Receives a proxy request and
// returns a proxy response
func Handler(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	if ginLambda == nil {
		// stdout and stderr are sent to AWS CloudWatch Logs
		log.Printf("Gin cold start")
		r := gin.Default()
		r.GET("/companies", listCompanies)
		r.GET("/companies/:id", getCompanyData)
		r.POST("/companies", createCompany)

		ginLambda = ginadapter.New(r)
	}

	return ginLambda.Proxy(req)
}

func main() {
	lambda.Start(Handler)
}
